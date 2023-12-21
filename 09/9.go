package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаем каналы для передачи данных между горутинами
	inputCh := make(chan int)
	outputCh := make(chan int)

	// Используем WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Горутина для записи чисел в канал inputCh
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(inputCh)

		numbers := []int{1, 2, 3, 4, 5}

		for _, num := range numbers {
			inputCh <- num
		}
	}()

	// Горутина для умножения чисел на 2 и записи в канал outputCh
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(outputCh)

		for num := range inputCh {
			outputCh <- num * 2
		}
	}()

	// Горутина для вывода результатов из канала outputCh в stdout
	wg.Add(1)
	go func() {
		defer wg.Done()

		for result := range outputCh {
			fmt.Println(result)
		}
	}()

	// Ожидание завершения всех горутин
	wg.Wait()
}
