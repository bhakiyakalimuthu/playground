package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/h1", http.HandlerFunc(handler))
	http.HandleFunc("/h2", handler)
	var router router
	http.ListenAndServe(":6666", &router)
	fmt.Println("vim-go")
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("welcome to my http handler"))
}

type router struct{}

func (router *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/h3":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("welcome to my h3 handler"))
	case "/h4":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("welcome to my h4 handler"))
	default:
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("welcome to my default handler"))

	}

}
