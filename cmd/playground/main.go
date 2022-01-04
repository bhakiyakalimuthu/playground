package main

import "fmt"

func main() {

	s := []int{1, 2, 3, 4}

	s = append(s[:2], s[1:]...)
	s[2] = 6
	fmt.Print(s)
}
