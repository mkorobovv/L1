package main

import (
	"fmt"
)

func RecognizeType(unknown interface{}) {
	switch unknown.(type) {
	case int:
		fmt.Println("int value")
	case bool:
		fmt.Println("bool value")
	case string:
		fmt.Println("string value")
	case chan int:
		fmt.Println("chan int value")
	default:
		fmt.Println("Unknown value")
	}
}

func main() {
	chanvar := make(chan int)
	RecognizeType(chanvar)

	intvar := 5
	RecognizeType(intvar)

	boolvar := false
	RecognizeType(boolvar)

	stringvar := "hello"
	RecognizeType(stringvar)
}
