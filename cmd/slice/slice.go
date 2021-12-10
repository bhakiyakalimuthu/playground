package main

import (
	"fmt"
)

func main() {
	var s1 []string
	s2 := []string{"y"}
	s3 := make([]string, 0, 2)
	s4 := make([]string, 0)
	s5 := make([]string, 5)

	fmt.Printf("s1 equal nil : %t len: %d cap : %d\n", s1 == nil, len(s1), cap(s1))
	fmt.Printf("s2 equal nil : %t len: %d cap : %d\n", s2 == nil, len(s2), cap(s2))
	fmt.Printf("s3 equal nil : %t len: %d cap : %d\n", s3 == nil, len(s3), cap(s3))
	fmt.Printf("s4 equal nil : %t len: %d cap : %d\n", s4 == nil, len(s4), cap(s4))
	fmt.Printf("s5 equal nil : %t len: %d cap : %d\n", s5 == nil, len(s5), cap(s5))

	s := "playground"
	for _, v := range s {
		fmt.Println(string(v))
	}

	list := []int{1, 2, 3}
	l := list[1:2:2]
	fmt.Println(l)
	fmt.Printf("capacity : %d\n", cap(l))

	fmt.Printf("string first item : %s last item %s\n", string(s[0]), string(s[len(s)-1]))
	fmt.Printf("rune first item : %d last item %d\n", s[0], s[len(s)-1])
	s = s[:len(s)-1]
	fmt.Println(s)

}
