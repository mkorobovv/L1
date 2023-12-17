package task19

func ReverseLetters(str string) string {
	// Так как тип strin неизменяемый, создадим массив рун из строки
	var resultString = []rune(str)
	// Будем использовать метод двух указателей на начало и конец массива
	left, right := 0, len(resultString)-1
	for left < right {
		// Меняем местами конец и начало
		resultString[left], resultString[right] = resultString[right], resultString[left]
		// Двигаем указатели ближе к середине
		left++
		right--
	}
	/*
		Здесь в string кастуется срез (слайс), созданный из массива resultString.
		То есть фактически даже новая память не выделяется, в то время как в функции
		ToString память выделяется под новый срез (слайс), который по сути, того же размера что и исходный массив.
	*/
	return string(resultString[:])
}
