package task20

import "strings"

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
