package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	one()
}


func one() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		one_(ctx)
		wg.Done()
	}()
	wg.Wait()
	fmt.Printf("main exiting - reason: %v", ctx.Err())
}

func one_(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context is done")
			return
		default:
			fmt.Println("running default - before waiting")
			<-time.After(time.Second * 20)
			fmt.Println("running default - after waiting")
			return
		}
	}
}
