package main

import "fmt"
// Logical operators are used to combine multiple conditions.
// && (AND)
// || (OR)
// ! (NOT)
func main(){
	age := 15
	if age > 60 {
		fmt.Println("Senior Citizen")
	}
	if age>=18 && age<=60 {
		fmt.Println("Adult")
	} else if age > 12 {
		fmt.Println("Teenager")
	} else {
		fmt.Println("Minor")
	}

	var role = "admin"
	var hasPermission = true
	if role == "admin" || hasPermission {
		fmt.Println("You have access to the system")
	} else {
		fmt.Println("You do not have access to the system")
	}

	// we can declare a variable inside the if statement
	if age := 15; age > 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}

	
}
