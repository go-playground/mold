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
		Email string `mod:"trimspace"`
	}

	tt := Test{Email: email}
	err := conform.Struct(context.Background(), &tt)
	if err != nil {
		log.Fatal(err)
	}
	if tt.Email != "Dean.Karn@gmail.com" {
		t.Fatalf("Unexpected value '%s'\n", tt.Email)
	}

	err = conform.Field(context.Background(), &email, "trimspace")
	if err != nil {
		log.Fatal(err)
	}
	if email != "Dean.Karn@gmail.com" {
		t.Fatalf("Unexpected value '%s'\n", tt.Email)
	}
}
