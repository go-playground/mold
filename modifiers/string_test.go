package modifiers

import (
	"context"
	"log"
	"testing"
)

// NOTES:
// - Run "go test" to run tests
// - Run "gocov test | gocov report" to report on test converage by file
// - Run "gocov test | gocov annotate -" to report on all code and functions, those ,marked with "MISS" were never called
//
// or
//
// -- may be a good idea to change to output path to somewherelike /tmp
// go test -coverprofile cover.out && go tool cover -html=cover.out -o cover.html
//

func TestEmails(t *testing.T) {
	conform := New()

	email := "           Dean.Karn@gmail.com            "

	type Test struct {
		Email string `mod:"trim"`
	}

	tt := Test{Email: email}
	err := conform.Struct(context.Background(), &tt)
	if err != nil {
		log.Fatal(err)
	}
	if tt.Email != "Dean.Karn@gmail.com" {
		t.Fatalf("Unexpected value '%s'\n", tt.Email)
	}

	err = conform.Field(context.Background(), &email, "trim")
	if err != nil {
		log.Fatal(err)
	}
	if email != "Dean.Karn@gmail.com" {
		t.Fatalf("Unexpected value '%s'\n", tt.Email)
	}

	var iface interface{}
	err = conform.Field(context.Background(), &iface, "trim")
	if err != nil {
		log.Fatal(err)
	}
	if iface != nil {
		t.Fatalf("Unexpected value '%v'\n", nil)
	}

	iface = "    test     "
	err = conform.Field(context.Background(), &iface, "trim")
	if err != nil {
		log.Fatal(err)
	}
	if iface != "test" {
		t.Fatalf("Unexpected value '%s'\n", "test")
	}
}

func TestTrimLeft(t *testing.T) {
	conform := New()

	s := "#$%_test"
	expected := "test"

	type Test struct {
		String string `mod:"ltrim=#_$%"`
	}

	tt := Test{String: s}
	err := conform.Struct(context.Background(), &tt)
	if err != nil {
		log.Fatal(err)
	}
	if tt.String != expected {
		t.Fatalf("Unexpected value '%s'\n", tt.String)
	}

	err = conform.Field(context.Background(), &s, "ltrim=%_$#")
	if err != nil {
		log.Fatal(err)
	}
	if s != expected {
		t.Fatalf("Unexpected value '%s'\n", s)
	}

	var iface interface{}
	err = conform.Field(context.Background(), &iface, "ltrim=%_$#")
	if err != nil {
		log.Fatal(err)
	}
	if iface != nil {
		t.Fatalf("Unexpected value '%v'\n", nil)
	}

	iface = s
	err = conform.Field(context.Background(), &iface, "ltrim=%_$#")
	if err != nil {
		log.Fatal(err)
	}
	if iface != expected {
		t.Fatalf("Unexpected value '%v'\n", iface)
	}
}

func TestTrimRight(t *testing.T) {
	conform := New()

	s := "test#$%_"
	expected := "test"

	type Test struct {
		String string `mod:"rtrim=#_$%"`
	}

	tt := Test{String: s}
	err := conform.Struct(context.Background(), &tt)
	if err != nil {
		log.Fatal(err)
	}
	if tt.String != expected {
		t.Fatalf("Unexpected value '%s'\n", tt.String)
	}

	err = conform.Field(context.Background(), &s, "rtrim=#_$%")
	if err != nil {
		log.Fatal(err)
	}
	if s != expected {
		t.Fatalf("Unexpected value '%s'\n", s)
	}

	var iface interface{}
	err = conform.Field(context.Background(), &iface, "rtrim=#_$%")
	if err != nil {
		log.Fatal(err)
	}
	if iface != nil {
		t.Fatalf("Unexpected value '%v'\n", nil)
	}

	iface = s
	err = conform.Field(context.Background(), &iface, "rtrim=#_$%")
	if err != nil {
		log.Fatal(err)
	}
	if iface != expected {
		t.Fatalf("Unexpected value '%v'\n", iface)
	}
}

