package main

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"github.com/go-playground/mold/v3"
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

func transformMyData(_ context.Context, _ *mold.Transformer, value reflect.Value, _ string) error {
	value.SetString("test")
	return nil
}
