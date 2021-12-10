package main

import (
	"log"
	"net/http"
)

type Site struct {
	URL string
}

type Result struct {
	URL        string
	StatusCode int
}

func crawl(wID int, jobs <-chan Site, result chan<- Result) {
	for site := range jobs {
		log.Printf("wID:%d", wID)
		resp, err := http.Get(site.URL)
		if err != nil {
			log.Println(err.Error())

		}
		log.Printf("wID : %d, URL : %s, statusCode : %d ", wID, site.URL, resp.StatusCode)
		result <- Result{URL: site.URL, StatusCode: resp.StatusCode}

	}
}
func main() {

	jobs := make(chan Site, 5)
	result := make(chan Result, 5)

	for w := 1; w <= 5; w++ {
		go crawl(w, jobs, result)

	}
	urls := []string{
		"https://google.se",
		"https://youtube.com",
		"https://amazon.se",
		"https://facebook.com",
		"https://bing.com",
	}
	for _, url := range urls {
		jobs <- Site{URL: url}

	}
	close(jobs)
	for i := 0; i < len(urls); i++ {
		r := <-result
		log.Println(r)
	}

}
