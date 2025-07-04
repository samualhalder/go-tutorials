package main

import (
	"fmt"
	"time"
)

func workers(id int, jobs <-chan int, results chan<- int) error {
	for job := range jobs {
		fmt.Printf("worker %d is doing job %d\n", id, job)
		time.Sleep(time.Second)
		results <- job * 2
	}
	return nil
}

func main() {
	numOfJobs := 5
	jobs := make(chan int, numOfJobs)
	results := make(chan int, numOfJobs)
	for id := 1; id <= 3; id++ {
		go workers(id, jobs, results)
	}

	for job := 1; job <= numOfJobs; job++ {
		jobs <- job
	}
	close(jobs)
	for i := 0; i < numOfJobs; i++ {
		fmt.Println(<-results)
	}

}
