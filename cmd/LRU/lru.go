package main

import (
	"container/list"
)

func main() {

}

// create hashmap to store key,value & linked list to store the least/most recently used items
// TODO:Read op
// If item in the cache , return the value + move/update the item to head of the linked list
// If item not in the cache, return nil

// TODO: Write op

// When new put comes, If the linked list is not full, add the item to map & add the item to head of the list
// If the list is full, remove the item from the tail of the list as well as from map
// then add the item to map & add item to head of the list

type LRU struct {
	cache map[string]int
	list  *list.List
	data  int
}

func New(cap int) {
}
