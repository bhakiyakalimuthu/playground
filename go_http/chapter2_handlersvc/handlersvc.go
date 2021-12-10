package main

import (
	"fmt"
	"net/http"
)

func main() {
	// http handler --> Handler interface --> ServerHttp(http.ResponseWriter, *http.Request) method
	// ResponseWriter() interface --> header() method --> Write([]byte) method --> WriterHeader(statusCode int)

	// http handler

	// http handler func

	// http server

	// http server mux

	handler := &HandlerSvc{}
	if err := http.ListenAndServe(":8080", handler); err != nil {
		panic(err)
	}

}

type HandlerSvc struct{}

func (h *HandlerSvc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("serving http request")
	w.Write([]byte(`hello welcome to handlersvc`))
}
