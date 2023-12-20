package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var result int64
	nums := []int64{2, 4, 6, 8, 10}

	for _, num := range nums {
		wg.Add(1)
		go func(wg *sync.WaitGroup, num int64) {
			defer wg.Done()
			atomic.AddInt64(&result, num*num)
		}(&wg, num)
	}

	wg.Wait()
	fmt.Println("Result	", result)
}
