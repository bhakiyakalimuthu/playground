package main

func processJob(job int) {
	// do something
}

func main() {
	jobs := [2]int{1, 2}
	for job := range jobs {
		go func() {
			processJob(job)
		}()
	}
}
