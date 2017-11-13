package mold

import (
	"context"
	"errors"
	"reflect"
	"strings"
	"testing"
	"time"

	. "gopkg.in/go-playground/assert.v1"
)

// NOTES:
// - Run "go test" to run tests
// - Run "gocov test | gocov report" to report on test converage by file
// - Run "gocov test | gocov annotate -" to report on all code and functions, those ,marked with "MISS" were never called
//
// or
//
// -- may be a good idea to change to output path to somewherelike /tmp
// go test -coverprofile cover.out && go tool cover -html=cover.out -o cover.html
//

func TestBadValues(t *testing.T) {
	tform := New()
	tform.Register("blah", func(ctx context.Context, t *Transformer, value reflect.Value, param string) error { return nil })

	type Test struct {
		unexposed string
		Ignore    string `mold:"-"`
		String    string `mold:"blah,,blah"`
	}

	var tt Test

	err := tform.Struct(context.Background(), &tt)
	NotEqual(t, err, nil)
	Equal(t, err.Error(), "invalid tag '' found on field String")

	err = tform.Struct(context.Background(), tt)
	NotEqual(t, err, nil)
	Equal(t, err.Error(), "mold: Struct(non-pointer mold.Test)")

	err = tform.Struct(context.Background(), nil)
	NotEqual(t, err, nil)
	Equal(t, err.Error(), "mold: Struct(nil)")

	var i int
	err = tform.Struct(context.Background(), &i)
	NotEqual(t, err, nil)
	Equal(t, err.Error(), "mold: (nil *int)")

	var iface interface{}
	err = tform.Struct(context.Background(), iface)
	NotEqual(t, err, nil)
	Equal(t, err.Error(), "mold: Struct(nil)")

	iface = nil
	err = tform.Struct(context.Background(), &iface)
	NotEqual(t, err, nil)
	Equal(t, err.Error(), "mold: (nil *interface {})")

	var tst *Test
	err = tform.Struct(context.Background(), tst)
	NotEqual(t, err, nil)
	Equal(t, err.Error(), "mold: Struct(nil *mold.Test)")

	var tm *time.Time
	err = tform.Field(context.Background(), tm, "blah")
	NotEqual(t, err, nil)
	Equal(t, err.Error(), "mold: Field(nil *time.Time)")

	PanicMatches(t, func() { tform.Register("", nil) }, "Function Key cannot be empty")
	PanicMatches(t, func() { tform.Register("test", nil) }, "Function cannot be empty")
	PanicMatches(t, func() {
		tform.Register(",", func(ctx context.Context, t *Transformer, value reflect.Value, param string) error { return nil })
	}, "Tag ',' either contains restricted characters or is the same as a restricted tag needed for normal operation")

	PanicMatches(t, func() { tform.RegisterAlias("", "") }, "Alias cannot be empty")
	PanicMatches(t, func() { tform.RegisterAlias("test", "") }, "Aliased tags cannot be empty")
	PanicMatches(t, func() { tform.RegisterAlias(",", "test") }, "Alias ',' either contains restricted characters or is the same as a restricted tag needed for normal operation")
}

