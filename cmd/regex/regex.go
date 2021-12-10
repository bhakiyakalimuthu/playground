package main

import (
	"fmt"
	"regexp"
)

var namePat = regexp.MustCompile(`^[A-Z]{2}\d{3}$`)

func main() {
	fmt.Println(namePat.String())
}
