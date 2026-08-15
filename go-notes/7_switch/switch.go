package main

import (
	"flag"
	"fmt"
	"reflect"
	"strings"

	// "time"
	// "reflect"
	"bufio"
	"os"
)

// format specifiers are:
// %d -> Integer
// %f -> floating point Number
// %s -> String
// %t -> Boolean


func typeofVars(key any) string {
	if reflect.TypeOf(key) == reflect.TypeOf(0){
		return "Integer"
	} else if reflect.TypeOf(key) == reflect.TypeOf(""){
		return "String"
	} else if reflect.TypeOf(key) == reflect.TypeOf(true){
		return "Bool"
	} else if reflect.TypeOf(key) == reflect.TypeOf(1.1) {
		return "Float"
	} else{
		return "Not recognizable"
	}
}



func main(){
	// i := 4
	// switch i {
	// case 1: fmt.Println("one")
	// case 3: fmt.Println("three")
	// case 2: fmt.Println("two")
	// default: fmt.Println("default")
	// }

	// // multiple conditions
	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("Weekends")
	// default:
	// 	fmt.Println("Workday")
	// }

	// // empty interface is of type any
	// // interface{} is empty of type any

	// // type switch
	// // you cannot use i.(type) just anywhere. It can only be used inside the switch statement conditional block.
	// whoami := func(i interface{}) {
	// 	switch i.(type) {
	// 	case int:
	// 		fmt.Println("Integer")
	// 	case string:
	// 		fmt.Println("String")
	// 	case bool:
	// 		fmt.Println("Boolean")
	// 	default:
	// 		println("unknown type", reflect.TypeOf(i))
	// 	}
	// }
	// whoami(23)



	// // Method 1
	// // The fmt package reads whitespace-separated tokens. It stops reading a single input the moment the user hits the spacebar
	// var name string
	// var age int

	// fmt.Println("Enter your name and age")

	// _, err := fmt.Scanln(&name, &age)
	// if err != nil{
	// 	fmt.Println("Error Reading Input:", err)
	// 	return
	// }
	// // When a terminal program finishes executing, many modern shells like Zsh, automatically print a % to indicate that the output ended abruptly without a newline. This visually warns you that the program's output and the next terminal prompt are sharing the same line.
	// // Add \n to the end of your string inside fmt.Printf. This pushes the terminal cursor to the next line safely.
	// fmt.Printf("Name is %s and Age is %d\n", name, age)

	// fmt.Printf("Name is of type %s and age is of type %s\n",typeofVars(name), typeofVars(age))



	// // Method 2:
	// // If you need to capture inputs that include spaces (like a full name, home address, or sentence), you should read from standard input (os.Stdin) using a Scanner. It reads until the user presses the Enter key.
	
	// // os.Stdin is raw and low-level. It doesn't know what a "line" or a "sentence" is; it just streams raw computer bytes.
	// // Because os.Stdin only deals in raw bytes, it is hard to work with directly. This is where bufio.NewScanner() comes in. 
	// scanner := bufio.NewScanner(os.Stdin)

	// fmt.Print("Enter your full name: ")
	
	// // Scan() blocks execution until the user presses Enter
	// // When your program hits if scanner.Scan(), the following sequence happens:It pauses (blocks) your program: The code stops executing and waits patiently because it looks at os.Stdin and sees nothing there yet.It listens for the Enter key: You type your name (e.g., Karan Gupta) and press Enter.It slices the text: The Scan() function actively searches for a newline character (the Enter key). Once it finds it, it grabs everything you typed before that Enter key, saves it inside its internal memory buffer, and stops reading.It returns a signal: Scan() returns a boolean value (true or false). If it successfully read your text, it returns true. This causes your if condition to pass, allowing the code inside the curly braces {} to run.
	// if scanner.Scan() {
	// 	// Text() extracts the input as a string
	// 	// When you run fullName := scanner.Text(), you are saying: "Take that text you just buffered, convert it into a clean Go string, and hand it over to my fullName variable."
	// 	fullName := scanner.Text()
	// 	fmt.Printf("Welcome, %s!\n", fullName)
	// }


	// Method 3:
	port := flag.Int("port",4000,"port number on which server will run")
	flag.Parse() // parses the flag from execution command

	fmt.Println("Server started on port",*port)



}



func readUntilUserEnd() []string {
	var answers = []string{}
	
	scanner := bufio.NewScanner((os.Stdin))

	for{
		fmt.Println("you: ")

		if !scanner.Scan(){
			break
		}

		input := scanner.Text()

		cleanedInput := strings.TrimSpace(strings.ToLower(input))

		if cleanedInput == "exit" || cleanedInput == "quit" {
			fmt.Println("Bye Bye")
			break
		}

		answers = append(answers, input)
		fmt.Printf("Bot you said %s\n",input)		
	}

	// Every operating system gives a running program three standard communication pipelines (streams):
	// os.Stdin (Standard Input): The pipeline for receiving data (e.g., your keyboard input).
	// os.Stdout (Standard Output): The pipeline for standard messages and normal data output (where fmt.Println sends text).
	// os.Stderr (Standard Error): A completely separate pipeline specifically reserved for diagnostic messages, warnings, and errors.
	// The F in Fprintln stands for File (or Formatted File). In Go, terminal streams like os.Stdout and os.Stderr are treated exactly like files that you can write data into.Unlike fmt.Println (which automatically forces text into os.Stdout), fmt.Fprintln requires you to explicitly choose which file or stream to write to as its very first argument.textfmt.Fprintln(  os.Stderr ,  "Your Error Message"  )
	// 																	│                 │
	// 															Where to send it    What to write

	if err := scanner.Err(); err != nil{
		fmt.Fprintln(os.Stderr, "Error", err )
	} 

	return answers

}
