package modifiers

import (
	"bytes"
	"context"
	"reflect"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/go-playground/mold/v3"
	"github.com/segmentio/go-camelcase"
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

// TitleCase converts string to title case, e.g. "this is a sentence" -> "This Is A Sentence"
func TitleCase(ctx context.Context, t *mold.Transformer, v reflect.Value, param string) error {
	s, ok := v.Interface().(string)
	if !ok {
		return nil
	}
	v.SetString(strings.Title(s))
	return nil
}

// UppercaseFirstCharacterCase converts a string so that it has only the first capital letter. Example: "all lower" -> "All lower"
func UppercaseFirstCharacterCase(_ context.Context, _ *mold.Transformer, v reflect.Value, _ string) error {
	s, ok := v.Interface().(string)
	if !ok {
		return nil
	}
	if s == "" {
		return nil
	}
	toRune, size := utf8.DecodeRuneInString(s)
	if !unicode.IsLower(toRune) {
		return nil
	}
	buf := &bytes.Buffer{}
	buf.WriteRune(unicode.ToUpper(toRune))
	buf.WriteString(s[size:])
	v.SetString(buf.String())
	return nil
}

var stripNumRegex = regexp.MustCompile("[^0-9]")

// StripAlphaCase removes all non-numeric characters. Example: "the price is €30,38" -> "3038". Note: The struct field will remain a string. No type conversion takes place.
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
func StripNumCase(_ context.Context, _ *mold.Transformer, v reflect.Value, _ string) error {
	s, ok := v.Interface().(string)
	if !ok {
		return nil
	}
	v.SetString(stripAlphaRegex.ReplaceAllLiteralString(s, ""))
	return nil
}

var stripNumUnicodeRegex = regexp.MustCompile(`[^\pL]`)

// StripNumUnicodeCase removes non-alpha unicode characters. Example: "!@£$%^&'()Hello 1234567890 World+[];\" -> "HelloWorld"
func StripNumUnicodeCase(ctx context.Context, t *mold.Transformer, v reflect.Value, param string) error {
	s, ok := v.Interface().(string)
	if !ok {
		return nil
	}
	v.SetString(stripNumUnicodeRegex.ReplaceAllLiteralString(s, ""))
	return nil
}

var stripAlphaUnicode = regexp.MustCompile(`[\pL]`)

// StripAlphaUnicodeCase removes alpha unicode characters. Example: "Everything's here but the letters!" -> "' !"
func StripAlphaUnicodeCase(ctx context.Context, t *mold.Transformer, v reflect.Value, param string) error {
	s, ok := v.Interface().(string)
	if !ok {
		return nil
	}
	v.SetString(stripAlphaUnicode.ReplaceAllLiteralString(s, ""))
	return nil
}

// CamelCase converts string to camel case
func CamelCase(ctx context.Context, t *mold.Transformer, v reflect.Value, param string) error {
	s, ok := v.Interface().(string)
	if !ok {
		return nil
	}
	v.SetString(camelcase.Camelcase(s))
	return nil
}