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
	mainChanel := make(chan int)
	const workerCount = 5

	for i := 0; i < workerCount; i += 1 {
		go worker(mainChanel, i)
	}
	info := make(chan os.Signal, 1)
	signal.Notify(info, syscall.SIGINT, os.Interrupt)
	for {
		select {
		case <-info:
			close(mainChanel)
			close(info)
			return
		default:
			mainChanel <- rand.Intn(100)
			time.Sleep(2 * time.Second)
		}
	}
}

func worker(ch chan int, workerId int) {
	for msg := range ch {
		fmt.Printf("mes: %d; WID: %d\n", msg, workerId)
	}
}
