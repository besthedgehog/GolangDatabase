package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup, jobs <-chan int, results chan<- int) {
	defer wg.Done()
	for number := range jobs {
		fmt.Printf("Worker %v is working \n", id)
		results <- number * number
	}
}

func giveJobs(jobs chan<- int, numberOfJobs int) {
	for i := range numberOfJobs {
		jobs <- i
	}
	close(jobs)
}

func collectRusults(results <-chan int, numberOfJobs int) []int {
	var resultList []int
	for tmpRes := range results {
		resultList = append(resultList, tmpRes)
	}
	return resultList
}

func runWorkers(numberOfWorkers int, wg *sync.WaitGroup, jobs <-chan int, results chan<- int) {
	for idx := range numberOfWorkers {
		wg.Add(1)
		go worker(idx, wg, jobs, results)
	}
}

func main() {
	var numberOfWorkers int = 10
	var numberOfJobs int = 100

	wg := &sync.WaitGroup{}

	// var wg *sync.WaitGroup

	chanJobs := make(chan int)
	chanResults := make(chan int)

	runWorkers(numberOfWorkers, wg, chanJobs, chanResults)

	go func() {
		wg.Wait()
		close(chanResults)
	}()

	go giveJobs(chanJobs, numberOfJobs)

	resultList := collectRusults(chanResults, numberOfJobs)
	fmt.Println(resultList)

}
