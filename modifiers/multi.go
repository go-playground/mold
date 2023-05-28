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

// setValue allows setting of a specified value
func setValue(ctx context.Context, fl mold.FieldLevel) error {
	switch fl.Field().Kind() {
	case reflect.String:
		fl.Field().SetString(fl.Param())

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
		value, err := strconv.Atoi(fl.Param())
		if err != nil {
			return err
		}
		fl.Field().SetInt(int64(value))

	case reflect.Int64:
		var value int64

		if fl.Field().Type() == durationType {
			d, err := time.ParseDuration(fl.Param())
			if err != nil {
				return err
			}
			value = int64(d)
		} else {
			i, err := strconv.Atoi(fl.Param())
			if err != nil {
				return err
			}
			value = int64(i)
		}
		fl.Field().SetInt(value)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		value, err := strconv.Atoi(fl.Param())
		if err != nil {
			return err
		}
		fl.Field().SetUint(uint64(value))

	case reflect.Float32, reflect.Float64:
		value, err := strconv.ParseFloat(fl.Param(), 64)
		if err != nil {
			return err
		}
		fl.Field().SetFloat(value)

	case reflect.Bool:
		value, err := strconv.ParseBool(fl.Param())
		if err != nil {
			return err
		}
		fl.Field().SetBool(value)

	case reflect.Map:
		var n int
		var err error
		if fl.Param() != "" {
			n, err = strconv.Atoi(fl.Param())
			if err != nil {
				return err
			}
		}
		fl.Field().Set(reflect.MakeMapWithSize(fl.Field().Type(), n))

	case reflect.Slice:
		var cap int
		var err error
		if fl.Param() != "" {
			cap, err = strconv.Atoi(fl.Param())
			if err != nil {
				return err
			}
		}
		fl.Field().Set(reflect.MakeSlice(fl.Field().Type(), 0, cap))

	case reflect.Struct:
		if fl.Field().Type() == timeType {
			if fl.Param() != "" {
				if strings.ToLower(fl.Param()) == "utc" {
					fl.Field().Set(reflect.ValueOf(time.Now().UTC()))
				} else {
					t, err := time.Parse(time.RFC3339Nano, fl.Param())
					if err != nil {
						return err
					}
					fl.Field().Set(reflect.ValueOf(t))
				}
			} else {
				fl.Field().Set(reflect.ValueOf(time.Now()))
			}
		}
	case reflect.Chan:
		var buffer int
		var err error
		if fl.Param() != "" {
			buffer, err = strconv.Atoi(fl.Param())
			if err != nil {
				return err
			}
		}
		fl.Field().Set(reflect.MakeChan(fl.Field().Type(), buffer))

	case reflect.Ptr:
		fl.Field().Set(reflect.New(fl.Field().Type().Elem()))
	}
	return nil
}

// empty sets the field to the zero value of the field type
func empty(ctx context.Context, fl mold.FieldLevel) error {
	zeroValue := reflect.Zero(fl.Field().Type())
	fl.Field().Set(zeroValue)
	return nil
}
