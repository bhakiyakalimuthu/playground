package main

import (
	"fmt"
	"playground/cmd"
	"sync"
)

var (
	mutex sync.Mutex
)

func init() {
	cmd.Balance += 100.0
}

func main() {

	s := []int{1, 2}
	var wg sync.WaitGroup
	wg.Add(len(s))
	for _, _ = range s {

		func(group *sync.WaitGroup) {
			var wgL sync.WaitGroup
			wgL.Add(3)
			go func() {
				deposit(50.0, &wgL)
			}()
			//wg.Add(1)
			go func() {
				withdraw(20, &wgL)
			}()
			//wg.Add(1)
			go func() {
				deposit(15, &wgL)
			}()
			wgL.Wait()
			group.Done()
		}(&wg)

	}

	wg.Wait()
	fmt.Println(cmd.GetBalance())
}

func deposit(amount float32, wg *sync.WaitGroup) {
	mutex.Lock()
	cmd.Balance += amount
	mutex.Unlock()
	wg.Done()
}

func withdraw(amount float32, wg *sync.WaitGroup) {
	mutex.Lock()
	cmd.Balance -= amount
	mutex.Unlock()
	wg.Done()
}

//func getBalance()  float32{
//	return balance
//}
