package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type AtomicCounter struct {
	counter int64
}

func NewAtomicCounter() *AtomicCounter {
	return &AtomicCounter{
		counter: 0,
	}
}

func (c *AtomicCounter) get() int64 { // Создадим метод AtomicCounter get
	return c.counter
}

func (c *AtomicCounter) add() { // Создадим метод AtomicCounter add
	atomic.AddInt64(&c.counter, 1)
}

func startWorker(iters int, c *AtomicCounter, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < iters; i++ {
		c.add()
	}
}

func main() {
	atomicCounter := NewAtomicCounter()
	var workersNum = 20
	var wg sync.WaitGroup
	var iters = 100
	wg.Add(workersNum)
	for i := 0; i < workersNum; i++ {
		go startWorker(iters, atomicCounter, &wg)
	}

	wg.Wait()
	fmt.Println("Counted ", atomicCounter.get(), " iters") // Итоговое количество итераций счетчика = workersNum * iters; 20 * 100 = 2000
}
