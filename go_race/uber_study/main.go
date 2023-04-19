package main

import (
	"fmt"
	"time"
)

func main() {
	s := []int{1, 6, 0, 5, 8}
	for _, val := range s {
		//go func(in int) {
		//	process(in)
		//}(val)
		go process(val)
	}

	<-time.After(time.Second * 1)
}

func process(in int) {

	fmt.Printf("\nval: %d address: %v", in, &in)

}
