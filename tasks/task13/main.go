package main

import "fmt"

func Swap(a, b int) (int, int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return a, b
}

func main() {
	a, b := 10, 20
	fmt.Println(Swap(a, b))
}
