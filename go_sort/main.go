package main

import (
	"fmt"
	"sort"
)

type P struct {
	name string
	age  int
}

var persons = []P{
	{"huawei", 15}, {"nokia", 6}, {"bluejeans", 18}, {"kindred", 12}, {"voi", 18}, {"arango", 3},
}

type Persons []P

func (p Persons) Len() int {
	return len(p)
}
func (p Persons) Less(i, j int) bool {
	return p[i].age < p[j].age

}
func (p Persons) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func main() {
	sort.Sort(Persons(persons))
	fmt.Println(persons)
}
