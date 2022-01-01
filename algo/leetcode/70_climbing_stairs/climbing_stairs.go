package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func findWays1(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	return findWays(n-1) + findWays(n-2)

}

func findWays(n int) int {
	m := map[int]int{
		0: 0,
		1: 1,
		2: 2,
	}
	return dp(n, m)
}

func dp(n int, cache map[int]int) int {
	if v, ok := cache[n]; ok {
		return v
	}
	cache[n] = cache[dp(n-1, cache)] + cache[dp(n-2, cache)]
	return cache[n]
}
