package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("Worker ", id, "started job ", job)
		time.Sleep(time.Second)

		results <- job
		fmt.Println("Worker", id, "finished job", job)
	}
}

func main() {
	const n = 20
	jobs := make(chan int, n)
	results := make(chan int, n)
	var wg sync.WaitGroup

	for i := range n {
		jobs <- i
	}
	// trabajadores creados
	wg.Add(1)
	for i := 1; i <= 4; i++ {
		go func() {
			defer wg.Done()
			worker(i, jobs, results)
		}()
	}

	go func() {
		wg.Wait()
		close(jobs)
	}()

	// Resultados
	for range n {
		<-results
	}
}
