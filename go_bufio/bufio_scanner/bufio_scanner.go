package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("vim-go")
	str := "do not communicate by sharing memory"
	reader := strings.NewReader(str)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())

	}
}
