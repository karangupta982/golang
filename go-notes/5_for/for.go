package main

import "fmt"

func main(){
	// basic for loop
	for i:=0; i<5; i++ {
		fmt.Println(i)
		if i==3 {
			continue
		}
		if i==4 {
			break
		}
	}

	// infinite loop
	k := 0
	for {
		if k==4{
			break
		}
		fmt.Println(k)
	}

	// range loop
	for j := range 6 {
		fmt.Println(j)
	}

	// nested loop
	for i:=0; i<5; i++ {
		for j:=0; j<5; j++ {
			fmt.Println(i, j)
		}
	}
}