func TestTrimPrefix(t *testing.T) {
	conform := New()

	s := "pre-test"
	expected := "test"

	type Test struct {
		String string `mod:"tprefix=pre-"`
	}

	tt := Test{String: s}
	err := conform.Struct(context.Background(), &tt)
	if err != nil {
		log.Fatal(err)
	}
	if tt.String != expected {
		t.Fatalf("Unexpected value '%s'\n", tt.String)
	}

	err = conform.Field(context.Background(), &s, "tprefix=pre-")
	if err != nil {
		log.Fatal(err)
	}
	if s != expected {
		t.Fatalf("Unexpected value '%s'\n", s)
	}

	var iface interface{}
	err = conform.Field(context.Background(), &iface, "tprefix=pre-")
	if err != nil {
		log.Fatal(err)
	}
	if iface != nil {
		t.Fatalf("Unexpected value '%v'\n", nil)
	}

	iface = s
	err = conform.Field(context.Background(), &iface, "tprefix=pre-")
	if err != nil {
		log.Fatal(err)
	}
	if iface != expected {
		t.Fatalf("Unexpected value '%v'\n", iface)
	}
}

func TestTrimSuffix(t *testing.T) {
	conform := New()

	s := "test-suffix"
	expected := "test"

	type Test struct {
		String string `mod:"tsuffix=-suffix"`
	}

	tt := Test{String: s}
	err := conform.Struct(context.Background(), &tt)
	if err != nil {
		log.Fatal(err)
	}
	if tt.String != expected {
		t.Fatalf("Unexpected value '%s'\n", tt.String)
	}

	err = conform.Field(context.Background(), &s, "tsuffix=-suffix")
	if err != nil {
		log.Fatal(err)
	}
	if s != expected {
		t.Fatalf("Unexpected value '%s'\n", s)
	}

	var iface interface{}
	err = conform.Field(context.Background(), &iface, "tsuffix=-suffix")
	if err != nil {
		log.Fatal(err)
	}
	if iface != nil {
		t.Fatalf("Unexpected value '%v'\n", nil)
	}

	iface = s
	err = conform.Field(context.Background(), &iface, "tsuffix=-suffix")
	if err != nil {
		log.Fatal(err)
	}
	if iface != expected {
		t.Fatalf("Unexpected value '%v'\n", iface)
	}
}

func TestToLower(t *testing.T) {
	conform := New()

	s := "TEST"
	expected := "test"

	type Test struct {
		String string `mod:"lcase"`
	}

	tt := Test{String: s}
	err := conform.Struct(context.Background(), &tt)
	if err != nil {
		log.Fatal(err)
	}
	if tt.String != expected {
		t.Fatalf("Unexpected value '%s'\n", tt.String)
	}

	err = conform.Field(context.Background(), &s, "lcase")
	if err != nil {
		log.Fatal(err)
	}
	if s != expected {
		t.Fatalf("Unexpected value '%s'\n", s)
	}

	var iface interface{}
	err = conform.Field(context.Background(), &iface, "lcase")
	if err != nil {
		log.Fatal(err)
	}
	if iface != nil {
		t.Fatalf("Unexpected value '%v'\n", nil)
	}

	iface = s
	err = conform.Field(context.Background(), &iface, "lcase")
	if err != nil {
		log.Fatal(err)
	}
	if iface != expected {
		t.Fatalf("Unexpected value '%v'\n", iface)
	}
}

func TestToUpper(t *testing.T) {
	conform := New()

	s := "test"
	expected := "TEST"

	type Test struct {
		String string `mod:"ucase"`
	}

	tt := Test{String: s}
	err := conform.Struct(context.Background(), &tt)
	if err != nil {
		log.Fatal(err)
	}
	if tt.String != expected {
		t.Fatalf("Unexpected value '%s'\n", tt.String)
	}

	err = conform.Field(context.Background(), &s, "ucase")
	if err != nil {
		log.Fatal(err)
	}
	if s != expected {
		t.Fatalf("Unexpected value '%s'\n", s)
	}

	var iface interface{}
	err = conform.Field(context.Background(), &iface, "ucase")
	if err != nil {
		log.Fatal(err)
	}
	if iface != nil {
		t.Fatalf("Unexpected value '%v'\n", nil)
	}

	iface = s
	err = conform.Field(context.Background(), &iface, "ucase")
	if err != nil {
		log.Fatal(err)
	}
	if iface != expected {
		t.Fatalf("Unexpected value '%v'\n", iface)
	}
}

