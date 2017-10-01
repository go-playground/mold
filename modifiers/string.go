package modifiers

import (
	"context"
	"reflect"
	"strings"

	"github.com/go-playground/mold"
	snakecase "github.com/segmentio/go-snakecase"
)

// TrimSpace trims extra space from text
func TrimSpace(ctx context.Context, t *mold.Transformer, v reflect.Value) error {
	v.Set(reflect.ValueOf(strings.TrimSpace(v.String())))
	return nil
}

// ToLower convert string to lower case
func ToLower(ctx context.Context, t *mold.Transformer, v reflect.Value) error {
	v.Set(reflect.ValueOf(strings.ToLower(v.String())))
	return nil
}

// ToUpper convert string to upper case
func ToUpper(ctx context.Context, t *mold.Transformer, v reflect.Value) error {
	v.Set(reflect.ValueOf(strings.ToUpper(v.String())))
	return nil
}

// SnakeCase converts string to snake case
func SnakeCase(ctx context.Context, t *mold.Transformer, v reflect.Value) error {
	v.Set(reflect.ValueOf(snakecase.Snakecase(v.String())))
	return nil
}

// TODO: Add more
// - Snake_Case - can be combined with lowercase
// - CamelCase
// - many more
