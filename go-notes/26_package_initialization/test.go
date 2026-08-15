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

// Common Uses:
// - Driver registration
// - Configuration setup
// - Global initialization