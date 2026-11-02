package main

import (
	"fmt"
	"sync"
)

// Mutex is used in case of race condition (two or more resources trying to modify the same resource or parameter or field)

type post struct{
	views int
	mu sync.Mutex
}

func (p *post) increase(wg *sync.WaitGroup){
	defer func(){
		p.mu.Unlock()
		wg.Done()
	}()

	p.mu.Lock()
	p.views += 1
	
}


func main(){

	var wg sync.WaitGroup
	myPost := post{
		views:0,
	}

	for i:=0;i<100;i++{
		wg.Add(1)
		go myPost.increase(&wg)
	}

	

	wg.Wait()
	fmt.Println(myPost.views)
	
}