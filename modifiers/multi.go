package modifiers

import (
	"context"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/mold/v4"
)

var (
	durationType = reflect.TypeOf(time.Duration(0))
	timeType     = reflect.TypeOf(time.Time{})
)

// defaultValue allows setting of a default value IF no value is already present.
func defaultValue(ctx context.Context, fl mold.FieldLevel) error {
	if !fl.Field().IsZero() {
		return nil
	}
	return setValue(ctx, fl)
}

func setValue(_ context.Context, fl mold.FieldLevel) error {
	return setValueInner(fl.Field(), fl.Param())
}

// setValue allows setting of a specified value
func setValueInner(field reflect.Value, param string) error {
	switch field.Kind() {
	case reflect.String:
		field.SetString(param)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
		value, err := strconv.Atoi(param)
		if err != nil {
			return err
		}
		field.SetInt(int64(value))

	case reflect.Int64:
		var value int64

		if field.Type() == durationType {
			d, err := time.ParseDuration(param)
			if err != nil {
				return err
			}
			value = int64(d)
		} else {
			i, err := strconv.Atoi(param)
			if err != nil {
				return err
			}
			value = int64(i)
		}
		field.SetInt(value)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		value, err := strconv.Atoi(param)
		if err != nil {
			return err
		}
		field.SetUint(uint64(value))

	case reflect.Float32, reflect.Float64:
		value, err := strconv.ParseFloat(param, 64)
		if err != nil {
			return err
		}
		field.SetFloat(value)

	case reflect.Bool:
		value, err := strconv.ParseBool(param)
		if err != nil {
			return err
		}
		field.SetBool(value)

	case reflect.Map:
		var n int
		var err error
		if param != "" {
			n, err = strconv.Atoi(param)
			if err != nil {
				return err
			}
		}
		field.Set(reflect.MakeMapWithSize(field.Type(), n))

	case reflect.Slice:
		var cap int
		var err error
		if param != "" {
			cap, err = strconv.Atoi(param)
			if err != nil {
				return err
			}
		}
		field.Set(reflect.MakeSlice(field.Type(), 0, cap))

	case reflect.Struct:
		if field.Type() == timeType {
			if param != "" {
				if strings.ToLower(param) == "utc" {
					field.Set(reflect.ValueOf(time.Now().UTC()))
				} else {
					t, err := time.Parse(time.RFC3339Nano, param)
					if err != nil {
						return err
					}
					field.Set(reflect.ValueOf(t))
				}
			} else {
				field.Set(reflect.ValueOf(time.Now()))
			}
		}
	case reflect.Chan:
		var buffer int
		var err error
		if param != "" {
			buffer, err = strconv.Atoi(param)
			if err != nil {
				return err
			}
		}
		field.Set(reflect.MakeChan(field.Type(), buffer))

	case reflect.Ptr:

		field.Set(reflect.New(field.Type().Elem()))
		return setValueInner(field.Elem(), param)
	}
	return nil
}

// empty sets the field to the zero value of the field type
func empty(_ context.Context, fl mold.FieldLevel) error {
	zeroValue := reflect.Zero(fl.Field().Type())
	fl.Field().Set(zeroValue)
	return nil
}
