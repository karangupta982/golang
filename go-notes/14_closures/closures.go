package main

import "fmt"


func outer() func() int{
	count := 0

	return func() int{
		count += 1
		return count
	}
}


func main(){
	increament := outer()
	fmt.Println(increament())
	fmt.Println(increament())

}