func TestSnakeCase(t *testing.T) {
	conform := New()

	s := "ThisIsSNAKEcase"
	expected := "this_is_snakecase"

	type Test struct {
		String string `mod:"snake"`
	}

	tt := Test{String: s}
	err := conform.Struct(context.Background(), &tt)
	if err != nil {
		log.Fatal(err)
	}
	if tt.String != expected {
		t.Fatalf("Unexpected value '%s'\n", tt.String)
	}

	err = conform.Field(context.Background(), &s, "snake")
	if err != nil {
		log.Fatal(err)
	}
	if s != expected {
		t.Fatalf("Unexpected value '%s'\n", s)
	}

	var iface interface{}
	err = conform.Field(context.Background(), &iface, "snake")
	if err != nil {
		log.Fatal(err)
	}
	if iface != nil {
		t.Fatalf("Unexpected value '%v'\n", nil)
	}

	iface = s
	err = conform.Field(context.Background(), &iface, "snake")
	if err != nil {
		log.Fatal(err)
	}
	if iface != expected {
		t.Fatalf("Unexpected value '%v'\n", iface)
	}
}

func TestTitleCase(t *testing.T) {
	conform := New()

	s := "this is a sentence"
	expected := "This Is A Sentence"

	type Test struct {
		String string `mod:"title"`
	}

	tt := Test{String: s}
	err := conform.Struct(context.Background(), &tt)
	if err != nil {
		log.Fatal(err)
	}
	if tt.String != expected {
		t.Fatalf("Unexpected value '%s'\n", tt.String)
	}

	err = conform.Field(context.Background(), &s, "title")
	if err != nil {
		log.Fatal(err)
	}
	if s != expected {
		t.Fatalf("Unexpected value '%s'\n", s)
	}

	var iface interface{}
	err = conform.Field(context.Background(), &iface, "title")
	if err != nil {
		log.Fatal(err)
	}
	if iface != nil {
		t.Fatalf("Unexpected value '%v'\n", nil)
	}

	iface = s
	err = conform.Field(context.Background(), &iface, "title")
	if err != nil {
		log.Fatal(err)
	}
	if iface != expected {
		t.Fatalf("Unexpected value '%v'\n", iface)
	}
}

func TestNumCase(t *testing.T) {
	conform := New()

	s := "the price is €30,38"
	expected := "3038"

	type Test struct {
		String string `mod:"strip_alpha"`
	}

	tt := Test{String: s}
	err := conform.Struct(context.Background(), &tt)
	if err != nil {
		log.Fatal(err)
	}
	if tt.String != expected {
		t.Fatalf("Unexpected value '%s'\n", tt.String)
	}

	err = conform.Field(context.Background(), &s, "strip_alpha")
	if err != nil {
		log.Fatal(err)
	}
	if s != expected {
		t.Fatalf("Unexpected value '%s'\n", s)
	}

	var iface interface{}
	err = conform.Field(context.Background(), &iface, "strip_alpha")
	if err != nil {
		log.Fatal(err)
	}
	if iface != nil {
		t.Fatalf("Unexpected value '%v'\n", nil)
	}

	iface = s
  
	err = conform.Field(context.Background(), &iface, "strip_alpha")
	if err != nil {
		log.Fatal(err)
	}
	if iface != expected {
		t.Fatalf("Unexpected value '%v'\n", iface)
	}
}

