package main

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/go-playground/form/v4"
	"github.com/go-playground/mold/v4/modifiers"
	"github.com/go-playground/mold/v4/scrubbers"
	"github.com/go-playground/validator/v10"
)

// This example is centered around a form post, but doesn't have to be
// just trying to give a well rounded real life example.

// <form method="POST">
//   <input type="text" name="Name" value="joeybloggs"/>
//   <input type="text" name="Age" value="3"/>
//   <input type="text" name="Gender" value="Male"/>
//   <input type="text" name="Address[0].Name" value="26 Here Blvd."/>
//   <input type="text" name="Address[0].Phone" value="9(999)999-9999"/>
//   <input type="text" name="Address[1].Name" value="26 There Blvd."/>
//   <input type="text" name="Address[1].Phone" value="1(111)111-1111"/>
//   <input type="text" name="active" value="true"/>
//   <input type="submit"/>
// </form>

var (
	conform  = modifiers.New()
	scrub    = scrubbers.New()
	validate = validator.New()
	decoder  = form.NewDecoder()
)

// Address contains address information
type Address struct {
	Name  string `mod:"trim" validate:"required"`
	Phone string `mod:"trim" validate:"required"`
}

// User contains user information
type User struct {
	Name    string            `mod:"trim"      validate:"required"              scrub:"name"`
	Age     uint8             `                validate:"required,gt=0,lt=130"`
	Gender  string            `                validate:"required"`
	Email   string            `mod:"trim"      validate:"required,email"        scrub:"emails"`
	Address []Address         `                validate:"required,dive"`
	Active  bool              `form:"active"`
	Misc    map[string]string `mod:"dive,keys,trim,endkeys,trim"`
}

func main() {
	// this simulates the results of http.Request's ParseForm() function
	values := parseForm()

	var user User

	// must pass a pointer
	err := decoder.Decode(&user, values)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Decoded:%+v\n\n", user)

	// great now lets conform our values, after all a human input the data
	// nobody's perfect
	err = conform.Struct(context.Background(), &user)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Conformed:%+v\n\n", user)

	// that's better all those extra spaces are gone
	// let's validate the data
	err = validate.Struct(user)
	if err != nil {
		log.Panic(err)
	}

	// ok now we know our data is good, let's do something with it like:
	// save to database
	// process request
	// etc....

	// ok now I'm done working with my data
	// let's log or store it somewhere
	// oh wait a minute, we have some sensitive PII data
	// let's make sure that's de-identified first
	err = scrub.Struct(context.Background(), &user)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Scrubbed:%+v\n\n", user)
}

// this simulates the results of http.Request's ParseForm() function
func parseForm() url.Values {
	return url.Values{
		"Name":             []string{"  joeybloggs  "},
		"Age":              []string{"3"},
		"Gender":           []string{"Male"},
		"Email":            []string{"Dean.Karn@gmail.com  "},
		"Address[0].Name":  []string{"26 Here Blvd."},
		"Address[0].Phone": []string{"9(999)999-9999"},
		"Address[1].Name":  []string{"26 There Blvd."},
		"Address[1].Phone": []string{"1(111)111-1111"},
		"active":           []string{"true"},
		"Misc[  b4  ]":     []string{"  b4  "},
	}
}
