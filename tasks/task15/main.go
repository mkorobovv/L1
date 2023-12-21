package main

import (
	"fmt"
	"strings"
)

func createHugeString(size int) string {
	var str strings.Builder

	for i := 0; i < size; i++ {
		fmt.Fprint(&str, "a")
	}

	return str.String()
}

var justString string

func someFunc() {
	// Главной проблемой прошлой программы было то, что мы срезаем не по количеству символов, а по количеству байт
	v := createHugeString(1 << 10)
	// Поэтому преобразуем строку в слайс рун, в таком случае мы срежем по количеству рун
	r := []rune(v)
	justString = string(r[:100])
}

func main() {
	someFunc()
}
