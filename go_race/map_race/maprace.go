package main

import (
	"fmt"
	"sync"
	"time"
)

type RLockAndLockStruct struct {
	mu sync.RWMutex

	mapEx map[string]string
}

func main() {
	r := &RLockAndLockStruct{}
	r.mapEx = make(map[string]string)

	go r.RLockAndLockTest("test", "goroutine 1 - ")
	go r.RLockAndLockTest("test", "goroutine 2 - ")
	time.Sleep(4000 * time.Millisecond)
}

func (r *RLockAndLockStruct) RLockAndLockTest(value string, goroutine string) string {
	r.mu.Lock()
	defer r.mu.Unlock()
	fmt.Printf("%sRLock\n", goroutine)
	t := r.mapEx[value]
	fmt.Printf("%sRUnlock\n", goroutine)
	if len(t) <= 0 {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("%sLock\n", goroutine)
		r.mapEx[value] = value
		fmt.Printf("%sUnlock\n", goroutine)
		return r.mapEx[value]
	}
	return t
}