func TestBasicTransform(t *testing.T) {

	type Test struct {
		String string `r:"repl"`
	}

	var tt Test

	set := New()
	set.SetTagName("r")
	set.Register("repl", func(ctx context.Context, t *Transformer, value reflect.Value, param string) error {
		value.SetString("test")
		return nil
	})

	val := reflect.ValueOf(tt)
	// trigger a wait in struct parsing
	for i := 0; i < 3; i++ {
		_, err := set.extractStructCache(val)
		Equal(t, err, nil)
	}
	err := set.Struct(context.Background(), &tt)
	Equal(t, err, nil)
	Equal(t, tt.String, "test")

	type Test2 struct {
		Test   Test
		String string `r:"repl"`
	}

	var tt2 Test2

	err = set.Struct(context.Background(), &tt2)
	Equal(t, err, nil)
	Equal(t, tt2.Test.String, "test")
	Equal(t, tt2.String, "test")

	type Test3 struct {
		Test
		String string `r:"repl"`
	}

	var tt3 Test3

	err = set.Struct(context.Background(), &tt3)
	Equal(t, err, nil)
	Equal(t, tt3.Test.String, "test")
	Equal(t, tt3.String, "test")

	type Test4 struct {
		Test   *Test
		String string `r:"repl"`
	}

	var tt4 Test4

	err = set.Struct(context.Background(), &tt4)
	Equal(t, err, nil)
	Equal(t, tt4.Test, nil)
	Equal(t, tt4.String, "test")

	tt5 := Test4{Test: &Test{}}

	err = set.Struct(context.Background(), &tt5)
	Equal(t, err, nil)
	Equal(t, tt5.Test.String, "test")
	Equal(t, tt5.String, "test")

	type Test6 struct {
		Test   *Test  `r:"default"`
		String string `r:"repl"`
	}

	var tt6 Test6

	set.Register("default", func(ctx context.Context, t *Transformer, value reflect.Value, param string) error {
		value.Set(reflect.New(value.Type().Elem()))
		return nil
	})
	err = set.Struct(context.Background(), &tt6)
	Equal(t, err, nil)
	NotEqual(t, tt6.Test, nil)
	Equal(t, tt6.Test.String, "test")
	Equal(t, tt6.String, "test")

	tt6.String = "BAD"
	var tString string

	// wil invoke one processing and one waiting
	go func() {
		err := set.Field(context.Background(), &tString, "repl")
		Equal(t, err, nil)
	}()
	err = set.Field(context.Background(), &tt6.String, "repl")
	Equal(t, err, nil)
	Equal(t, tt6.String, "test")

	err = set.Field(context.Background(), &tt6.String, "")
	Equal(t, err, nil)

	err = set.Field(context.Background(), &tt6.String, "-")
	Equal(t, err, nil)

	err = set.Field(context.Background(), tt6.String, "test")
	NotEqual(t, err, nil)
	Equal(t, err.Error(), "mold: Field(non-pointer string)")

	err = set.Field(context.Background(), nil, "test")
	NotEqual(t, err, nil)
	Equal(t, err.Error(), "mold: Field(nil)")

	var iface interface{}
	err = set.Field(context.Background(), iface, "test")
	NotEqual(t, err, nil)
	Equal(t, err.Error(), "mold: Field(nil)")

	done := make(chan struct{})
	go func() {
		err := set.Field(context.Background(), &tString, "nonexistant")
		NotEqual(t, err, nil)
		close(done)
	}()

	err = set.Field(context.Background(), &tt6.String, "nonexistant")
	NotEqual(t, err, nil)
	Equal(t, err.Error(), "unregistered/undefined transformation 'nonexistant' found on field")

	<-done
	set.Register("dummy", func(ctx context.Context, t *Transformer, value reflect.Value, param string) error { return nil })
	err = set.Field(context.Background(), &tt6.String, "dummy")
	Equal(t, err, nil)
}

func TestAlias(t *testing.T) {

	type Test struct {
		String string `r:"repl,repl2"`
	}

	var tt Test

	set := New()
	set.SetTagName("r")
	set.Register("repl", func(ctx context.Context, t *Transformer, value reflect.Value, param string) error {
		value.SetString("test")
		return nil
	})
	set.Register("repl2", func(ctx context.Context, t *Transformer, value reflect.Value, param string) error {
		value.SetString("test2")
		return nil
	})

	err := set.Struct(context.Background(), &tt)
	Equal(t, err, nil)
	Equal(t, tt.String, "test2")

	set.RegisterAlias("rep", "repl,repl2")
	set.RegisterAlias("bad", "repl,,repl2")
	type Test2 struct {
		String string `r:"rep"`
	}

	var tt2 Test2

	err = set.Struct(context.Background(), &tt2)
	Equal(t, err, nil)
	Equal(t, tt.String, "test2")

	var s string
	err = set.Field(context.Background(), &s, "bad")
	NotEqual(t, err, nil)

	// var s string
	err = set.Field(context.Background(), &s, "repl,rep,bad")
	NotEqual(t, err, nil)
}

func TestArray(t *testing.T) {
	type Test struct {
		Arr []string `s:"defaultArr,dive,defaultStr"`
	}

	set := New()
	set.SetTagName("s")
	set.Register("defaultArr", func(ctx context.Context, t *Transformer, value reflect.Value, param string) error {
		if HasValue(value) {
			return nil
		}
		value.Set(reflect.MakeSlice(value.Type(), 2, 2))
		return nil
	})
	set.Register("defaultStr", func(ctx context.Context, t *Transformer, value reflect.Value, param string) error {
		if value.String() == "ok" {
			return errors.New("ALREADY OK")
		}
		value.SetString("default")
		return nil
	})

	var tt Test

	err := set.Struct(context.Background(), &tt)
	Equal(t, err, nil)
	Equal(t, len(tt.Arr), 2)
	Equal(t, tt.Arr[0], "default")
	Equal(t, tt.Arr[1], "default")

	tt2 := Test{
		Arr: make([]string, 1),
	}

	err = set.Struct(context.Background(), &tt2)
	Equal(t, err, nil)
	Equal(t, len(tt2.Arr), 1)
	Equal(t, tt2.Arr[0], "default")

	tt3 := Test{
		Arr: []string{"ok"},
	}

	err = set.Struct(context.Background(), &tt3)
	NotEqual(t, err, nil)
	Equal(t, err.Error(), "ALREADY OK")
}

