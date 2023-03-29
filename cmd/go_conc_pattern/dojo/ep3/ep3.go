package main

import (
	"fmt"
	"time"
)

func main() {
	tChan := make(chan int, 1)

	go func() {

		fmt.Println(<-tChan)
	}()
	tChan <- 6
	tChan <- 6

	time.Sleep(time.Second * 5)

}
