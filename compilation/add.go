package main

import "fmt"

func main() {
	a := 1
	b := 2
	if true {
		add(a, b)
	}
}

func add(a, b int) {
	fmt.Println(a + b)
}
