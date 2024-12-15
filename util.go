package mold

import (
	"reflect"
	"strconv"
)

// extractType gets the actual underlying type of field value.
func (t *Transformer) extractType(current reflect.Value) (reflect.Value, reflect.Kind) {
	switch current.Kind() {
	case reflect.Ptr:
		if current.IsNil() {
			return current, reflect.Ptr
		}
		return t.extractType(current.Elem())

	case reflect.Interface:
		if current.IsNil() {
			return current, reflect.Interface
		}
		return t.extractType(current.Elem())

	default:
		if fn := t.interceptors[current.Type()]; fn != nil {
			return t.extractType(fn(current))
		}
		return current, current.Kind()
	}
}

// HasValue determines if a reflect.Value is it's default value
func HasValue(field reflect.Value) bool {
	switch field.Kind() {
	case reflect.Slice, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Chan, reflect.Func:
		return !field.IsNil()
	default:
		return field.IsValid() && field.Interface() != reflect.Zero(field.Type()).Interface()
	}
}

func GetPrimitiveValue(typ reflect.Kind, value string) (reflect.Value, error) {
	switch typ {

	case reflect.String:
		return reflect.ValueOf(value), nil

	case reflect.Int:
		value, err := strconv.Atoi(value)
		if err != nil {
			return reflect.Value{}, &ErrFailedToParseValue{typ: typ, err: err}
		}
		return reflect.ValueOf(value), nil

	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return reflect.Value{}, &ErrFailedToParseValue{typ: typ, err: err}
		}
		return reflect.ValueOf(int64(value)), nil

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		value, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return reflect.Value{}, &ErrFailedToParseValue{typ: typ, err: err}
		}
		return reflect.ValueOf(uint64(value)), nil

	case reflect.Float32, reflect.Float64:
		value, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return reflect.Value{}, &ErrFailedToParseValue{typ: typ, err: err}
		}
		return reflect.ValueOf(value), nil

	case reflect.Bool:
		value, err := strconv.ParseBool(value)
		if err != nil {
			return reflect.Value{}, &ErrFailedToParseValue{typ: typ, err: err}
		}
		return reflect.ValueOf(value), nil
	}
	return reflect.Value{}, &ErrUnsupportedType{typ: typ}
}
