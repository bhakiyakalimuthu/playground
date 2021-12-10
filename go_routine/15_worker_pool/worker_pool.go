package main

import "fmt"

func main() {
	fmt.Println("init main:1")
	jobs := make(chan int, 100)
	result := make(chan int, 100)
	go worker(jobs, result)
	go worker(jobs, result)
	go worker(jobs, result)
	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)
	for res := range result {
		fmt.Println(res)
	}
}

func worker(jobs, result chan int) {
	for job := range jobs {
		result <- fib(job)
	}
	close(result)
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)

}
