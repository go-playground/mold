package modifiers

import (
	"context"
	"github.com/go-playground/mold/v3"
	"reflect"
	"strings"

	snakecase "github.com/segmentio/go-snakecase"
)

// TrimSpace trims extra space from text
func TrimSpace(_ context.Context, _ *mold.Transformer, v reflect.Value, _ string) error {
	s, ok := v.Interface().(string)
	if !ok {
		return nil
	}
	v.SetString(strings.TrimSpace(s))
	return nil
}

// TrimLeft trims extra left hand side of string using provided cutset
func TrimLeft(_ context.Context, _ *mold.Transformer, v reflect.Value, param string) error {
	s, ok := v.Interface().(string)
	if !ok {
		return nil
	}
	v.SetString(strings.TrimLeft(s, param))
	return nil
}

// TrimRight trims extra right hand side of string using provided cutset
func TrimRight(_ context.Context, _ *mold.Transformer, v reflect.Value, param string) error {
	s, ok := v.Interface().(string)
	if !ok {
		return nil
	}
	v.SetString(strings.TrimRight(s, param))
	return nil
}

// TrimPrefix trims the string of a prefix
func TrimPrefix(_ context.Context, _ *mold.Transformer, v reflect.Value, param string) error {
	s, ok := v.Interface().(string)
	if !ok {
		return nil
	}
	v.SetString(strings.TrimPrefix(s, param))
	return nil
}

// TrimSuffix trims the string of a suffix
func TrimSuffix(_ context.Context, _ *mold.Transformer, v reflect.Value, param string) error {
	s, ok := v.Interface().(string)
	if !ok {
		return nil
	}
	v.SetString(strings.TrimSuffix(s, param))
	return nil
}

// ToLower convert string to lower case
func ToLower(_ context.Context, _ *mold.Transformer, v reflect.Value, _ string) error {
	s, ok := v.Interface().(string)
	if !ok {
		return nil
	}
	v.SetString(strings.ToLower(s))
	return nil
}

// ToUpper convert string to upper case
func ToUpper(_ context.Context, _ *mold.Transformer, v reflect.Value, _ string) error {
	s, ok := v.Interface().(string)
	if !ok {
		return nil
	}
	v.SetString(strings.ToUpper(s))
	return nil
}

// SnakeCase converts string to snake case
func SnakeCase(_ context.Context, _ *mold.Transformer, v reflect.Value, _ string) error {
	s, ok := v.Interface().(string)
	if !ok {
		return nil
	}
	v.SetString(snakecase.Snakecase(s))
	return nil
}

// TODO: Add more
// - Snake_Case - can be combined with lowercase
// - CamelCase
// - many more
