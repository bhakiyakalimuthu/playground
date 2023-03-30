package main

import (
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("seek.txt")
	if err != nil {
		log.Fatalf("failed to open")
	}

	cur, err := f.Seek(16, io.SeekStart)
	if err != nil {
		log.Fatalf("failed to seek %v", err)
	}
	log.Printf("\n current locaiton %d", cur)

	buf := make([]byte, 3)

	n, err := f.Read(buf)
	if err != nil {
		if err == io.EOF {
			log.Fatalf("reached end of file")
		}

		log.Fatalf("read failed %v", err)
	}
	log.Printf("string value %s", string(buf[:n]))

	cur, err = f.Seek(-6, io.SeekEnd)
	if err != nil {
		log.Fatalf("failed to seek %v", err)
	}
	log.Printf("\n current locaiton %d", cur)
	buf = make([]byte, 6)

	n, err = f.Read(buf)
	if err != nil {
		if err == io.EOF {
			log.Fatalf("reached end of file")
		}

		log.Fatalf("read failed %v", err)
	}
	log.Printf("string value %s", string(buf[:n]))
}
