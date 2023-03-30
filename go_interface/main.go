package main

import "fmt"

type i interface {
	hello(in string)
	nihao(in string)
}

var _ i = (*s)(nil)

type s struct{}

func (s s) nihao(in string) {
	//TODO implement me
	panic("implement me")
}

//func (s s) nihao(in string) {
//	//TODO implement me
//	panic("implement me")
//}

func (s *s) hello(in string) { fmt.Println(in) }
func main() {
	x := s{}
	fmt.Println(x.hello)
}
