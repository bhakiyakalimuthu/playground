package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	two()
}

func two() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	go func() {
		fmt.Println("running")
		<-time.After(time.Second * 1)
	}()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("calling done")
			switch ctx.Err() {
			case context.DeadlineExceeded:
				fmt.Println("deadline exceeded")
				return
			case context.Canceled:
				fmt.Println("canceled")
				return
			}
		}
	}
}
