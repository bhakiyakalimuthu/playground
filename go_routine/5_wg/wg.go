package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("init main:1")
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func() {
		count("goroutine")
		wg.Done()
	}()
	go func() {
		count("go channels")
		wg.Done()
	}()
	wg.Wait()
}

func count(in string) {
	for i := 0; i < 5; i++ {
		println(len(in), in)
		time.Sleep(time.Millisecond * 500)

	}
}
