package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int, 2)
	ch1 <- 1
	ch1 <- 2
	fmt.Println("sending data to channel 1")
	val := <-ch1
	fmt.Printf("receiving from channel 1 val: %d\n", val)
	ch2 := make(chan int, 1)

	go send(ch2)
	go receive(ch2)

	time.Sleep(5 * time.Second)
}

func send(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("sending value to channel 2")
}

func receive(ch chan int) {
	time.Sleep(1 * time.Second)
	fmt.Println("timeout finished")
	val := <-ch
	fmt.Printf("receiving value from channel 2 val:%d", val)

}
