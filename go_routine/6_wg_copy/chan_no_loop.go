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
	fmt.Println("out:", &wg)
	go func(wgc *sync.WaitGroup) {
		count("goroutine")
		fmt.Println("in1:", &wgc)
		wgc.Done()
	}(wg)
	go func(wgc *sync.WaitGroup) {
		count("go channels")
		fmt.Println("in2:", &wgc)
		wgc.Done()
	}(wg)
	wg.Wait()
}

func count(in string) {
	for i := 0; i < 2; i++ {
		println(len(in), in)
		time.Sleep(time.Millisecond * 500)

	}
}
