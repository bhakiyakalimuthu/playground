package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	timeout := time.Millisecond * 150
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	wChan := make(chan string, 2)
	go func(wchan chan string) {
		rand.Seed(time.Now().UnixNano())
		timeout := time.Millisecond * time.Duration(rand.Intn(200))
		<-time.After(timeout)
		fmt.Println(timeout)
		wchan <- "hello"

	}(wChan)

	fmt.Println("before select")
	select {
	case data := <-wChan:
		fmt.Printf("received data %s", data)
	case <-ctx.Done():
		fmt.Printf("received cancellation")
	}

}
