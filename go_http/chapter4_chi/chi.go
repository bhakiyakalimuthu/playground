package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Get("/h1", chiHandlerFunc)
	if err := http.ListenAndServe(":8083", r); err != nil {
		panic(err)
	}
}

func chiHandlerFunc(w http.ResponseWriter, r *http.Request) {

	fmt.Println("serving http request - chiHandlerFunc")
	w.Write([]byte(`welcome to chiHandlerFunc`))
}
