package main

import "net/http"

func main() {
	// http handle take in handlerfunc which implements http.ServeHTTP()
	http.Handle("/h1", http.HandlerFunc(handlerFunc))
	// http handleFunc  which implements http.ServeHTTP()
	http.HandleFunc("/h2", handlerFunc)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`hello handlerfunc`))
}
