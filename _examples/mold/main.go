package main

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"gopkg.in/go-playground/mold.v2"
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

func transformMyData(ctx context.Context, t *mold.Transformer, value reflect.Value, param string) error {
	value.SetString("test")
	return nil
}
