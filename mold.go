package mold

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"
)

var (
	timeType           = reflect.TypeOf(time.Time{})
	restrictedAliasErr = "Alias '%s' either contains restricted characters or is the same as a restricted tag needed for normal operation"
	restrictedTagErr   = "Tag '%s' either contains restricted characters or is the same as a restricted tag needed for normal operation"
)

// TODO - ensure StructLevel and Func get passed an interface and not *Transform directly

// Transform represents a subset of the current *Transformer that is executing the current transformation.
type Transform interface {
	Struct(ctx context.Context, v interface{}) error
	Field(ctx context.Context, v interface{}, tags string) error
}

// Func defines a transform function for use.
type Func func(ctx context.Context, fl FieldLevel) error

//
// StructLevelFunc accepts all values needed for struct level manipulation.
//
// Why does this exist? For structs for which you may not have access or rights to add tags too,
// from other packages your using.
//
type StructLevelFunc func(ctx context.Context, sl StructLevel) error

//
// InterceptorFunc is a way to intercept custom types to redirect the functions to be applied to an inner typ/value.
// eg. sql.NullString, the manipulation should be done on the inner string.
//
type InterceptorFunc func(current reflect.Value) (inner reflect.Value)

// Transformer is the base controlling object which contains
// all necessary information
type Transformer struct {
	tagName          string
	aliases          map[string]string
	transformations  map[string]Func
	structLevelFuncs map[reflect.Type]StructLevelFunc
	interceptors     map[reflect.Type]InterceptorFunc
	cCache           *structCache
	tCache           *tagCache
}

// New creates a new Transform object with default tag name of 'mold'
func New() *Transformer {
	tc := new(tagCache)
	tc.m.Store(make(map[string]*cTag))

	sc := new(structCache)
	sc.m.Store(make(map[reflect.Type]*cStruct))

	return &Transformer{
		tagName:         "mold",
		aliases:         make(map[string]string),
		transformations: make(map[string]Func),
		interceptors:    make(map[reflect.Type]InterceptorFunc),
		cCache:          sc,
		tCache:          tc,
	}
}

// SetTagName sets the given tag name to be used.
// Default is "trans"
func (t *Transformer) SetTagName(tagName string) {
	t.tagName = tagName
}

// Register adds a transformation with the given tag
//
// NOTES:
// - if the key already exists, the previous transformation function will be replaced.
// - this method is not thread-safe it is intended that these all be registered before hand
func (t *Transformer) Register(tag string, fn Func) {
	if len(tag) == 0 {
		panic("Function Key cannot be empty")
	}

	if fn == nil {
		panic("Function cannot be empty")
	}

	_, ok := restrictedTags[tag]

	if ok || strings.ContainsAny(tag, restrictedTagChars) {
		panic(fmt.Sprintf(restrictedTagErr, tag))
	}
	t.transformations[tag] = fn
}

// RegisterAlias registers a mapping of a single transform tag that
// defines a common or complex set of transformations to simplify adding transforms
// to structs.
//
// NOTE: this function is not thread-safe it is intended that these all be registered before hand
func (t *Transformer) RegisterAlias(alias, tags string) {
	if len(alias) == 0 {
		panic("Alias cannot be empty")
	}

	if len(tags) == 0 {
		panic("Aliased tags cannot be empty")
	}

	_, ok := restrictedTags[alias]

	if ok || strings.ContainsAny(alias, restrictedTagChars) {
		panic(fmt.Sprintf(restrictedAliasErr, alias))
	}
	t.aliases[alias] = tags
}

// RegisterStructLevel registers a StructLevelFunc against a number of types.
// Why does this exist? For structs for which you may not have access or rights to add tags too,
// from other packages your using.
//
// NOTES:
// - this method is not thread-safe it is intended that these all be registered prior to any validation
func (t *Transformer) RegisterStructLevel(fn StructLevelFunc, types ...interface{}) {
	if t.structLevelFuncs == nil {
		t.structLevelFuncs = make(map[reflect.Type]StructLevelFunc)
	}

	for _, typ := range types {
		t.structLevelFuncs[reflect.TypeOf(typ)] = fn
	}
}

//
// RegisterInterceptor registers a new interceptor functions agains one or more types.
// This InterceptorFunc allows one to intercept the incoming to to redirect the application of modifications
// to an inner type/value.
//
// eg. sql.NullString
//
func (t *Transformer) RegisterInterceptor(fn InterceptorFunc, types ...interface{}) {
	for _, typ := range types {
		t.interceptors[reflect.TypeOf(typ)] = fn
	}
}

// Struct applies transformations against the provided struct
func (t *Transformer) Struct(ctx context.Context, v interface{}) error {
	orig := reflect.ValueOf(v)

	if orig.Kind() != reflect.Ptr || orig.IsNil() {
		return &ErrInvalidTransformValue{typ: reflect.TypeOf(v), fn: "Struct"}
	}

	val := orig.Elem()
	typ := val.Type()

	if val.Kind() != reflect.Struct || val.Type() == timeType {
		return &ErrInvalidTransformation{typ: reflect.TypeOf(v)}
	}
	return t.setByStruct(ctx, orig, val, typ)
}

