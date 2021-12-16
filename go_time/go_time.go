package main

import (
	"fmt"
	"time"
)

func main() {
	//uint8  : 0 to 255
	//uint16 : 0 to 65535
	//uint32 : 0 to 4294967295
	//uint64 : 0 to 18446744073709551615
	//int8   : -128 to 127
	//int16  : -32768 to 32767
	//int32  : -2147483648 to 2147483647
	//int64  : -9223372036854775808 to 9223372036854775807

	// time now
	now := time.Now()
	fmt.Println(now)

	// unix time stamp
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixMilli())

	// sub
	fmt.Println(now.Sub(time.Unix(1639332412, 0)))
	fmt.Println(now.Add(time.Hour))

	// parse
	d, _ := time.Parse("2006-01-02", "2021-12-12")
	fmt.Printf("parse %v\n", d)

	// format

	fmt.Printf("format %s", now.Format("2006-01-02:16:30:00"))

	//
}
