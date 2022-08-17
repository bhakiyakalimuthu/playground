package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	srv := NewServer()
	ctx, cancel := context.WithCancel(context.Background())
	// Handle graceful shutdown
	done := make(chan struct{}, 1)
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-time.Tick(time.Second * 5)
		log.Printf("sending termination")
		shutdown <- syscall.SIGINT
	}()
	go func() {
		<-shutdown
		signal.Stop(shutdown)

		log.Printf("Shutting down....")
		cancel()
		close(done)
	}()
	// Start the server
	log.Printf("starting server....")
	log.Fatal("closing server....\n", srv.start(ctx))
	<-done
}

type server struct {
}

func NewServer() *server {
	return &server{}
}
func (s *server) start(ctx context.Context) error {
	x := make(chan int)
	go send(ctx, x)

	for y := range x {
		log.Printf("value of x %d", y)
	}
	return nil
}

func send(ctx context.Context, x chan int) {
	for {
		select {
		case <-ctx.Done():
			close(x)
			return
		default:
			<-time.Tick(time.Second * 1)
			x <- 10

		}
	}

}

//func (s *server) start(ctx context.Context) error {
//	x := make(chan int)
//	go send(x)
//	for {
//		select {
//		case y := <-x:
//			log.Printf("value of x %d", y)
//		case <-ctx.Done():
//			return ctx.Err()
//		}
//	}
//	return nil
//}
//
//func send(x chan int) {
//
//	for i := 0; i < 10; i++ {
//		<-time.Tick(time.Second * 1)
//		x <- i
//	}
//}
