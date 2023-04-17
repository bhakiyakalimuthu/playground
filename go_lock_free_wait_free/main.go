package main

import (
	"fmt"
	"sync/atomic"
)

// Lock free data structure
type counterLF struct {
	value int64
}

func (cl *counterLF) Inc() {
	atomic.AddInt64(&cl.value, 1)
}

func (cl *counterLF) Value() int64 {
	return atomic.LoadInt64(&cl.value)
}

// Wait free data structure
type counterWF struct {
	value atomic.Value
}

func (cw *counterWF) Inc() {
	cw.value.Store(cw.value.Load().(int) + 1)
}

func (cw *counterWF) Value() int {
	return cw.value.Load().(int)
}

func main() {
	var val int
	val = 10
	// lock free
	lf := &counterWF{}
	lf.Inc()
	lf.Inc()
	lf.Inc()
	fmt.Println(lf.Value(), "\n")

	// wait free
	wf := &counterWF{}
	wf.Inc()
	wf.Inc()
	wf.Inc()
	wf.Inc()
	fmt.Println(wf.Value())
}
