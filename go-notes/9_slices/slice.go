package main

import (
	"fmt"
	"slices"
)

// Slice is a dynamic array
func main(){
	// uninitialized slice is nil
	// var nums []int
	// fmt.Println(nums == nil)
	// fmt.Println(nums, len(nums), cap(nums))

	// // initialized slice is not nil
	// nums = make([]int, 3)
	// fmt.Println(nums, len(nums), cap(nums))

	// // initialized slice with length = 3 and capacity = 5
	// nums = make([]int, 3, 5)
	// fmt.Println(nums, len(nums), cap(nums))

	// // append to slice
	// nums = append(nums, 1, 2, 3)
	// fmt.Println(nums, len(nums), cap(nums))

	// If capacity is less than length then capacity gets doubled


	// empty slice
	nums4 := []int{}
	nums4 = append(nums4, 1,2,3)
	// fmt.Println(nums4, len(nums4), cap(nums4))

	// copy function
	nums5 := make([]int, len(nums4))
	copy(nums5, nums4)
	fmt.Println(nums5, len(nums5), cap(nums5))

	// slice function
	nums6 := nums4[1:3]
	fmt.Println(nums6, len(nums6), cap(nums6))

	// slice function with capacity
	nums7 := nums4[1:3:3]
	fmt.Println(nums7, len(nums7), cap(nums7))

	// slice package
	fmt.Println(slices.Equal(nums4,nums5))

	var nums = [][]int{{1,2,3},{4,5,6}}
	fmt.Println(nums)






// 	This happens because of how Go handles memory layout and how it copies variables when passing them into functions.
// Go is strictly a "pass-by-value" language. This means whenever you pass an argument to a function, Go always makes a copy of it. However, what gets copied is completely different for an array versus a slice.
// ------------------------------
// ## 1. The Array: A Solid Block of Data
// An array variable is the entire block of data. If you have an array of 5 integers, that variable represents all 5 integers sitting side-by-side in memory.

// * What happens in the function? When you pass myArr into modifyArray(a [2]int), Go makes a 100% duplicate copy of the entire block of memory.
// * The Result: The function modifies its own local duplicate clone. The original array back in main() remains completely untouched.

// Original Array (in main):   [ 1 | 2 ]
//                                ↓ (Copies the whole box)
// Function's Clone Array:     [ 1 | 2 ]  → modified to →  [ 999 | 2 ]

// ------------------------------
// ## 2. The Slice: A Small Window (Header)
// A slice variable does not actually store your data. Instead, a slice is a tiny structural header (a descriptor) that contains three small things: [3, 4, 5] 

//    1. A pointer (the memory address) pointing to the actual underlying array where data is stored.
//    2. A Length (len).
//    3. A Capacity (cap). [6, 7, 8] 


// * What happens in the function? When you pass mySlice into modifySlice(s []int), Go still makes a copy—but it only copies that tiny header (the pointer address, length, and capacity). It does not copy the actual data items. [9] 
// * The Result: The cloned header inside the function still contains the exact same memory pointer address as the original header. When you change an index (s[0] = 999), Go follows that pointer directly to the original spot in memory and mutates it. [10] 

// Original Slice Header (main):    [ Pointer: 0xABC123 | Len: 2 | Cap: 2 ]
//                                          ↓
//                                    (Copies only the header)
//                                          ↓
// Function Slice Header:           [ Pointer: 0xABC123 | Len: 2 | Cap: 2 ]
//                                          ↓
//                                    Both headers point to the EXACT SAME physical memory:
//                                    Shared Memory Location (0xABC123): [ 999 | 2 ]
	
}
