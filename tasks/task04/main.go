package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	mainChan := make(chan int)

	var nWorkers int
	fmt.Print("Введите количество воркеров: ")
	fmt.Scan(&nWorkers)

	info := make(chan os.Signal, 1)
	signal.Notify(info, syscall.SIGINT, os.Interrupt)

	for i := 1; i <= nWorkers; i++ {
		go woker(mainChan, i)
	}

	for {
		select {
		case <-info:
			close(mainChan)
			close(info)
			return
		default:
			mainChan <- rand.Intn(100)
		}
	}
}

func woker(ch <-chan int, wid int) {
	for v := range ch {
		fmt.Printf("Получено число %d из канала от воркера %d\n", v, wid)
		time.Sleep(2 * time.Second)
	}
}
