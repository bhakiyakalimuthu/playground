package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		fmt.Fprintf(os.Stdout, "received request\n")
		select {
		case <-time.After(time.Second * 3):
			w.Write([]byte(`OK`))
		case <-ctx.Done():
			fmt.Fprintf(os.Stdout, "request cancelled\n")
		}
	}))
}
