package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("init main:1")
	ch := make(chan string)
	go count("go channels", ch)
	for {
		msg := <-ch
		fmt.Println(msg)
	}
}

func count(in string, ch chan string) {
	for i := 0; i < 2; i++ {
		ch <- in
		time.Sleep(time.Millisecond * 500)

	}
}
