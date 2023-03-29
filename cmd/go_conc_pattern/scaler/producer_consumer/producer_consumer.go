package main

import (
	"flag"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

func main() {
	cpuprofile := flag.String("cpuprofile", "", "profile file path")
	memprofile := flag.String("memprofile", "", "profile file path")
	msgCount := flag.Int("msgcount", 5, "message count")
	dataChan := make(chan int)
	doneChan := make(chan bool)
	producer := NewProducer(dataChan, doneChan)
	consumer := NewConsumer(dataChan)

	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())
	if *cpuprofile != "" {
		file, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatalf("unable to create file %v", err)
		}
		defer pprof.StopCPUProfile()
		if err := pprof.StartCPUProfile(file); err != nil {
			log.Fatalf("failed to start cpu profile %v", err)
		}
	}

	go producer.produce(*msgCount)
	go consumer.consume()
	<-doneChan
	if *memprofile != "" {
		file, err := os.Create(*memprofile)
		if err != nil {
			log.Fatalf("failed create file %v", err)
		}
		runtime.GC()
		if err := pprof.WriteHeapProfile(file); err != nil {
			log.Fatalf("failed to start mem profile %v", err)
		}
	}
}

type consumer struct{ dataChan chan int }

func NewConsumer(dataChan chan int) *consumer {

	return &consumer{dataChan: dataChan}
}

func (c *consumer) consume() {
	for data := range c.dataChan {
		log.Printf("consumed message %d", data)
	}
}

type producer struct {
	dataChan chan int
	doneChan chan bool
}

func NewProducer(dataChan chan int, doneChan chan bool) *producer {
	return &producer{dataChan: dataChan, doneChan: doneChan}
}

func (p *producer) produce(count int) {
	for i := 0; i < count; i++ {
		p.dataChan <- i
	}
	p.doneChan <- true
}
