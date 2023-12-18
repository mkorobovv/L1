package main

import "fmt"

func Intersection(firstSet []int, secondSet []int) []int {

	hashTable := make(map[int]bool)
	for _, elem := range firstSet {
		hashTable[elem] = true
	}
	var resulSet []int
	for _, elem := range secondSet {
		if _, ok := hashTable[elem]; ok {
			resulSet = append(resulSet, elem)
		}
	}
	return resulSet
}

func main() {
	setOne := []int{1, 2, 10, -30, 4, 6}
	setTwo := []int{4, 0, -10, -30, 1, 3}
	fmt.Println(Intersection(setOne, setTwo))
}
