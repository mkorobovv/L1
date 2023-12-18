package main

import "fmt"

func MakeTemepratureGroups(sequence []float64) map[int][]float64 {
	resultTable := make(map[int][]float64)

	for _, val := range sequence {
		group := (int(val) / 10) * 10
		resultTable[group] = append(resultTable[group], val)
	}

	return resultTable
}

func main() {
	var temperatures = []float64{22.3, 33.7, -12.9, 13.8, -22.1, -23.5}
	fmt.Println(MakeTemepratureGroups(temperatures))
}
