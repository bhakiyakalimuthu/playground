package main

import (
	"sync"
	"time"
)

//https://betterprogramming.pub/deep-dive-into-concurrency-of-go-93002344d37b

// set this env value  `GODEBUG=schedtrace=1000`

//gomaxprocs: Processors configured
//idleprocs: Processors are not in use. Goroutine running.
//threads: Threads in use.
//idlethreads: Threads are not in use.
//runqueue: Goroutines in the global queue.
//[1 0 0 0]: Goroutines in each processor's local run queue.

//SCHED 0ms: gomaxprocs=10 idleprocs=9 threads=2 spinningthreads=0 idlethreads=0 runqueue=0 [0 0 0 0 0 0 0 0 0 0]
//SCHED 1020ms: gomaxprocs=10 idleprocs=0 threads=11 spinningthreads=0 idlethreads=0 runqueue=1 [1 0 0 7 0 0 0 1 0 0]
//SCHED 2019ms: gomaxprocs=10 idleprocs=0 threads=11 spinningthreads=0 idlethreads=0 runqueue=8 [0 0 0 0 0 0 1 1 0 0]
//SCHED 3040ms: gomaxprocs=10 idleprocs=0 threads=11 spinningthreads=0 idlethreads=0 runqueue=9 [0 1 0 0 0 0 0 0 0 0]
//SCHED 4051ms: gomaxprocs=10 idleprocs=0 threads=11 spinningthreads=0 idlethreads=0 runqueue=9 [0 0 0 0 0 0 0 0 1 0]
//SCHED 5068ms: gomaxprocs=10 idleprocs=0 threads=11 spinningthreads=0 idlethreads=0 runqueue=8 [0 0 0 0 0 1 0 0 1 0]
//SCHED 6066ms: gomaxprocs=10 idleprocs=0 threads=11 spinningthreads=0 idlethreads=0 runqueue=8 [0 0 0 0 1 0 0 0 0 1]
//SCHED 7076ms: gomaxprocs=10 idleprocs=0 threads=11 spinningthreads=0 idlethreads=0 runqueue=8 [0 0 0 0 1 0 0 1 0 0]
//SCHED 8087ms: gomaxprocs=10 idleprocs=0 threads=11 spinningthreads=0 idlethreads=0 runqueue=8 [0 0 0 0 1 0 0 0 1 0]

var wg sync.WaitGroup

func main() {
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go work()
	}
	wg.Wait()
}

func work() {
	time.Sleep(time.Second)
	var counter int
	for i := 0; i < 1e10; i++ {
		counter++
	}
	wg.Done()
}
