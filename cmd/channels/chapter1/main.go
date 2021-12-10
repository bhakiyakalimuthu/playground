package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go send(ch)
	go receive(ch)

	time.Sleep(1 * time.Second)
}

func send(ch chan int) {
	ch <- 1
}

func receive(ch chan int) {
	val := <-ch
	fmt.Println(val)
}
