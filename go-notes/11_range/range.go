package main

import "fmt"

// range is used for iterating over data structures
func main(){
	nums := []int{1,2,3}

	// for i:=0;i<len(nums);i++{
	// 	fmt.Println(nums[i])	
	// }

	for ind,num := range nums{
		fmt.Println(num, ind)
	}

	// range over map
	mpp := map[string]string{}
	mpp["first_name"] = "Karan"
	mpp["Last_name"] = "Gupta"

	// for key, val := range mpp{
	// 	fmt.Println(key, val)
	// }

	// range over string
	// c-> unicode, i is not index
	for i,c := range "golang"{
		fmt.Println(string(c), i)
	}
}
