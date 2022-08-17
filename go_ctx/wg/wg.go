package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	ch := make(chan string)
	go func() {
		for {
			channel, ok := <-ch
			if ok {
				fmt.Println("Shut Down")
				defer wg.Done()
				return
			}
			fmt.Println(channel)
			<-time.After(time.Second)
		}
	}()
	ch <- "Start"
	ch <- "Processing"
	ch <- "Finishing"
	close(ch)

	wg.Wait()
}
