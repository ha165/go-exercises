package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d completed job %d\n", id, job)
	}

}

func main() {
	var wg sync.WaitGroup

	jobs := make(chan int)

	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	for job := 1; job <= 8; job++ {
		jobs <- job
	}
	close(jobs)
	wg.Wait()
}
