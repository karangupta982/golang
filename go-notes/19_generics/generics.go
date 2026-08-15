package main

import (
	"fmt"
	"cmp"

)

// Generics was introduced in 1.18 version of Go, initially it was not present
// Generics allow you to write functions, structures, and interfaces without hardcoding the exact data type upfront.


// func iterateOverIntegerArray(nums []int){
// 	for _,val := range nums{
// 		fmt.Println(val)
// 	}
// }
// func iterateOverStringArray(names []string){
// 	for _,val := range names{
// 		fmt.Println(val)
// 	}
// }
//						T any or T interface{} or T comparable can also be written
func iterateOverArray[T int | string | bool, V string](nums []T, name V){
	for _, val := range nums{
		fmt.Println(val, name)
	}
}

type stack[T int | string ] struct{
	elements []T
}


// 1. Using standard constraints (cmp.Ordered)Go provides built-in packages like cmp for common operations. cmp.Ordered restricts your function to types that naturally support ordering operators like <, >, and == (integers, floats, and strings).

// T can now only be types that are orderable
func GetMax[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}


// 2. Creating Custom Constraints (The ~ operator)You can define a custom interface that acts as a constraint list using a pipe symbol (|) for "OR" relationships.
// Numbers can only be int, int64, or float64
type Numbers interface {
	~int | ~int64 | ~float64
}

func Sum[T Numbers](a, b T) T {
	return a + b
}
// (The tilde ~ symbol is a neat trick—it means the function also accepts custom types derived from those primitives, like type MyAge int).
 


func main(){

	nums := []int{1,2,3}
	names := []string {"golang","python"}

	// iterateOverIntegerArray(nums)

	// iterateOverStringArray(names)

	iterateOverArray(nums, "Golang")
	iterateOverArray(names, "Golang") 

	// myStack := stack[string]{
	// 	elements: []string{"golang","python"},
	// } 

}


// Interface vs. Generics (When to use what?)It is easy to confuse interfaces (any) and generics ([T any]), but they serve different goals:Use Interfaces when you focus on behavior. If you want a type because it knows how to Speak() or Write(), use an interface.Use Generics when you focus on type safety and efficiency. If you are writing utilities that move data around (like sorting arrays, filtering maps, or caching structures) without changing the data itself, use generics.