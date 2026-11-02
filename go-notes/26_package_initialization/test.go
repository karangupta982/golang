package main

import "fmt"

func init(){
	fmt.Println("init called")
}

var package_variable = package_Var()

func package_Var() int{
	fmt.Println("package variable called")
	return 1
}



func main(){
	fmt.Println("main called")
}

// output: 
// package variable called
// init called
// main called

// Package Initialization

// Order:
// 1. Initialize imported packages
// 2. Initialize package-level variables
// 3. Execute init() functions
// 4. Execute main()

// Purpose:
// Prepare a package before the application starts.


// 1. Imports: Go look at all the packages you imported. If those packages have their own imports, variables, or init() functions, Go resolves them first.
// 2. Package Variables: Go sets up all the global variables in your current package. If a variable calls a function to get its value (like your package_Var() function), that function runs right then.
// 3. init() Functions: Go runs the init() functions in your package.
// 4. main() Function: Finally, Go hands control over to your main() function, and your core application logic begins.