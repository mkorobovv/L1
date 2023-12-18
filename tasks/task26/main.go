package main

import (
	"fmt"
	"strings"
)

func UniqueSymbols(word string) bool {
	cache := make(map[rune]struct{})
	lowerString := strings.ToLower(word)
	for _, v := range lowerString {

		if _, ok := cache[v]; ok {
			return false
		}
		cache[v] = struct{}{}
	}

	return true
}

func main() {
	letter1 := "abcdefghijklmnopqrstuvwxyz"
	fmt.Println(UniqueSymbols(letter1))
	letter2 := "aBbcde"
	fmt.Println(UniqueSymbols(letter2))
}
