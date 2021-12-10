package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("init main:1")
	count("goroutine")
	count("go channels")
}

func count(in string) {
	for i := 0; true; i++ {
		println(len(in), in)
		time.Sleep(time.Millisecond * 500)

	}
}
