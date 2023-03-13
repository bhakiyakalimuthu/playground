package main

import (
	"fmt"
)

func main() {
	var s1 []string
	s2 := []string{"y"}
	s3 := make([]string, 1, 2)
	s3[0] = "x"
	s4 := make([]string, 0)
	s5 := make([]string, 5)

	fmt.Printf("s1 equal nil : %t len: %d cap : %d\n", s1 == nil, len(s1), cap(s1))
	fmt.Printf("s2 equal nil : %t len: %d cap : %d\n", s2 == nil, len(s2), cap(s2))
	fmt.Printf("s3 equal nil : %t len: %d cap : %d\n", s3 == nil, len(s3), cap(s3))
	fmt.Printf("s4 equal nil : %t len: %d cap : %d\n", s4 == nil, len(s4), cap(s4))
	fmt.Printf("s5 equal nil : %t len: %d cap : %d\n", s5 == nil, len(s5), cap(s5))
	fmt.Printf(s3[0])
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
	s = s[:]
	fmt.Println(s)

	list1 := []int{1, 2, 3}
	list1 = append(list1, 0)
	fmt.Println(list1)
	fmt.Println(list1[len(list1)-1])
	fmt.Println(int(^uint(0) >> 1))

	i := []int{1, 5, 9, 17, 18, 25, 60, 99}
	fmt.Printf("slice[:2] -> %d\n", i[:2])
	fmt.Printf("slice[5:] -> %d\n", i[5:])
	i = append(i[:3], i[2:]...)
	fmt.Printf("merge slice -> %v\n", i)

	in := make([]int, 0)
	in = append(in, 2)
	fmt.Printf(" before update in value : %v\n in address : %p\n", in, in)
	x(in)
	in = append(in, 3)
	fmt.Printf(" afer update in value : %v\n in address : %p\n", in, in)
}

func x(s []int) {

	fmt.Printf(" before update s value : %v\n s address : %p\n", s, s)
	//s = append(s, 5)
	//s1 := make([]int, 0)
	//s1 = append(s1, 1)
	//noOfItemCopied := copy(s, s1)
	//fmt.Printf(" noOfItemCopied  : %d\n", noOfItemCopied)
	//s2 := make([]int, 0)
	//s2 = append(s2, 3)
	//s = append(s, s2...)
	fmt.Printf(" after update s value : %v\n s address : %p\n", s, s)
}
