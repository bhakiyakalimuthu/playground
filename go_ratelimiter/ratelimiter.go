package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*6)
	defer cancel()
	ratelimiter(ctx, func() {
		fmt.Println("running ratelimiter")
	})
}

func ratelimiter(ctx context.Context, limiter func()) {

	ticker := time.NewTicker(time.Second * 5)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("context cancelled")
			return

		case <-ticker.C:
			limiter()

		}

	}

}
