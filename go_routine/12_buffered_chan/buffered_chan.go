package main

import (
	"fmt"
)

func main() {
	fmt.Println("init main:1")
	ch := make(chan string, 2)
	ch <- "hello"
	ch <- "hi"

	fmt.Println(<-ch, <-ch)
}
