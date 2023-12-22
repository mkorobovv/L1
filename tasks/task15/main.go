package main

import (
	"fmt"
	"strings"
)

func createHugeString(size int) string {
	var str strings.Builder
	// Создадим большую строку с помощью стринг билдера
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
	byteSlice := make([]byte, 100)

	copy(byteSlice, v[:100])
	justString := string(byteSlice)
	v = ""
	fmt.Println(justString)
}

func main() {
	someFunc()
}
