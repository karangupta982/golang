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

// func modifySlice(s []int) {
//     s[0] = 999       // Updates the underlying array! Original sees this.
//     s = append(s, 4) // Changes the local COPY of the header's length. Original DOES NOT see this.
//     fmt.Println("Inside function:", s) // [999, 2, 3, 4]
// }

// func main() {
//     nums := []int{1, 2, 3}
//     modifySlice(nums)
//     fmt.Println("Outside function:", nums) // [999, 2, 3] (The '4' is missing!)
// }


// ---

// type NumArray struct {
//     prefixSum []int
// }


// func Constructor(nums []int) NumArray {
//     arr := make([]int, len(nums) + 1)  // Use make() to dynamically allocate a slice instead of a fixed array

//     for i := range nums {
//         arr[i + 1] = arr[i] + nums[i] 
//     }
//     return NumArray{
//         prefixSum: arr,
//     }
// }


// func (this *NumArray) SumRange(left int, right int) int {
//     return this.prefixSum[right+1] - this.prefixSum[left]
// }


----

// why b := a does not copy slice a?
// b := a does make a copy, but it only copies the slice header, not the underlying array where the actual data lives. 
// 1. What is a Slice really?
// A slice is not the data itself. It is a small descriptor struct (a header) containing just three things:

// A pointer pointing to an underlying array in memory.

// A length (len) tracking how many elements are currently in use.

// A capacity (cap) tracking the maximum size before reallocation is needed.

// 2. What happens during b := a?
// Go follows a strict rule: assignment always copies the value.

// When you write b := a, Go copies the value of a. Because a is a slice, its "value" is that small header (pointer + length + capacity).

// b gets its own brand-new header.

// However, the pointer inside b's header is copied from a, meaning it points to the exact same underlying array.

// The underlying array is always a rigid, continuous block of memory.

// The slice is just a remote control (header) pointing to it.

// When you outgrow the array, Go doesn't magically stretch the old memory block (it can't, because other data might be sitting right next to it in RAM). Instead, Go abandons the old array, builds a bigger continuous array somewhere else, and points your slice to the new one.