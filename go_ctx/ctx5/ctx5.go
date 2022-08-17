package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	select {
	case <-ctx.Done():
		fmt.Println("context done")
	case <-time.After(4 * time.Second):
		fmt.Println("waiting")

	}
	time.Sleep(2 * time.Second)
	queueName := "yoyo"
	
}
