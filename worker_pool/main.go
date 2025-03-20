package main

import (
	"fmt"
	"sync"
)

type Customer struct {
	Id   int
	Name string
}

func worker(id int, jobs <-chan Customer, results chan<- Customer, wg *sync.WaitGroup) {

	defer wg.Done()

	for j := range jobs {
		fmt.Println("worker", id, "started  job for customer", j.Name)
		results <- j
	}
}

func main() {

	var (
		customers []Customer
		wg        sync.WaitGroup
	)

	for i := 0; i <= 10; i++ {
		customers = append(customers, Customer{
			Id:   i,
			Name: fmt.Sprintf("Customer %v", i),
		})
	}

	// fmt.Println("===customers===")
	// fmt.Println(customers)

	var workerNumber int = 3
	var numJobs int = len(customers)
	jobs := make(chan Customer, len(customers))
	// jobOne := make(chan int, numJobs)
	results := make(chan Customer, numJobs)

	// for w := 1; w <= 3; w++ {
	// 	go workerOne(w, jobOne)
	// }

	for w := 1; w <= workerNumber; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// // its only need if you wanna result
	for _, customer := range customers {
		jobs <- customer
	}
	close(jobs)
	wg.Wait()

	// for a := 1; a <= numJobs; a++ {
	// 	<-results
	// }
}
