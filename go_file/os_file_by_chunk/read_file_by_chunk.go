package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("test1Mb.db")
	if err != nil {
		panic(err)
	}
	chunk := make([]byte, 4)
	for {
		read, err := file.Read(chunk)
		if err != nil {
			if err == io.EOF {
				fmt.Println(err)
			}
			break
		}

		fmt.Println(string(chunk[:read]))
	}
}
