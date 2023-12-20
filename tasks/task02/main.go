package main

import (
	"fmt"
	"sync"
)

func ConcurrencySquares(numbers []int, wg *sync.WaitGroup) {
	for _, v := range numbers {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			fmt.Println(v * v)

		}(v)
	}
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	ConcurrencySquares(numbers, &wg)
	wg.Wait()
}