func TestMap(t *testing.T) {
	type Test struct {
		Map map[string]string `s:"defaultMap,dive,defaultStr"`
	}

	set := New()
	set.SetTagName("s")
	set.Register("defaultMap", func(ctx context.Context, t *Transformer, value reflect.Value, param string) error {
		if HasValue(value) {
			return nil
		}
		value.Set(reflect.MakeMap(value.Type()))
		return nil
	})
	set.Register("defaultStr", func(ctx context.Context, t *Transformer, value reflect.Value, param string) error {
		if value.String() == "ok" {
			return errors.New("ALREADY OK")
		}
		value.SetString("default")
		return nil
	})

	var tt Test

	err := set.Struct(context.Background(), &tt)
	Equal(t, err, nil)
	Equal(t, len(tt.Map), 0)

	tt2 := Test{
		Map: map[string]string{"key": ""},
	}

	err = set.Struct(context.Background(), &tt2)
	Equal(t, err, nil)
	Equal(t, len(tt2.Map), 1)
	Equal(t, tt2.Map["key"], "default")

	tt3 := Test{
		Map: map[string]string{"key": "ok"},
	}

	err = set.Struct(context.Background(), &tt3)
	NotEqual(t, err, nil)
	Equal(t, err.Error(), "ALREADY OK")
}

func TestInterface(t *testing.T) {
	type Test struct {
		Iface interface{} `s:"default"`
	}

	type Inner struct {
		STR    string
		String string `s:"defaultStr"`
	}

	type Test2 struct {
		Iface interface{} `s:"default2"`
	}

	type Inner2 struct {
		String string `s:"error"`
	}

	set := New()
	set.SetTagName("s")
	set.Register("default", func(ctx context.Context, t *Transformer, value reflect.Value, param string) error {
		value.Set(reflect.ValueOf(Inner{STR: "test"}))
		return nil
	})
	set.Register("default2", func(ctx context.Context, t *Transformer, value reflect.Value, param string) error {
		value.Set(reflect.ValueOf(Inner2{}))
		return nil
	})
	set.Register("defaultStr", func(ctx context.Context, t *Transformer, value reflect.Value, param string) error {
		if HasValue(value) && value.String() == "ok" {
			return errors.New("ALREADY OK")
		}
		value.Set(reflect.ValueOf("default"))
		return nil
	})
	set.Register("error", func(ctx context.Context, t *Transformer, value reflect.Value, param string) error {
		return errors.New("BAD VALUE")
	})

	var tt Test

	err := set.Struct(context.Background(), &tt)
	Equal(t, err, nil)
	NotEqual(t, tt.Iface, nil)

	inner, ok := tt.Iface.(Inner)
	Equal(t, ok, true)
	Equal(t, inner.String, "default")
	Equal(t, inner.STR, "test")

	var tt2 Test2

	err = set.Struct(context.Background(), &tt2)
	NotEqual(t, err, nil)

	type Test3 struct {
		Iface interface{} `s:"defaultStr"`
	}

	var tt3 Test3
	tt3.Iface = "String"
	err = set.Struct(context.Background(), &tt3)
	Equal(t, err, nil)
	Equal(t, tt3.Iface.(string), "default")

	type Test4 struct {
		Iface interface{} `s:"defaultStr,defaultStr"`
	}

	var tt4 Test4
	tt4.Iface = nil
	err = set.Struct(context.Background(), &tt4)
	Equal(t, err, nil)
	Equal(t, tt4.Iface.(string), "default")

	type Test5 struct {
		Iface interface{} `s:"defaultStr,error"`
	}

	var tt5 Test5
	tt5.Iface = "String"
	err = set.Struct(context.Background(), &tt5)
	NotEqual(t, err, nil)
}