func TestNotNumCase(t *testing.T) {
	conform := New()

	s := "39472349D34a34v69e8932747"
	expected := "Dave"

	type Test struct {
		String string `mod:"strip_num"`
	}

	tt := Test{String: s}
	err := conform.Struct(context.Background(), &tt)
	if err != nil {
		log.Fatal(err)
	}
	if tt.String != expected {
		t.Fatalf("Unexpected value '%s'\n", tt.String)
	}

	err = conform.Field(context.Background(), &s, "strip_num")
	if err != nil {
		log.Fatal(err)
	}
	if s != expected {
		t.Fatalf("Unexpected value '%s'\n", s)
	}

	var iface interface{}
	err = conform.Field(context.Background(), &iface, "strip_num")
	if err != nil {
		log.Fatal(err)
	}
	if iface != nil {
		t.Fatalf("Unexpected value '%v'\n", nil)
	}

	iface = s
	err = conform.Field(context.Background(), &iface, "strip_num")
	if err != nil {
		log.Fatal(err)
	}
	if iface != expected {
		t.Fatalf("Unexpected value '%v'\n", iface)
	}
}

func TestAlphaCase(t *testing.T) {
	conform := New()

	s := "!@£$%^&'()Hello 1234567890 World+[];\\"
	expected := "HelloWorld"

	type Test struct {
		String string `mod:"strip_num_unicode"`
	}

	tt := Test{String: s}
	err := conform.Struct(context.Background(), &tt)
	if err != nil {
		log.Fatal(err)
	}
	if tt.String != expected {
		t.Fatalf("Unexpected value '%s'\n", tt.String)
	}

	err = conform.Field(context.Background(), &s, "strip_num_unicode")
	if err != nil {
		log.Fatal(err)
	}
	if s != expected {
		t.Fatalf("Unexpected value '%s'\n", s)
	}

	var iface interface{}
	err = conform.Field(context.Background(), &iface, "strip_num_unicode")
	if err != nil {
		log.Fatal(err)
	}
	if iface != nil {
		t.Fatalf("Unexpected value '%v'\n", nil)
	}

	iface = s
	err = conform.Field(context.Background(), &iface, "strip_num_unicode")
	if err != nil {
		log.Fatal(err)
	}
	if iface != expected {
		t.Fatalf("Unexpected value '%v'\n", iface)
	}
}

func TestNotAlphaCase(t *testing.T) {
	conform := New()

	s := "Everything's here but the letters!"
	expected := "'    !"

	type Test struct {
		String string `mod:"strip_alpha_unicode"`
	}

	tt := Test{String: s}
	err := conform.Struct(context.Background(), &tt)
	if err != nil {
		log.Fatal(err)
	}
	if tt.String != expected {
		t.Fatalf("Unexpected value '%s'\n", tt.String)
	}

	err = conform.Field(context.Background(), &s, "strip_alpha_unicode")
	if err != nil {
		log.Fatal(err)
	}
	if s != expected {
		t.Fatalf("Unexpected value '%s'\n", s)
	}

	var iface interface{}
	err = conform.Field(context.Background(), &iface, "strip_alpha_unicode")
	if err != nil {
		log.Fatal(err)
	}
	if iface != nil {
		t.Fatalf("Unexpected value '%v'\n", nil)
	}

	iface = s
	err = conform.Field(context.Background(), &iface, "strip_alpha_unicode")
	if err != nil {
		log.Fatal(err)
	}
	if iface != expected {
		t.Fatalf("Unexpected value '%v'\n", iface)
	}
}

func TestCamelCase(t *testing.T) {
	conform := New()

	s := "this_is_snakecase"
	expected := "thisIsSnakecase"

	type Test struct {
		String string `mod:"camel"`
	}

	tt := Test{String: s}
	err := conform.Struct(context.Background(), &tt)
	if err != nil {
		log.Fatal(err)
	}
	if tt.String != expected {
		t.Fatalf("Unexpected value '%s'\n", tt.String)
	}

	err = conform.Field(context.Background(), &s, "camel")
	if err != nil {
		log.Fatal(err)
	}
	if s != expected {
		t.Fatalf("Unexpected value '%s'\n", s)
	}

	var iface interface{}
	err = conform.Field(context.Background(), &iface, "camel")
	if err != nil {
		log.Fatal(err)
	}
	if iface != nil {
		t.Fatalf("Unexpected value '%v'\n", nil)
	}

	iface = s
	err = conform.Field(context.Background(), &iface, "camel")
	if err != nil {
		log.Fatal(err)
	}
	if iface != expected {
		t.Fatalf("Unexpected value '%v'\n", iface)
	}
}