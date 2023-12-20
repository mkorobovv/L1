package main

import "fmt"

func wahttype(unknown interface{}) {
	fmt.Printf("%T\n", unknown)
}

func main() {
	var i int
	var b bool
	var s string
	var f float32
	wahttype(i)
	wahttype(b)
	wahttype(s)
	wahttype(f)
}
