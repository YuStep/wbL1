package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	// Первый вариант нет сохранения порядка
	//arr := [5]int{2, 4, 6, 8, 10}
	//
	//wg := sync.WaitGroup{}
	//ch := make(chan int, 5)
	//for _, s := range arr {
	//
	//	wg.Add(1)
	//	go func(val int) {
	//		defer wg.Done()
	//		ch <- val * val
	//	}(s)
	//}
	//wg.Wait()
	//close(ch)
	//for s := range ch {
	//	fmt.Println(s)
	//}

	array := [5]int{2, 4, 6, 8, 10}
	var wg = sync.WaitGroup{}
	for _, v := range array {
		wg.Add(1)
		go increaseAndPrint(float64(v), &wg)
	}
	wg.Wait()
}

func increaseAndPrint(num float64, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%d: %f | ", int(num), math.Pow(float64(num), 2))
}
