package task10

import "math"

func MakeTemepratureGroups(sequence []float64) map[int][]float64 {
	resultTable := make(map[int][]float64)

	for _, val := range sequence {
		group := (int(math.Floor(val)) / 10) * 10
		resultTable[group] = append(resultTable[group], val)
	}

	return resultTable
}
