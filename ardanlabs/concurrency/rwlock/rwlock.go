package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var data []string
var rwMu sync.RWMutex

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func() {
		for i := 0; i < 10; i++ {

			Writer(i)
		}
		wg.Done()
		fmt.Println("\nwriter completed")
	}()
	for i := 0; i < 8; i++ {
		go func() {
			for {
				Reader()
			}

			fmt.Println("\nreader completed")

		}()
	}
	wg.Wait()

}

func Reader() {
	rwMu.RLock()
	<-time.After(time.Millisecond * time.Duration(rand.Intn(100)))
	fmt.Printf("\nReader reading data : %v", data)
	rwMu.RUnlock()
}
func Writer(val int) {
	rwMu.Lock()
	<-time.After(time.Millisecond * time.Duration(rand.Intn(100)))
	fmt.Printf("\nwriter writing data <----- %d", val)
	data = append(data, fmt.Sprintf("%d", val))
	rwMu.Unlock()

}
