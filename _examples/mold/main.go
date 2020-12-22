package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-playground/mold/v4"
)

var tform *mold.Transformer

func main() {
	tform = mold.New()
	tform.Register("set", transformMyData)

	type Test struct {
		String string `mold:"set"`
	}

	var tt Test

	err := tform.Struct(context.Background(), &tt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", tt)

	var myString string
	err = tform.Field(context.Background(), &myString, "set")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(myString)
}

func transformMyData(ctx context.Context, fl mold.FieldLevel) error {
	fl.Field().SetString("test")
	return nil
}
