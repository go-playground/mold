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

var namePatterns = []map[string]string{
	{`[^\pL-\s']`: ""}, // cut off everything except [ alpha, hyphen, whitespace, apostrophe]
	{`\s{2,}`: " "},    // trim more than two whitespaces to one
	{`-{2,}`: "-"},     // trim more than two hyphens to one
	{`'{2,}`: "'"},     // trim more than two apostrophes to one
	{`( )*-( )*`: "-"}, // trim enclosing whitespaces around hyphen
}

var nameRegex = regexp.MustCompile(`[\p{L}]([\p{L}|[:space:]\-']*[\p{L}])*`)

// NameCase Trims, strips numbers and special characters (except dashes and spaces separating names),
// converts multiple spaces and dashes to single characters, title cases multiple names.
// Example: "3493€848Jo-$%£@Ann " -> "Jo-Ann", " ~~ The Dude ~~" -> "The Dude", "**susan**" -> "Susan",
// " hugh fearnley-whittingstall" -> "Hugh Fearnley-Whittingstall"
func NameCase(ctx context.Context, t *mold.Transformer, v reflect.Value, param string) error {
	s, ok := v.Interface().(string)
	if !ok {
		return nil
	}
	v.SetString(strings.Title(nameRegex.FindString(onlyOne(strings.ToLower(s)))))
	return nil
}

func onlyOne(s string) string {
	for _, v := range namePatterns {
		for f, r := range v {
			s = regexp.MustCompile(f).ReplaceAllLiteralString(s, r)
		}
	}
	return s
}

// TODO: Add more
// - Snake_Case - can be combined with lowercase
// - CamelCase
// - many more
