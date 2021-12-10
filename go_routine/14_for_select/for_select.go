package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("init main:1")
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func(ch1 chan string) {
		for {
			ch1 <- "every 500 ms"
			time.Sleep(time.Millisecond * 500)
		}
	}(ch1)
	go func(ch2 chan string) {
		for {
			ch2 <- "every 2 secs"
			time.Sleep(time.Second * 2)
		}
	}(ch2)

	for {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}

}
