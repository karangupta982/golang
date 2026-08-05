package main

import "fmt"
import "time"
import "errors"

func main(){
	fmt.Println("Boolean : ", true)
	fmt.Println("Boolean : ", false)

	fmt.Println("String : ", "Hello world")
	fmt.Println("Number : ", 1)
	fmt.Println("Float : ", 1.5)

	fmt.Println("Array : ", [3]int{1, 2, 3})
	fmt.Println("Slice : ", []int{1, 2, 3})
	fmt.Println("Map : ", map[string]int{"a": 1, "b": 2})
	fmt.Println("Struct : ", struct {
		Name string
		Age int
	}{Name: "John", Age: 20})

	
	fmt.Println("Function : ", func() string { return "Hello world" }())
	fmt.Println("Interface : ", interface{}("Hello world"))
	fmt.Println("Channel : ", make(chan int))
	fmt.Println("Error : ", errors.New("error"))
	fmt.Println("Time : ", time.Now())
	fmt.Println("Duration : ", time.Duration(1000000000))
	fmt.Println("Timezone : ", time.Now().Location())
}