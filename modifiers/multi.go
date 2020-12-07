package modifiers

import (
	"context"
	"reflect"
	"strconv"
	"time"

	"github.com/go-playground/mold/v3"
)

//
// Default allows setting of a default value IF no value is already present.
//
func Default(ctx context.Context, t *mold.Transformer, v reflect.Value, param string) error {
	if !v.IsZero() {
		return nil
	}

	switch v.Interface().(type) {
	case string:
		v.SetString(param)
	case int, int8, int16, int32, int64:
		value, err := strconv.Atoi(param)
		if err != nil {
			return err
		}
		v.SetInt(int64(value))

	case uint, uint8, uint16, uint32, uint64:
		value, err := strconv.Atoi(param)
		if err != nil {
			return err
		}
		v.SetUint(uint64(value))
	case float32, float64:
		value, err := strconv.ParseFloat(param, 64)
		if err != nil {
			return err
		}
		v.SetFloat(value)
	case bool:
		value, err := strconv.ParseBool(param)
		if err != nil {
			return err
		}
		v.SetBool(value)
	case time.Duration:
		d, err := time.ParseDuration(param)
		if err != nil {
			return err
		}
		v.SetInt(int64(d))
	}
	return nil
}
