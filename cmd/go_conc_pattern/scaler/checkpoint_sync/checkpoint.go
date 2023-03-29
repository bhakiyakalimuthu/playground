package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	noOfWorkers = 3
)

func main() {
	parts := []string{"A", "B", "C", "D", "F"}
	wg := new(sync.WaitGroup)

	for i := 1; i <= noOfWorkers; i++ {
		wg.Add(len(parts))

		for _, item := range parts {
			go partsCheck(item, wg)
		}
		wg.Wait()
		fmt.Printf("\nparts verification done by worker %d\n", i)
	}

}

func partsCheck(partName string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("\nstarting Part check : %s\n", partName)
	time.Sleep(time.Millisecond * 100)
	fmt.Printf("\nfinishing Part check : %s\n", partName)

}
