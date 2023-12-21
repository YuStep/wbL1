package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	mainChanel := make(chan int)
	const N = 20
	go worker(mainChanel)
	go worker1(mainChanel)
	select {
	case <-time.After(N * time.Second):
		close(mainChanel)
	}
}
func worker(mainChan chan int) {
	for {
		vr := rand.Intn(1000)
		mainChan <- vr
		fmt.Println("message sent: ", vr)
		time.Sleep(2 * time.Second)
	}
}
func worker1(mainChan chan int) {
	for {
		fmt.Println("message received: ", <-mainChan)
	}
}
