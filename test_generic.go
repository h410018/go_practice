package main

import (
	"fmt"
	"reflect"
)

func testGeneric[T any](v T) {
	if reflect.TypeOf(v).Kind() == reflect.Int {
		fmt.Println(v)
	}
}

type MyType[T any] struct {
	Value T
}
