package main

import "fmt"

// variadic functions takes n number of arguements
func sum(nums ...int) int{
	add := 0
	for _, num := range nums{
		add = add + num
	}
	return add
}

func main(){

	nums := []int{1,2,3,4}
	res := sum(nums...)

	res1 := sum(1,2,3,4,5)
	fmt.Println(res1,res)
}

