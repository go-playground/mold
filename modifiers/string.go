package modifiers

import (
	"context"
	"reflect"
	"strings"

	"github.com/go-playground/mold"
)

// TrimSpace trims extra space from text
func TrimSpace(ctx context.Context, t *mold.Transformer, v reflect.Value) error {
	v.Set(reflect.ValueOf(strings.TrimSpace(v.String())))
	return nil
}

// TODO: Add more
// - Snake_Case - can be combined with lowercase
// - CamelCase
// - lowercase
// - many more
