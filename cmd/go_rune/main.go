package main

import "fmt"

func main() {
	s := "success"
	temp := make([]rune, 0)

	for _, val := range s {
		temp = append(temp, val)
	}
	fmt.Println(string(temp))

}
