package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int

func main() {

	relayURLs := []string{"url1", "url2"}
	result := false
	var lastRelayError error
	var wg sync.WaitGroup
	for _, url := range relayURLs {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			res, err := makeRequest()
			// Check for errors
			if err != nil {
				fmt.Println("error making request to relay")
				return
			}
			if res.Error != nil {
				fmt.Println("error reply from relay")
				lastRelayError = res.Error
				return
			}
			fmt.Printf("resp result %v\n", res.Result)
			// Decode the response
			_result := res.Result

			// Result should be true if any one relay responds true
			result = result || _result
			fmt.Printf("inside result %v\n", result)
		}(url)
	}
	wg.Wait()
	fmt.Printf("final result %v\n", result)
	fmt.Printf("lastRelayError %v\n", lastRelayError)
}

func makeRequest() (*response, error) {
	counter = counter + 1
	fmt.Printf("counter value %v\n", counter)
	time.Sleep(time.Second)
	if counter == 1 {
		return &response{
			Result: true,
			Error:  nil,
		}, nil
	}
	return &response{
		Result: false,
		Error:  nil,
	}, nil
}

type response struct {
	Result bool
	Error  error
}
