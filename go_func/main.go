package main

import (
	"fmt"
	"time"
)

type Search func(x string)

func main() {
	f1 := func(x string) {

		fmt.Printf("func f1 %s\n", x)
	}
	f2 := func(x string) {

		fmt.Printf("func f2 %s\n", x)
	}

	work("hello", f1, f2)
	<-time.After(time.Second * 2)
}

func work(query string, replica ...Search) {

	replicaSearch := func(i int) {
		replica[i](query)
	}

	for i := range replica {
		fmt.Println(i)
		go replicaSearch(i)

	}
}
