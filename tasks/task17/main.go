package main

import "fmt"

func BinarySearch(arr []int, val int) int {
	left, right := 0, len(arr)
	mid := (left + right) / 2

	for left < right {
		if arr[mid] == val {
			return mid
		} else if arr[mid] < val {
			left = mid
		} else {
			right = mid
		}
		mid = (left + right) / 2
	}
	return -1
}

func main() {

	var sortedArray = []int{3, 5, 7, 10, 22, 44, 55, 60, 100, 234, 245, 289, 321}
	var searchNumber = 44
	fmt.Println("Search Number Index: ", BinarySearch(sortedArray, searchNumber))
}
