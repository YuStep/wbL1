package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	var sum int

	for _, v := range arr {
		wg.Add(1)
		go increaseAndPrint(v, &wg, &sum)
	}
	wg.Wait()
	fmt.Printf("%d\n", sum)
}

func increaseAndPrint(num int, wg *sync.WaitGroup, sum *int) {
	defer wg.Done()
	*sum += num * num
}
