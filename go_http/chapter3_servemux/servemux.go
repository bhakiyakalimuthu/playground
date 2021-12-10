package main

import (
	"fmt"
	"net/http"
)

func main() {
	// serveMuxSvc
	mux := http.NewServeMux()
	handler := serveMuxSvc{}
	mux.Handle("/h1", handler)

	// serveMuxFunc to use httpHandlerFunc
	handlerFunc := http.HandlerFunc(serveMuxFunc)
	mux.Handle("/h2", handlerFunc)

	// serveMuxHandleFunc
	mux.HandleFunc("/h3", serveMuxFunc)

	if err := http.ListenAndServe(":8082", mux); err != nil {
		panic(err)
	}
}

type serveMuxSvc struct{}

func (serveMuxSvc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("serving http request - serveMuxSvc")
	w.Write([]byte(`welcome to serveMuxSvc`))
}

func serveMuxFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("serving http request - serveMuxFunc")
	w.Write([]byte(`welcome to serveMuxFunc`))

}