func (t *Transformer) setByStruct(ctx context.Context, parent, current reflect.Value, typ reflect.Type) (err error) {
	cs, ok := t.cCache.Get(typ)
	if !ok {
		if cs, err = t.extractStructCache(current); err != nil {
			return
		}
	}

	// run is struct has a corresponding struct level transformation
	if cs.fn != nil {
		if err = cs.fn(ctx, structLevel{
			transformer: t,
			parent:      parent,
			current:     current,
		}); err != nil {
			return
		}
	}

	var f *cField

	for i := 0; i < len(cs.fields); i++ {
		f = cs.fields[i]
		if err = t.setByField(ctx, current.Field(f.idx), f.cTags); err != nil {
			return
		}
	}
	return nil
}

// Field applies the provided transformations against the variable
func (t *Transformer) Field(ctx context.Context, v interface{}, tags string) (err error) {
	if len(tags) == 0 || tags == ignoreTag {
		return nil
	}

	val := reflect.ValueOf(v)

	if val.Kind() != reflect.Ptr || val.IsNil() {
		return &ErrInvalidTransformValue{typ: reflect.TypeOf(v), fn: "Field"}
	}
	val = val.Elem()

	// find cached tag
	ctag, ok := t.tCache.Get(tags)
	if !ok {
		t.tCache.lock.Lock()

		// could have been multiple trying to access, but once first is done this ensures tag
		// isn't parsed again.
		ctag, ok = t.tCache.Get(tags)
		if !ok {
			if ctag, _, err = t.parseFieldTagsRecursive(tags, "", "", false); err != nil {
				t.tCache.lock.Unlock()
				return
			}
			t.tCache.Set(tags, ctag)
		}
		t.tCache.lock.Unlock()
	}
	err = t.setByField(ctx, val, ctag)
	return
}

func (t *Transformer) setByField(ctx context.Context, orig reflect.Value, ct *cTag) (err error) {
	current, kind := t.extractType(orig)

	if ct != nil && ct.hasTag {
		for ct != nil {
			switch ct.typeof {
			case typeEndKeys:
				return
			case typeDive:
				ct = ct.next

				switch kind {
				case reflect.Slice, reflect.Array:
					err = t.setByIterable(ctx, current, ct)
				case reflect.Map:
					err = t.setByMap(ctx, current, ct)
				default:
					err = ErrInvalidDive
				}
				return

			default:
				if !current.CanAddr() {
					newVal := reflect.New(current.Type()).Elem()
					newVal.Set(current)
					if err = ct.fn(ctx, fieldLevel{
						transformer: t,
						parent:      orig,
						current:     newVal,
						param:       ct.param,
					}); err != nil {
						return
					}
					orig.Set(reflect.Indirect(newVal))
				} else {
					if err = ct.fn(ctx, fieldLevel{
						transformer: t,
						parent:      orig,
						current:     current,
						param:       ct.param,
					}); err != nil {
						return
					}
				}
				ct = ct.next
			}
		}
	}

	// need to do this again because one of the previous
	// sets could have set a struct value, where it was a
	// nil pointer before
	orig2 := current
	current, kind = t.extractType(current)

	if kind == reflect.Struct {
		typ := current.Type()
		if typ == timeType {
			return
		}

		if !current.CanAddr() {
			newVal := reflect.New(typ).Elem()
			newVal.Set(current)

			if err = t.setByStruct(ctx, orig, newVal, typ); err != nil {
				return
			}
			orig.Set(reflect.Indirect(newVal))
			return
		}
		err = t.setByStruct(ctx, orig2, current, typ)
	}
	return
}

func (t *Transformer) setByIterable(ctx context.Context, current reflect.Value, ct *cTag) (err error) {
	for i := 0; i < current.Len(); i++ {
		if err = t.setByField(ctx, current.Index(i), ct); err != nil {
			return
		}
	}
	return
}

func (t *Transformer) setByMap(ctx context.Context, current reflect.Value, ct *cTag) error {
	for _, key := range current.MapKeys() {
		newVal := reflect.New(current.Type().Elem()).Elem()
		newVal.Set(current.MapIndex(key))

		if ct != nil && ct.typeof == typeKeys && ct.keys != nil {
			// remove current map key as we may be changing it
			// and re-add to the map afterwards
			current.SetMapIndex(key, reflect.Value{})

			newKey := reflect.New(current.Type().Key()).Elem()
			newKey.Set(key)
			key = newKey

			// handle map key
			if err := t.setByField(ctx, key, ct.keys); err != nil {
				return err
			}

			// can be nil when just keys being validated
			if ct.next != nil {
				if err := t.setByField(ctx, newVal, ct.next); err != nil {
					return err
				}
			}
		} else {
			if err := t.setByField(ctx, newVal, ct); err != nil {
				return err
			}
		}
		current.SetMapIndex(key, newVal)
	}

	return nil
}
