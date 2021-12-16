package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("100kb.txt")
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(file)
	b := make([]byte, 1024)
	for {
		defer file.Close()
		_, err := r.Read(b)
		if err != nil {
			if err == io.EOF {
				fmt.Println("EOF")
				return
			}
			panic(err)

		}
		fmt.Println(string(b))

	}

}
