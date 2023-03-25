package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// simple check
	x := 6
	y := unsafe.Pointer(&x)
	z := &x
	fmt.Println(&y)
	fmt.Println(y)
	fmt.Println(&x)
	fmt.Println(z)

	fmt.Println("------------------done")
	// struct type casting

	_x := &myStructA{}
	_y := (*int)(unsafe.Pointer(_x))
	_z := &_x
	fmt.Println(&_x)
	fmt.Println(_z)
	fmt.Println(_y)

	// learn uintptr
	var unsignedPointerInt uintptr = 1024
}

type myStructA struct {
	a int
	b int
}

type myStructB struct {
	i string
	j string
}
