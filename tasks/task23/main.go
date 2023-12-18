package main

import "fmt"

func RemoveElemmentByIndex(array []int, index int) []int {
	array = append(array[0:index], array[index+1:]...)
	return array
}

func main() {
	var array = []int{1, 2, 3, 4, 5}
	var i = 3
	fmt.Println(RemoveElemmentByIndex(array, i))
}
