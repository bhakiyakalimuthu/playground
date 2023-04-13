package main

import (
	"bytes"
	"fmt"
	"io"
)

type BufWriter struct {
	buf *bytes.Buffer
	io.Closer
}

func NewBufWriter() *BufWriter {

	buf := new(bytes.Buffer)
	return &BufWriter{
		buf: buf,
	}

}

func (b *BufWriter) Close() error {
	fmt.Println("closing")
	return nil
}
func WriteTo(w *BufWriter, msg []byte) error {

	defer w.Close()
	n, err := w.buf.Write(msg)
	if err != nil {
		return err
	}
	fmt.Println("bytes written", n)
	return nil
}

func main() {
	buf := NewBufWriter()
	if err := WriteTo(buf, []byte("hello world")); err != nil {
		panic(err)
	}
	fmt.Println(buf.buf.String())
}
