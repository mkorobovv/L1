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
	/*
		Поскольку в Go подстрока ссылается на оригинальную строку, оригинальная строка не будет освобождена из памяти до тех пор, пока подстрока существует.
		Это может привести к накоплению большого количества неиспользуемой памяти.
		Чтобы исправить это, можно скопировать нужную часть строки в новую переменную, чтобы избежать сохранения ссылки на оригинальную строку.
	*/
	v := createHugeString(1 << 10)
	justString := v[:100]
	v = ""
	fmt.Println(string(justString))
}

func main() {
	someFunc()
}
