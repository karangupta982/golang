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
