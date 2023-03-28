package main

import "fmt"

func main() {
	in := []int{5, 4, 6, 8, 9, 1}

	s1Chan := stage1(in)

	s2Chan := stage2(s1Chan)

	for out := range s2Chan {

		fmt.Println(out)
	}
}

func stage1(in []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, val := range in {
			out <- val
		}
		close(out)
	}()
	return out
}

func stage2(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for val := range in {
			out <- val * val
		}
		close(out)
	}()
	return out
}
