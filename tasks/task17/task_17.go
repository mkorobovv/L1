package task17

func BinarySearch(arr []int, val int) (int, int) {
	left, right := 0, len(arr)
	mid := (left + right) / 2

	for left < right {
		if arr[mid] == val {
			return mid, val
		} else if arr[mid] < val {
			left = mid
		} else {
			right = mid
		}
		mid = (left + right) / 2
	}
	return -1, -1
}
