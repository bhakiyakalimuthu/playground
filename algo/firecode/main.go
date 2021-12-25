package main

import "fmt"

func main() {
	fmt.Printf("fib %d\n", fib(0))
	fmt.Printf("fib %d\n", fib(1))
	fmt.Printf("fib %d\n", fib(2))
	fmt.Printf("fib %d\n", fib(50))
}

func fibSlow(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func fib(n int) int {
	cache := map[int]int{
		0: 0,
		1: 1,
	}
	return compute(n, cache)
}
func compute(n int, cache map[int]int) int {
	if v, ok := cache[n]; ok {
		return v
	}
	cache[n] = compute(n-1, cache) + compute(n-2, cache)
	return cache[n]
}
