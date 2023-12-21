package main

import (
	"fmt"
	"reflect"
)

func RecognizeType(unknown interface{}) reflect.Value {
	return reflect.ValueOf(unknown)
}

func main() {
	arrOfTypes := []interface{}{'a', "hello", 5, 0.23, true, struct{}{}}
	for _, typ := range arrOfTypes {
		typ := RecognizeType(typ)
		fmt.Printf("Value: %v, type: %s\n", typ, typ.Kind().String())
	}
}
