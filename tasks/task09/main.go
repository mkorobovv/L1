package main

import "fmt"

func Read(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range nums {
			out <- v
		}
		close(out)
	}()
	return out
}

func PowFromChan(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}

func main() {
	ch := Read(1, 2, 3, 4, 5)
	pow2nums := PowFromChan(ch)
	for v := range pow2nums {
		fmt.Print(v, " ")
	}
}
