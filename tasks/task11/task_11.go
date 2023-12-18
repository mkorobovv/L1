package task11

func Intersection(firstSet []int, secondSet []int) []int {

	hashTable := make(map[int]bool)
	for _, elem := range firstSet {
		hashTable[elem] = true
	}
	var resulSet []int
	for _, elem := range secondSet {
		if _, ok := hashTable[elem]; ok {
			resulSet = append(resulSet, elem)
		}
	}
	return resulSet
}
