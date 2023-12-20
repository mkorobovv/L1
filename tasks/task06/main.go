package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func MethodChannels() {
	fmt.Println("Старт функции на каналах")

	done := make(chan struct{}) // Канал для сигнала завершения горутины

	go func() {
		defer fmt.Println("Горутина на каналах заверешена")
		defer close(done)
		for i := 0; i < 5; i++ {
			fmt.Println("Горутина на каналах:", i)
			time.Sleep(time.Second)
		}
	}()

	// Ждем сигнала завершения
	<-done
}

func MethodContext() {
	fmt.Println("Старт функции на контексте")
	ctx, cancel := context.WithCancel(context.Background())
	defer fmt.Println("Горутина на контексте завершена")
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println("Горутина на контексте работает...")
				time.Sleep(time.Second)
			}
		}
	}()

	fmt.Println("Нажмите Enter, чтобы остановить горутину на контексте.")
	fmt.Scanln()

}

func MethodWaitGroup() {
	fmt.Println("Старт функции Wait Group")
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Println("Горутина на Wait Group ", i)
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
	fmt.Println("\nГорутина на Wait Group завершена")
}

func main() {
	MethodChannels()
	fmt.Println()

	MethodContext()
	fmt.Println()

	MethodWaitGroup()
	fmt.Println()

	fmt.Println("Все Горутины завершили свои работы, программа завершена")

}
