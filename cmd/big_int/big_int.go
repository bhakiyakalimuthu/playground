package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	//var x uint64 = 7552855
	//y := big.NewInt(int64(x))
	////z := new(big.Int).SetUint64(x)
	//z := new(big.Int).SetUint64(0)
	//fmt.Println(big.NewInt(10))
	//fmt.Println(y)
	//fmt.Println(y.Cmp(z))
	bundleCh := make(chan []bundle)
	l := []uint64{1, 0, 2, 3}
	fmt.Println("l size", "len", len(l))
	bundles := make([]bundle, len(l))
	go check(bundleCh)
	for _, i := range l {
		if i == 0 {
			continue
		}
		b := bundle{
			x: i,
			y: new(big.Int).SetUint64(i),
		}
		bundles = append(bundles, b)
	}
	fmt.Println("bundle size", "len", len(bundles))
	bundleCh <- bundles
	<-time.After(time.Second * 10)
}

func check(bundleCh chan []bundle) {
	blockNum := new(big.Int).SetUint64(0)
	for {
		select {
		case bundle := <-bundleCh:
			for _, b := range bundle {
				fmt.Println(b.x)
				fmt.Println(blockNum.Cmp(b.y))
			}

		}
	}
}

type bundle struct {
	x uint64
	y *big.Int
}
