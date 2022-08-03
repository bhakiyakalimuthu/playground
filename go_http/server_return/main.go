package main

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	srv := NewServer()
	ctx, cancel := context.WithCancel(context.Background())
	// Handle graceful shutdown
	done := make(chan struct{}, 1)
	go func() {
		shutdown := make(chan os.Signal, 1)
		signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
		<-shutdown
		signal.Stop(shutdown)

		log.Printf("Shutting down....")
		srv.srv.Shutdown(ctx)
		cancel()
		close(done)
	}()
	// Start the server
	log.Printf("starting server....")
	log.Fatal(srv.start(ctx))
	<-done
}

type server struct {
	srv *http.Server
}

func NewServer() *server {
	return &server{}
}
func (s *server) start(ctx context.Context) error {

	r := mux.NewRouter()
	r.HandleFunc("/", health).Methods(http.MethodGet)
	s.srv = &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	err := s.srv.ListenAndServe()
	if err == http.ErrServerClosed {
		return nil
	}
	return err
}

func health(response http.ResponseWriter, r *http.Request) {
	response.Write([]byte(`ok`))
}
