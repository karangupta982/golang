package main

import "fmt"

// custom types
// type Mytype string

// enums are enumerated types

// type OrderStatus int

// const (
// 	Received OrderStatus = iota //iota is an auto increamented number starting from 0
// 	Confirmed
// 	Prepared
// 	Delivered 
// )

// func changeStatus(status OrderStatus){
// 	fmt.Println("Changing Status to", status)
// }



type OrderStatus string

const (
	Received OrderStatus = "received"
	Confirmed = "confirmed"
	Prepared = "prepared"
	Delivered = "delivered"
)

func changeOrderStatus(status OrderStatus){
	fmt.Println("Changing order status", status)
}


func main(){
	// changeStatus(Prepared)
	changeOrderStatus(Confirmed)
}



// type ServerState int

// const (
//     StateIdle ServerState = iota
//     StateConnected
//     StateError
//     StateRetrying
// )

// var stateName = map[ServerState]string{
//     StateIdle:      "idle",
//     StateConnected: "connected",
//     StateError:     "error",
//     StateRetrying:  "retrying",
// }

// func (ss ServerState) String() string {
//     return stateName[ss]
// }

// func main() {
//     ns := transition(StateIdle)
//     fmt.Println(ns)

//     ns2 := transition(ns)
//     fmt.Println(ns2)
// }

// func transition(s ServerState) ServerState {
//     switch s {
//     case StateIdle:
//         return StateConnected
//     case StateConnected, StateRetrying:

//         return StateIdle
//     case StateError:
//         return StateError
//     default:
//         panic(fmt.Errorf("unknown state: %s", s))
//     }
// }

// When fmt.Println(ns) is called, Go checks if the ServerState type implements the Stringer interface. Since you defined the method func (ss ServerState) String() string, Go automatically calls it. The map lookup stateName[StateConnected] returns "connected", which is printed to the console.


// The reason Go automatically finds and calls your method is due to Go's unique approach to interfaces called implicit implementation (often described as "duck typing").
// You do not need to explicitly declare that your type implements an interface. If your type implements a method with the exact signature required by an interface, Go considers that interface implemented.


// What is the Stringer interface?

// The fmt.Stringer interface is a built-in interface defined in Go's standard fmt package. Its definition looks exactly like this:

// package fmt

// type Stringer interface {
//     String() string
// }

// Any type that defines a String() string method automatically satisfies this interface.

// -----

// In Go, the formatting functions in the fmt package are categorized by their prefixes (Print, Fprint, Sprint) and their suffixes (none, f, ln).
// Understanding this grid allows you to easily decipher exactly what any function does during an interview.

// 🗺️ The Ultimate Core Rule Grid

// To understand any fmt function, just combine its Prefix (Where the output goes) with its Suffix (How it formats).

// 1. The Prefixes (Where does it write?)

// • Print...: Outputs directly to the Standard Output (the terminal console).
// • Sprint...: "String Print". It returns the result as a new string variable instead of printing it to the console.
// • Fprint...: "File/Format Print". It writes the output directly into an io.Writer (such as a file, network socket, or HTTP response body).

// 2. The Suffixes (How does it format?)

// • No Suffix: Arguments are printed using their default layouts. Spaces are added between arguments only if neither is a string.
// • ...f (Formatted): Uses a template string with verbs/placeholders (like %s, %d, %v) to give you strict control over spacing, alignment, and types.
// • ...ln (Line): Spaces are automatically added between all arguments, and a newline (\n) is always appended to the very end.

// 📋 Direct Comparisons & Use Cases


// The Standard Console Group (Print, Printf, Println)

// Use these when you want to quickly display information directly to a user running your CLI program.
// go
// name := "Alice"
// age := 25

// // 1. Println: Adds spaces between arguments and a newline at the end.
// // Use case: Quick debug logging, separated items.
// fmt.Println(name, "is", age) // Output: Alice is 25\n

// // 2. Printf: Evaluates format verbs. Does NOT add a newline automatically.
// // Use case: Structuring clean layouts, formatting decimals/floats.
// fmt.Printf("%s is %d years old.\n", name, age) // Output: Alice is 25 years old.\n

// // 3. Print: Default layout, no automatic spaces or newlines.
// // Use case: Building a single line piece by piece.
// fmt.Print(name, " is ", age, "\n") // Output: Alice is 25\n
// Use code with caution.

// The String Variable Group (Sprint, Sprintf, Sprintln)

// Use these when you don't want to show anything to the terminal yet, but need to build a string to save into a variable, database, or struct field.
// go
// // Use case: Creating error messages or dynamic text strings.
// errorMessage := fmt.Sprintf("Error: User %s not found (Code: %d)", name, 404)

// // errorMessage now holds the string. Nothing has printed to the console yet.
// Use code with caution.

// The Writer Group (Fprint, Fprintf, Fprintln)

// Use these in production systems to pass data directly into files, network connections, or web responses without allocating intermediary strings in memory.
// go
// // Use case: Writing to an open file or an HTTP response stream.
// var buf bytes.Buffer
// fmt.Fprintf(&buf, "Log entry for %s", name)
// Use code with caution.

// 💡 Interview Pro-Tip: Common Format Verbs (%v vs %+v vs %#v)

// Interviewers often ask how to debug-print a complex struct using Printf. You should memorize these three variations of the %v (value) verb:
// go
// type User struct { Name string; Age int }
// u := User{"Bob", 30}

// fmt.Printf("%v\n", u)   // Value only: {Bob 30}
// fmt.Printf("%+v\n", u)  // Field names + values: {Name:Bob Age:30}
// fmt.Printf("%#v\n", u)  // Go syntax representation: main.User{Name:"Bob", Age:30}
// Use code with caution.
// Would you like to look at the performance implications of using fmt.Sprintf versus basic string concatenation (+) inside loops, or should we talk about the built-in global println() function and why you should avoid it in production code?


