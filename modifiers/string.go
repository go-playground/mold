package modifiers

import (
	"bytes"
	"context"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/go-playground/mold/v4"
	"github.com/segmentio/go-camelcase"
	"github.com/segmentio/go-snakecase"
)

// TrimSpace trims extra space from text
func TrimSpace(ctx context.Context, fl mold.FieldLevel) error {
	s, ok := fl.Field().Interface().(string)
	if !ok {
		return nil
	}
	fl.Field().SetString(strings.TrimSpace(s))
	return nil
}

// TrimLeft trims extra left hand side of string using provided cutset
func TrimLeft(ctx context.Context, fl mold.FieldLevel) error {
	s, ok := fl.Field().Interface().(string)
	if !ok {
		return nil
	}
	fl.Field().SetString(strings.TrimLeft(s, fl.Param()))
	return nil
}

// TrimRight trims extra right hand side of string using provided cutset
func TrimRight(ctx context.Context, fl mold.FieldLevel) error {
	s, ok := fl.Field().Interface().(string)
	if !ok {
		return nil
	}
	fl.Field().SetString(strings.TrimRight(s, fl.Param()))
	return nil
}

// TrimPrefix trims the string of a prefix
func TrimPrefix(ctx context.Context, fl mold.FieldLevel) error {
	s, ok := fl.Field().Interface().(string)
	if !ok {
		return nil
	}
	fl.Field().SetString(strings.TrimPrefix(s, fl.Param()))
	return nil
}

// TrimSuffix trims the string of a suffix
func TrimSuffix(ctx context.Context, fl mold.FieldLevel) error {
	s, ok := fl.Field().Interface().(string)
	if !ok {
		return nil
	}
	fl.Field().SetString(strings.TrimSuffix(s, fl.Param()))
	return nil
}

// ToLower convert string to lower case
func ToLower(ctx context.Context, fl mold.FieldLevel) error {
	s, ok := fl.Field().Interface().(string)
	if !ok {
		return nil
	}
	fl.Field().SetString(strings.ToLower(s))
	return nil
}

// ToUpper convert string to upper case
func ToUpper(ctx context.Context, fl mold.FieldLevel) error {
	s, ok := fl.Field().Interface().(string)
	if !ok {
		return nil
	}
	fl.Field().SetString(strings.ToUpper(s))
	return nil
}

// SnakeCase converts string to snake case
func SnakeCase(ctx context.Context, fl mold.FieldLevel) error {
	s, ok := fl.Field().Interface().(string)
	if !ok {
		return nil
	}
	fl.Field().SetString(snakecase.Snakecase(s))
	return nil
}

// TitleCase converts string to title case, e.g. "this is a sentence" -> "This Is A Sentence"
func TitleCase(ctx context.Context, fl mold.FieldLevel) error {
	s, ok := fl.Field().Interface().(string)
	if !ok {
		return nil
	}
	fl.Field().SetString(strings.Title(s))
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
func NameCase(ctx context.Context, fl mold.FieldLevel) error {
	s, ok := fl.Field().Interface().(string)
	if !ok {
		return nil
	}
	fl.Field().SetString(strings.Title(nameRegex.FindString(onlyOne(strings.ToLower(s)))))
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

// UppercaseFirstCharacterCase converts a string so that it has only the first capital letter. Example: "all lower" -> "All lower"
func UppercaseFirstCharacterCase(_ context.Context, fl mold.FieldLevel) error {
	s, ok := fl.Field().Interface().(string)
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
	fl.Field().SetString(buf.String())
	return nil
}

var stripNumRegex = regexp.MustCompile("[^0-9]")

// StripAlphaCase removes all non-numeric characters. Example: "the price is €30,38" -> "3038". Note: The struct field will remain a string. No type conversion takes place.
func StripAlphaCase(_ context.Context, fl mold.FieldLevel) error {
	s, ok := fl.Field().Interface().(string)
	if !ok {
		return nil
	}
	fl.Field().SetString(stripNumRegex.ReplaceAllLiteralString(s, ""))
	return nil
}

var stripAlphaRegex = regexp.MustCompile("[0-9]")

// StripNumCase removes all numbers. Example "39472349D34a34v69e8932747" -> "Dave". Note: The struct field will remain a string. No type conversion takes place.
func StripNumCase(_ context.Context, fl mold.FieldLevel) error {
	s, ok := fl.Field().Interface().(string)
	if !ok {
		return nil
	}
	fl.Field().SetString(stripAlphaRegex.ReplaceAllLiteralString(s, ""))
	return nil
}

var stripNumUnicodeRegex = regexp.MustCompile(`[^\pL]`)

// StripNumUnicodeCase removes non-alpha unicode characters. Example: "!@£$%^&'()Hello 1234567890 World+[];\" -> "HelloWorld"
func StripNumUnicodeCase(ctx context.Context, fl mold.FieldLevel) error {
	s, ok := fl.Field().Interface().(string)
	if !ok {
		return nil
	}
	fl.Field().SetString(stripNumUnicodeRegex.ReplaceAllLiteralString(s, ""))
	return nil
}

var stripAlphaUnicode = regexp.MustCompile(`[\pL]`)

// StripAlphaUnicodeCase removes alpha unicode characters. Example: "Everything's here but the letters!" -> "' !"
func StripAlphaUnicodeCase(ctx context.Context, fl mold.FieldLevel) error {
	s, ok := fl.Field().Interface().(string)
	if !ok {
		return nil
	}
	fl.Field().SetString(stripAlphaUnicode.ReplaceAllLiteralString(s, ""))
	return nil
}

// CamelCase converts string to camel case
func CamelCase(ctx context.Context, fl mold.FieldLevel) error {
	s, ok := fl.Field().Interface().(string)
	if !ok {
		return nil
	}
	fl.Field().SetString(camelcase.Camelcase(s))
	return nil
}
