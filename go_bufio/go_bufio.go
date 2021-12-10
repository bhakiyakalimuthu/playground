package main

import (
	"bufio"
	"fmt"
)

func main() {
	bufferCheck()
}

type writer int

func (w writer) Write(p []byte) (int, error) {
	fmt.Println(len(p))
	return len(p), nil
}
func bufferCheck() {
	//	Unbuffered IO
	fmt.Println("Unbuffered IO")
	w := new(writer)
	w.Write([]byte(`a`))
	w.Write([]byte(`b`))
	w.Write([]byte(`c`))
	w.Write([]byte(`d`))

	// Buffered IO
	fmt.Println("Buffered IO")
	bf := bufio.NewWriterSize(w, 3)
	bf.Write([]byte(`a`))
	bf.Write([]byte(`b`))
	bf.Write([]byte(`c`))
	bf.Write([]byte(`d`))
	if err := bf.Flush(); err != nil {
		fmt.Println(err)
	}

	// Large write
	fmt.Println("buffered large write")
	bfl := bufio.NewWriterSize(w, 3)
	bfl.Write([]byte(`abcd`))

}
