package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	dataCh := make(chan []byte)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go produce(wg, dataCh)
	go consume(wg, dataCh)
	wg.Wait()

}

func produce(wg *sync.WaitGroup, dataCh chan []byte) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		dataCh <- []byte(fmt.Sprintf("yoyo_%d ", i))
	}
	<-time.After(time.Second)
	close(dataCh)
}

func consume(wg *sync.WaitGroup, dataCh chan []byte) error {
	file, err := os.OpenFile("./test.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("failed to open a file, reason:", err)
		return err
	}
	writeToFile(wg, file, dataCh)
	return nil
}

func writeToFile(wg *sync.WaitGroup, file *os.File, dataCh chan []byte) {
	defer file.Close()
	for data := range dataCh {
		fmt.Printf("writing data :%s", string(data))
		_, err := file.Write(data)
		if err != nil {
			fmt.Println(err)
		}
	}
}
