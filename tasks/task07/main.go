package main

import (
	"fmt"
	"sync"
)

// Структура безопасной мапы
type SafeMap struct {
	mu   sync.RWMutex // RWMutex для безопасного чтения и записи в мапу
	data map[string]int
}

func NewMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]int),
	}
}

func (s *SafeMap) Set(key string, value int) {
	s.mu.Lock() // Блокируем запись в мапу для других горутин, горутина ждет завершения метода Set
	defer s.mu.Unlock()
	s.data[key] = value
}

func (s *SafeMap) Get(key string) int {
	s.mu.RLock() // Блокирует только чтение из мапы, при этом запись не заблокирована
	defer s.mu.RUnlock()
	return s.data[key]
}

func main() {
	var wg sync.WaitGroup // Создадим объект WaitGroup для ожидания выполнения горутин

	safemap := NewMap()
	stringelems := []string{"test", "hello", "world", "mutex", "gopher"}
	wg.Add(2)
	go func(wg *sync.WaitGroup) {
		for i, v := range stringelems {
			safemap.Set(v, i)
		}
		wg.Done()
	}(&wg)

	go func(wg *sync.WaitGroup) {
		for _, v := range stringelems {
			fmt.Println(safemap.Get(v))
		}
		wg.Done()
	}(&wg)

	wg.Wait() // Программа ждет выполнение всех горутин
}
