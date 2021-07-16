package modifiers

import (
	"context"
	"reflect"
	"strconv"
	"time"

	"github.com/go-playground/mold/v4"
)

var (
	durationType = reflect.TypeOf(time.Duration(0))
)

//
// defaultValue allows setting of a default value IF no value is already present.
//
func defaultValue(ctx context.Context, fl mold.FieldLevel) error {
	if !fl.Field().IsZero() {
		return nil
	}

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

	}
	return nil
}
