package main

import (
	"container/list"
	"crypto/ecdsa"
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

type Action interface {
	Add(key, value string)
	Get(key string) string
}
type LRU struct {
	cache map[string]list.Element
	list  *list.List
	cap   int
	data  int
}

type pair struct {
	key, value string
}

func New(cap int) Action {
	return &LRU{
		cache: make(map[string]list.Element, cap),
		list:  list.New(),
		cap:   cap,
		data:  0,
	}
}

func (l *LRU) Add(key, value string) {
	p := pair{key: key, value: value}
	if len(l.cache) == l.cap {
		// fetch last element
		last := l.list.Back()
		// remove from the list
		l.list.Remove(last)
		// delete the same element from cache
		delete(l.cache, last.Value.(pair).key)
	}
	l.cache[key] = list.Element{Value: p}
	// fetch the first element
	front := l.list.Front()
	if front == nil {
		l.list.PushFront(p)
		return
	}
	// insert new key before front element
	l.list.InsertBefore(p, front)
}

func (l *LRU) Get(key string) string {
	// fetch front element from the list
	front := l.list.Front()

	if v, ok := l.cache[key]; ok {
		if front.Value.(pair).key != key {
			// insert the latest key before front
			l.list.InsertBefore(v, front)
		}
		return v.Value.(pair).value
	}
	return ""

}

func crypto() {
	ecdsa.GenerateKey()
}
