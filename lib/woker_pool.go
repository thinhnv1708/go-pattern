package lib

import (
	"fmt"
	"runtime"
)

func PoolWorker() {
	number := 500
	numberOfWorkers := runtime.NumCPU()

	jobs := make(chan int, number)
	resultChan := make(chan int, number)

	for n := 0; n < numberOfWorkers; n++ {
		go worker(jobs, resultChan)
	}

	for job := 0; job <= number; job++ {
		jobs <- job
	}

	close(jobs)
	fmt.Println("Quat")
	for n := 0; n <= number; n++ {
		fmt.Println(<-resultChan)
	}

}

func worker(jobs <-chan int, resultChan chan<- int) {
	for param := range jobs {
		resultChan <- fib(param)
	}

}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}