func TestInterfacePtr(t *testing.T) {
	type Test struct {
		Iface interface{} `s:"default"`
	}

	type Inner struct {
		String string `s:"defaultStr"`
	}

	set := New()
	set.SetTagName("s")
	set.Register("default", func(ctx context.Context, t *Transformer, value reflect.Value, param string) error {
		value.Set(reflect.ValueOf(new(Inner)))
		return nil
	})
	set.Register("defaultStr", func(ctx context.Context, t *Transformer, value reflect.Value, param string) error {
		if value.String() == "ok" {
			return errors.New("ALREADY OK")
		}
		value.SetString("default")
		return nil
	})

	var tt Test

	err := set.Struct(context.Background(), &tt)
	Equal(t, err, nil)
	NotEqual(t, tt.Iface, nil)

	inner, ok := tt.Iface.(*Inner)
	Equal(t, ok, true)
	Equal(t, inner.String, "default")

	type Test2 struct {
		Iface interface{}
	}

	var tt2 Test2
	tt2.Iface = Inner{}
	err = set.Struct(context.Background(), &tt2)
	Equal(t, err, nil)
}

func TestStructLevel(t *testing.T) {
	type Test struct {
		String string
	}

	set := New()
	set.RegisterStructLevel(func(ctx context.Context, t *Transformer, value reflect.Value) error {
		s := value.Interface().(Test)
		if s.String == "error" {
			return errors.New("BAD VALUE")
		}
		s.String = "test"
		value.Set(reflect.ValueOf(s))
		return nil
	}, Test{})

	var tt Test
	err := set.Struct(context.Background(), &tt)
	Equal(t, err, nil)
	Equal(t, tt.String, "test")

	tt.String = "error"
	err = set.Struct(context.Background(), &tt)
	NotEqual(t, err, nil)
}

func TestTimeType(t *testing.T) {

	var tt time.Time

	set := New()
	set.Register("default", func(ctx context.Context, t *Transformer, value reflect.Value, param string) error {
		value.Set(reflect.ValueOf(time.Now()))
		return nil
	})

	err := set.Field(context.Background(), &tt, "default")
	Equal(t, err, nil)

	err = set.Field(context.Background(), &tt, "default,dive")
	NotEqual(t, err, nil)
	Equal(t, err.Error(), "Invalid dive tag configuration")
}

func TestParam(t *testing.T) {

	type Test struct {
		String string `r:"ltrim=#$_"`
	}

	set := New()
	set.SetTagName("r")
	set.Register("ltrim", func(ctx context.Context, t *Transformer, value reflect.Value, param string) error {
		value.SetString(strings.TrimLeft(value.String(), param))
		return nil
	})

	tt := Test{String: "_test"}

	err := set.Struct(context.Background(), &tt)
	Equal(t, err, nil)
	Equal(t, tt.String, "test")
}

func TestDiveKeys(t *testing.T) {

	type Test struct {
		Map map[string]string `s:"dive,keys,default,endkeys,default"`
	}

	set := New()
	set.SetTagName("s")
	set.Register("default", func(ctx context.Context, t *Transformer, value reflect.Value, param string) error {
		value.Set(reflect.ValueOf("after"))
		return nil
	})
	set.Register("err", func(ctx context.Context, t *Transformer, value reflect.Value, param string) error {
		return errors.New("err")
	})

	test := Test{
		Map: map[string]string{
			"b4": "b4",
		},
	}

	err := set.Struct(context.Background(), &test)
	Equal(t, err, nil)

	val := test.Map["after"]
	Equal(t, val, "after")

	m := map[string]string{
		"b4": "b4",
	}

	err = set.Field(context.Background(), &m, "dive,keys,default,endkeys,default")
	Equal(t, err, nil)

	val = m["after"]
	Equal(t, val, "after")

	err = set.Field(context.Background(), &m, "keys,endkeys,default")
	Equal(t, err, ErrInvalidKeysTag)

	err = set.Field(context.Background(), &m, "dive,endkeys,default")
	Equal(t, err, ErrUndefinedKeysTag)

	err = set.Field(context.Background(), &m, "dive,keys,undefinedtag")
	Equal(t, err, ErrUndefinedTag{tag: "undefinedtag"})

	err = set.Field(context.Background(), &m, "dive,keys,err,endkeys")
	NotEqual(t, err, nil)

	m = map[string]string{
		"b4": "b4",
	}
	err = set.Field(context.Background(), &m, "dive,keys,default,endkeys,err")
	NotEqual(t, err, nil)
}
