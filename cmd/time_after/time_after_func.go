package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	r := Resubmitter{}
	go r.newTask(6*time.Second, time.Second, func() error {
		fmt.Println("hellow world")
		return nil
	})
	<-time.After(time.Minute)
}

type Resubmitter struct {
	mu     sync.Mutex
	cancel context.CancelFunc
}

func (r *Resubmitter) newTask(repeatFor time.Duration, interval time.Duration, fn func() error) error {
	repeatUntilCh := time.After(repeatFor)

	r.mu.Lock()
	if r.cancel != nil {
		r.cancel()
	}
	ctx, cancel := context.WithCancel(context.Background())
	r.cancel = cancel
	r.mu.Unlock()

	firstRunErr := fn()

	go func() {
		for ctx.Err() == nil {
			select {
			case <-ctx.Done():
				return
			case <-repeatUntilCh:
				cancel()
				return
			case <-time.After(interval):
				fn()

			}
		}
	}()

	return firstRunErr
}
