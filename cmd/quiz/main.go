package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

const (
	f float64 = 89.00
)

func CopyFile(d, s string) (w int64, err error) {
	src, err := os.Open(s)
	if err != nil {
		return
	}
	dst, err := os.Create(d)
	defer dst.Close()
	if err != nil {
		return
	}
	w, err = io.Copy(dst, src)
	return

}

func forL() {

	i := 1
	for i < 10 {
		println(i)
	}
}

func funcT() {
	b := []string{"osprey", "stork"}
	f := []string{"shark", "salmon"}
	var a []string
	a = f
	a = append(a, b[:1]...)
	fmt.Println(a)
}
func randTest() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println(rand.Intn(899) + 100)
}
func main() {
	randTest()
	fmt.Println("----------------")
	u := &User{
		Name: "afa",
		Role: 1,
	}
	makeAdmin(*u)
	a := hasAccess(u)
	fmt.Println(a)
	funcT()
	//m := make(go_map[string]int)
	//v, x := m["testplay"]
	//
	//fmt.Println(v)
	//fmt.Println(x)
	//var b bytes.Buffer
	//takeReader(&b)
}

func takeReader(b io.Writer) {

}

func makeAdmin(u User) {
	u.Role = 1
}
func hasAccess(u *User) bool {
	return u.Role == 1
}

type User struct {
	Name string
	Role int
}
