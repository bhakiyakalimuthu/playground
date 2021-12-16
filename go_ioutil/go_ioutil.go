package main

import (
	"fmt"
	"io/ioutil"

func main() {
	ioutil.ReadFile("testfile.txt")
	fmt.Println("vim-go")
}
