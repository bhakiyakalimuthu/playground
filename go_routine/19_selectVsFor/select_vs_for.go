package main

import "fmt"

func main() {
	ch := make(chan string)
	count := 0
	for {
		if count == 10 {
			return
		}
		select {
		case <-ch:
		default:
			count++
			fmt.Println("default case executed")
		}
	}

}
