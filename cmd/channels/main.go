package main

import "fmt"

func main() {

	// nil channel
	var x chan string
	fmt.Println(x)

	// hchan struct is initialized to default
	y := make(chan int)
	fmt.Println(y)
}
