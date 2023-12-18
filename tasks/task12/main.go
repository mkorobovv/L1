package main

import "fmt"

func MakeSetFromSequence(sequence []string) map[string]bool {
	set := make(map[string]bool)
	for _, elem := range sequence {
		set[elem] = true
	}
	return set
}

func main() {
	elements := []string{"cat", "dog", "bird", "dog", "dog", "cat"}
	fmt.Println(MakeSetFromSequence(elements))
}
