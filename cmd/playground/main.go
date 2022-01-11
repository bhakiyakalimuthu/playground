package main

import (
	"fmt"
	"sort"
)

func main() {

	s := []int{1, 2, 3, 4}

	s = append(s[:2], s[1:]...)
	s[2] = 6
	fmt.Print(s)

	sliceSort()
}

func sliceSort() {
	Author := []struct {
		a_name    string
		a_article int
		a_id      int
	}{
		{"Mina", 304, 1098},
		{"Cina", 634, 102},
		{"Tina", 104, 105},
		{"Rina", 10, 108},
		{"Sina", 234, 103},
		{"Vina", 237, 106},
		{"Rohit", 56, 107},
		{"Mohit", 300, 104},
		{"Riya", 4, 101},
		{"Sohit", 20, 110},
	}

	sort.Slice(Author, func(i, j int) bool {
		return Author[i].a_article < Author[j].a_article
	})
	fmt.Printf("Author %v", Author)
}
