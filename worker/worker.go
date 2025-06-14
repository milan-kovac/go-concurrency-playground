package worker

import (
	"fmt"
	"time"
)

type Job struct {
	ID int
}

func Worker(id int, jobs <-chan Job, results chan<- string) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job.ID)
		time.Sleep(1 * time.Second)
		results <- fmt.Sprintf("Worker %d finished job %d", id, job.ID)
	}
}

func Run() {
	const totalJobs = 5
	const totalWorkers = 3

	jobs := make(chan Job, totalJobs)
	results := make(chan string, totalJobs)

	for worker := 1; worker <= totalWorkers; worker++ {
		go Worker(worker, jobs, results)
	}

	for job := 1; job <= totalJobs; job++ {
		jobs <- Job{ID: job}
	}

	close(jobs)

	for a := 1; a <= totalJobs; a++ {
		fmt.Println(<-results)
	}
}
