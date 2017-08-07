package main

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"github.com/go-playground/mold"
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

func transformMyData(ctx context.Context, t *mold.Transformer, value reflect.Value) error {
	value.SetString("test")
	return nil
}
