package main

import "fmt"

func main() {
	c := car{size: 10}
	b := bike{height: 20}
	convert(c)
	convert(b)
}

func convert(s fmt.Stringer) {
	fmt.Println(s.String())
}

type car struct {
	size int
}

func (c car) String() string {
	return fmt.Sprintf("car size : %d", c.size)
}

type bike struct {
	height int
}

func (b bike) String() string {
	return fmt.Sprintf("bike height : %d", b.height)
}
