package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	alp = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	z   = 1 << iota
)

func main() {
	var a uint8 = 0b10101100
	var b uint8 = 0b11110000
	a &= b
	fmt.Printf("%b\n %d\n", a, a)
	a |= b
	fmt.Printf("%b\n %d\n", a, a)

	var x int16 = 100
	fmt.Printf("%016b\n", x<<1)
	var y int32 = 300
	fmt.Printf("%032b\n", y>>1)
	v := ^uint(0)
	fmt.Printf("v : %d\n", v)
	uintSize := 32 << (^uint(0) >> 32 & 1)
	fmt.Printf("%d\n", uintSize)

	fmt.Printf("z : %d\n", z)

	letterIdxBits := 10 // 6 bits to represent a letter index
	letterIdxMask := 2 << (letterIdxBits - 1)
	fmt.Printf("letterIdxMask : %d\n", letterIdxMask)
	byteArr := make([]byte, 3)
	for i, _ := range byteArr {
		byteArr[i] = alp[rand.Intn(len(alp))]
	}
	rand.Seed(time.Now().UnixNano())
	partOne := fmt.Sprintf("%v\n", string(byteArr))
	fmt.Printf(partOne)

	fmt.Printf("lef shift %d\n", 4<<10)
	fmt.Printf("lef shift %d\n", 1<<12)
}
