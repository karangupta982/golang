package main

import "fmt"

// Arrays are fixed size, contiguous memory allocation
// If you know size then use it
// Memory optimized
// Constant time access

func main(){
	var nums [4]int
	fmt.Println(len(nums))

	nums[0] = 1
	fmt.Println(nums[0])

	nums2 := [4]int{1,2,2,3}
	fmt.Println(nums2)

	// 2D arrays
	nums3 := [2][2]int{{1,2},{2,3}}
	fmt.Println(nums3)


}
