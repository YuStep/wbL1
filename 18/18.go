package main

import (
	"fmt"
	"time"
)

type inc struct {
	count int
}

func (i *inc) inc() {
	i.count += 1
	fmt.Println("new count:", i.count)
}

func main() {

	signal := make(chan bool)

	var counter inc
	counter.count = 0

	go func() {
		for {
			select {
			case <-signal:
				return
			default:
				counter.inc()
				time.Sleep(2 * time.Second)
			}
		}
	}()

	<-time.After(10 * time.Second)
	signal <- true
	fmt.Println("program ended")
}
