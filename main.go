package main

import (
	"fmt"
	"sync"
)

var counter int

func main() {
	var wg sync.WaitGroup
	wg.Add(5)

	// Запуск нескольких горутин, которые изменяют глобальную переменную без синхронизации
	for i := 0; i < 5; i++ {
		go func() {
			// Здесь происходит изменение глобальной переменной `counter` без блокировки
			counter++
			fmt.Println(counter) // Это может вызвать ошибку гонки, так как несколько горутин изменяют counter
			wg.Done()
		}()
	}

	wg.Wait()
}

func MaxInt(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
