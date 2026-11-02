package main

import (
	"fmt"
	"maps"
)

// maps -> hash, object, dictionary
func main(){
	// creating map, map[keytype]valueType
	// mpp := make(map[string]string)

	// mpp["name"] = "golang"
	// mpp["area"] = "backend"

	// fmt.Println(mpp["name"], mpp["area"])

	// // if accessing a value not present in map then map will return empty string in case of string, false in case of boolean, 0 in case of int
	// fmt.Println(len(mpp))
	
	// // deleting from map
	// delete(mpp, "area")
	// clear(mpp)

	// fmt.Println((mpp))


	mpp2 := map[string]int{}

	mpp2["a"] = 5
	mpp2["b"] = 1
	// fmt.Println(len(mpp2))

	val, isPresent := mpp2["c"]

	if isPresent{
		fmt.Println("exist", val)
	} else{
		fmt.Println("does not exist", val)
	}

	// maps package
	mpp3 := map[string]int{}
	maps.Copy(mpp3,mpp2)
	fmt.Println(mpp3)

	fmt.Println(maps.Equal(mpp2,mpp3))
}


// 1. What is struct{}{}?
// struct{} is an empty struct—a struct type with zero fields.

// struct{}{} is the literal value of that type (instantiating an empty struct).

// Because an empty struct has no fields, it takes up zero bytes of memory (unsafe.Sizeof(struct{}{}) == 0).

// map[int]bool: Takes 1 byte per entry for the boolean (true/false), Slightly shorter to insert: set[x] = true

// map[int]struct{}: Takes 0 bytes per entry, Slightly more verbose: set[x] = struct{}{}

// bool set code
	// set := make(map[int]bool)

	// // Insert
	// set[10] = true

	// // Check existence
	// if set[10] {
	//     fmt.Println("Exists")
	// }

// struct set code
	// set := make(map[int]struct{})

	// // Insert (assigns an empty, 0-byte value)
	// set[10] = struct{}{}

	// // Check existence (the standard "comma ok" idiom)
	// if _, exists := set[10]; exists {
	//     fmt.Println("Exists")
	// }