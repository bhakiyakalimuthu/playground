package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Job 1 is executed only once per hour at minute 17 (HH:17)
// Job 2 is executed every 4th minute. (HH:04, HH:08, HH:12, etc)
// Job 3 is executed every 6th minute with a 1-minute offset. (HH:07, HH:13,...)

func main() {
	fmt.Println("vim-go")

	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		ticker := time.NewTicker(time.Hour)
		for {
			select {
			case <-ticker.C:
				now := time.Now()
				if now.Minute() == int(17) {
					fmt.Println("job 1 running, current minute", now.Minute())
				}
			case <-ctx.Done():
				fmt.Println("shutting down job1")
				return
			}
		}
	}(ctx)
	go func(ctx context.Context) {
		ticker := time.NewTicker(time.Minute * 4)
		for {
			select {
			case <-ticker.C:
				now := time.Now()
				if now.Minute()%4 == 0 {
					fmt.Println("job 2 running")
				}
			case <-ctx.Done():
				fmt.Println("shutting down job2")
				return
			}
		}
	}(ctx)
	go func(ctx context.Context) {
		ticker := time.NewTicker(time.Minute * 6)
		for {
			select {
			case <-ticker.C:
				now := time.Now()
				if now.Minute()%6 == 0 {
					<-time.After(time.Minute * 1)
					fmt.Println("job 3 running")
				}
			case <-ctx.Done():
				fmt.Println("shutting down job3")
				return
			}
		}
	}(ctx)
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGTERM, syscall.SIGINT)
	<-done
	cancel()
}

func main_alternate() {
	var job1_last_executed, job2_last_executed, job3_last_executed time.Time

	for {
		now := time.Now()

		if now.Minute() == 17 && job1_last_executed.Before(now.Truncate(time.Hour)) {
			fmt.Println("Job 1 executed at", now)
			job1_last_executed = now
			continue
		}

		if now.Minute()%4 == 0 && job2_last_executed.Before(now.Truncate(time.Minute)) {
			fmt.Println("Job 2 executed at", now)
			job2_last_executed = now
			continue
		}

		minutes_since_last_exec := int((now.Sub(job3_last_executed)) / time.Minute)
		if minutes_since_last_exec >= 6 {
			exec_time := now.Truncate(time.Minute).Add(time.Minute)
			fmt.Println("Job 3 executed at", exec_time)
			job3_last_executed = exec_time
			continue
		}

		time.Sleep(1 * time.Second)
	}
}
