package main

import (
	"fmt"
	"strconv"
	
	"sync"
)



func worker(i int, jobs <-chan int, res chan<- string , wg *sync.WaitGroup){
	defer wg.Done()

	for job := range jobs{
		fmt.Printf("job processed by workerId %d and job no %d\n",i, job)
		responseMessage := "job Completed" + strconv.Itoa(job)
		res <- responseMessage
	}
}

func main(){
	var wg sync.WaitGroup

	var jobs = make(chan int)
	var res = make(chan string)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(i, jobs, res, &wg)
	}

	arr := []string{}
	for j := 0; j < 12; j++ {
		jobs <- j
		arr = append(arr, <-res)
	} 
	
	close(jobs)

	wg.Wait()

	fmt.Println("RESPONSE covered")
	for val := range arr{
		fmt.Printf("response %d", val)
	}

}