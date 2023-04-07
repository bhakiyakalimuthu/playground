package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Server struct {
	server *http.Server
	db     *Store
}

func NewServer() *Server {
	return &Server{
		db: NewStore(),
	}
}

func (s *Server) Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.HandleRequest)
	s.server = &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	s.server.ListenAndServe()
}

func (s *Server) Stop() {
	if s.server != nil {
		_ = s.server.Shutdown(context.Background())
	}
}

func (s *Server) HandleRequest(w http.ResponseWriter, r *http.Request) {
	_id := r.URL.Query().Get("id")
	fmt.Println(_id)
	id, err := strconv.Atoi(_id)
	if err != nil {

		WriteResponse(w, http.StatusInternalServerError, nil)
		return
	}

	//user, ok := s.db.TryCache(id)
	//if ok {
	//	json.NewEncoder(w).Encode(user)
	//	return
	//}
	user, ok := s.db.Get(id)
	if !ok {
		//WriteResponse(w, http.StatusNotFound, nil)
		panic("user not found")
		return
	}
	s.db.dbhit++
	//fmt.Println(user)
	//s.db.cache[id] = user
	json.NewEncoder(w).Encode(user)
}

func WriteResponse(w http.ResponseWriter, statusCode int, response []byte) {
	w.WriteHeader(statusCode)
	w.Write(response)

}

type User struct {
	ID   int
	Name string
}

type Store struct {
	db    map[int]*User
	dbhit int
	cache map[int]*User
}

func NewStore() *Store {
	db := make(map[int]*User)
	cache := make(map[int]*User)
	for i := 1; i <= 100; i++ {
		db[i] = &User{
			ID:   i,
			Name: fmt.Sprintf("name_%d", i),
		}
	}
	return &Store{db: db, cache: cache}
}

func (str *Store) Get(id int) (*User, bool) {
	if id == 999 {
		fmt.Println("inside")
	}
	v, ok := str.db[id]
	fmt.Printf("\n getting from db %v", v)
	return v, ok
}

func (str *Store) TryCache(id int) (*User, bool) {
	v, ok := str.cache[id]
	//fmt.Printf("\n getting from db %v", v)
	return v, ok
}

func main() {
	fmt.Println("vim-go")
}
