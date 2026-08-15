package main

import "fmt"

// by value
func changeNum(num int) {
	num = 5
	fmt.Println("num value after change", num)
}

func changeNum2(num *int){
	*num = 10
	fmt.Println("num2 after change", *num) 
}

func main(){

	num := 1
	changeNum(num)
	fmt.Println("Num value before change in main", num)

	num2 := 3
	changeNum2(&num2)
	fmt.Println("num2 after change in main", num2)

}
