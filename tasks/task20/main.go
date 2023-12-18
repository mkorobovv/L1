package main

import (
	"fmt"
	"strings"
)

func ReverseSentence(s string) string {
	var arrayOfWords = strings.Split(s, " ")
	left, right := 0, len(arrayOfWords)-1

	for left < right {
		arrayOfWords[left], arrayOfWords[right] = arrayOfWords[right], arrayOfWords[left]
		left++
		right--
	}
	return strings.Join(arrayOfWords, " ")
}

func main() {
	var words = "sun dog snow"
	fmt.Println("Исходное предложение: ", words, "\nПосле исполнения функции: ", ReverseSentence(words))
}
