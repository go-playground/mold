package modifiers

import (
	"context"
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/mold/v3"

	snakecase "github.com/segmentio/go-snakecase"
)

// TrimSpace trims extra space from text
func TrimSpace(ctx context.Context, t *mold.Transformer, v reflect.Value, param string) error {
	s, ok := v.Interface().(string)
	if !ok {
		return nil
	}
	v.SetString(strings.TrimSpace(s))
	return nil
}

// TrimLeft trims extra left hand side of string using provided cutset
func TrimLeft(ctx context.Context, t *mold.Transformer, v reflect.Value, param string) error {
	s, ok := v.Interface().(string)
	if !ok {
		return nil
	}
	v.SetString(strings.TrimLeft(s, param))
	return nil
}

// TrimRight trims extra right hand side of string using provided cutset
func TrimRight(ctx context.Context, t *mold.Transformer, v reflect.Value, param string) error {
	s, ok := v.Interface().(string)
	if !ok {
		return nil
	}
	v.SetString(strings.TrimRight(s, param))
	return nil
}

// TrimPrefix trims the string of a prefix
func TrimPrefix(ctx context.Context, t *mold.Transformer, v reflect.Value, param string) error {
	s, ok := v.Interface().(string)
	if !ok {
		return nil
	}
	v.SetString(strings.TrimPrefix(s, param))
	return nil
}

// TrimSuffix trims the string of a suffix
func TrimSuffix(ctx context.Context, t *mold.Transformer, v reflect.Value, param string) error {
	s, ok := v.Interface().(string)
	if !ok {
		return nil
	}
	v.SetString(strings.TrimSuffix(s, param))
	return nil
}

// ToLower convert string to lower case
func ToLower(ctx context.Context, t *mold.Transformer, v reflect.Value, param string) error {
	s, ok := v.Interface().(string)
	if !ok {
		return nil
	}
	v.SetString(strings.ToLower(s))
	return nil
}

// ToUpper convert string to upper case
func ToUpper(ctx context.Context, t *mold.Transformer, v reflect.Value, param string) error {
	s, ok := v.Interface().(string)
	if !ok {
		return nil
	}
	v.SetString(strings.ToUpper(s))
	return nil
}

// SnakeCase converts string to snake case
func SnakeCase(ctx context.Context, t *mold.Transformer, v reflect.Value, param string) error {
	s, ok := v.Interface().(string)
	if !ok {
		return nil
	}
	v.SetString(snakecase.Snakecase(s))
	return nil
}

// TitleCase converts string to title case, e.g. "this is a sentence" -> "This Is A Sentence"
func TitleCase(ctx context.Context, t *mold.Transformer, v reflect.Value, param string) error {
	s, ok := v.Interface().(string)
	if !ok {
		return nil
	}
	v.SetString(strings.Title(s))
	return nil
}

var stripNumRegex = regexp.MustCompile("[^0-9]")

// StripAlphaCase removes all non-numeric characters. Example: "the price is €30,38" -> "3038". Note: The struct field will remain a string. No type conversion takes place.
// From https://github.com/leebenson/conform
func StripAlphaCase(_ context.Context, _ *mold.Transformer, v reflect.Value, _ string) error {
	s, ok := v.Interface().(string)
	if !ok {
		return nil
	}
	v.SetString(stripNumRegex.ReplaceAllLiteralString(s, ""))
	return nil
}

var stripAlphaRegex = regexp.MustCompile("[0-9]")

// StripNumCase removes all numbers. Example "39472349D34a34v69e8932747" -> "Dave". Note: The struct field will remain a string. No type conversion takes place.
// From https://github.com/leebenson/conform
func StripNumCase(_ context.Context, _ *mold.Transformer, v reflect.Value, _ string) error {
	s, ok := v.Interface().(string)
	if !ok {
		return nil
	}
	v.SetString(stripAlphaRegex.ReplaceAllLiteralString(s, ""))
	return nil
}

// TODO: Add more
// - Snake_Case - can be combined with lowercase
// - CamelCase
// - many more
