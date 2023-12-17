package task23

func RemoveElemmentByIndex(array []int, index int) []int {
	array = append(array[0:index], array[index+1:]...)
	return array
}
