package modifiers

import (
	"context"
	"strconv"
	"time"

	"github.com/go-playground/mold/v4"
)

//
// Default allows setting of a default value IF no value is already present.
//
func Default(ctx context.Context, fl mold.FieldLevel) error {
	if !fl.Field().IsZero() {
		return nil
	}

	switch fl.Field().Interface().(type) {
	case string:
		fl.Field().SetString(fl.Param())

	case int, int8, int16, int32, int64:
		value, err := strconv.Atoi(fl.Param())
		if err != nil {
			return err
		}
		fl.Field().SetInt(int64(value))

	case uint, uint8, uint16, uint32, uint64:
		value, err := strconv.Atoi(fl.Param())
		if err != nil {
			return err
		}
		fl.Field().SetUint(uint64(value))

	case float32, float64:
		value, err := strconv.ParseFloat(fl.Param(), 64)
		if err != nil {
			return err
		}
		fl.Field().SetFloat(value)

	case bool:
		value, err := strconv.ParseBool(fl.Param())
		if err != nil {
			return err
		}
		fl.Field().SetBool(value)

	case time.Duration:
		d, err := time.ParseDuration(fl.Param())
		if err != nil {
			return err
		}
		fl.Field().SetInt(int64(d))
	}
	return nil
}
