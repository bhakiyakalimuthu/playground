package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	x := -5
	fmt.Println(reflect.TypeOf(x))

	y := uint(x)
	fmt.Println(reflect.TypeOf(y))
	fmt.Println(y)

	s := "bhaki"
	fmt.Printf("s : %s\n", s[1:2])

	i, err := strconv.ParseInt("+10", 0, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(i)

	j, err := strconv.Atoi("-6")
	if err != nil {
		panic(err)
	}
	fmt.Println(j)

}
