package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("100kb.txt")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}
	stat, err := file.Stat()
	fmt.Printf("%#v %v", stat, err)
}
