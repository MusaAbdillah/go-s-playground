package main

import (
	"fmt"
)

type Customer struct {
	Id   int
	Name string
}

func worker(id int, jobs <-chan Customer, results chan<- Customer) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job for customer", j.Name)
		results <- j
	}
}

func main() {

	var (
		customers []Customer
	)

	for i := 0; i <= 10; i++ {
		customers = append(customers, Customer{
			Id:   i,
			Name: fmt.Sprintf("Customer %v", i),
		})
	}

	// fmt.Println("===customers===")
	// fmt.Println(customers)

	var numJobs int = len(customers)
	jobs := make(chan Customer, len(customers))
	// jobOne := make(chan int, numJobs)
	results := make(chan Customer, numJobs)

	// for w := 1; w <= 3; w++ {
	// 	go workerOne(w, jobOne)
	// }

	for w := 1; w <= numJobs; w++ {
		go worker(w, jobs, results)
	}

	// // its only need if you wanna result
	for _, customer := range customers {
		jobs <- customer
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
