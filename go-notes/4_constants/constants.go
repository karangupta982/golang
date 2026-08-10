package main

import "fmt"

// Vairables and constants are declared at the package level.
var location string = "India"
const currentGoal string = "CNCF contributions"

// But shorthand declaration is only allowed at the function level.
// learning := "Goland" // wrong syntax

func main(){
	// Constants are immutable values that are defined at compile time and cannot be changed during runtime.
	// const name string = "Karan"
	// const age int = 24


	const name = "Karan"
	const age = 24

	fmt.Println(name, age, location, currentGoal)

	const (
		port = 5000
		host = "localhost"
	)
	fmt.Println(port, host)

	var (
		problem = "Discipline"
		goal = "CNCF contributions"
	)
	fmt.Println(problem, goal)
}
