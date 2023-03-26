package main

import "fmt"

func main() {
	ex := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.14,
	}
	fmt.Printf("%+v\n", ex)
	fmt.Printf("%+v\n", &ex.flag)
	fmt.Printf("%+v\n", &ex.counter)
	fmt.Printf("%+v\n", &ex.pi)
	var ex1 example
	fmt.Println(ex1)
	var ex2 example1
	ex1 = example(ex2)
}

type example struct {
	flag    bool
	counter int16
	pi      float32
}
type example1 struct {
	flag    bool
	counter int16
	pi      float32
}
