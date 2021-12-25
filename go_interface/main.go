package main

import "fmt"

type i interface {
	hello()
}

type s struct{}

func (s *s) hello() { fmt.Println("hello") }
func main() {
	x := s{}
	fmt.Println(x.hello)
}
