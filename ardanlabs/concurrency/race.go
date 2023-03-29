package main

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"
)

var counter int
var counter32 int32

func main() {
	//example()
	//useAtomic()
	useMutex()
}

func example() {
	const grs = 2
	wg := new(sync.WaitGroup)
	wg.Add(grs)

	for i := 0; i < grs; i++ {
		go func() {
			for i := 0; i < grs; i++ {
				value := counter
				value++
				log.Println(counter)
				counter = value
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("final counter %d", counter)
}

func useAtomic() {
	const grs = 2
	wg := new(sync.WaitGroup)
	wg.Add(grs)

	for i := 0; i < grs; i++ {
		go func() {
			for i := 0; i < grs; i++ {
				atomic.AddInt32(&counter32, 1)
				log.Println(counter32)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("final counter %d", counter32)
}
func useMutex() {
	const grs = 2
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)
	wg.Add(grs)

	for i := 0; i < grs; i++ {
		go func() {
			for i := 0; i < grs; i++ {
				mu.Lock()
				value := counter
				value++
				counter = value
				mu.Unlock()
				log.Println(counter)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("final counter %d", counter)
}
