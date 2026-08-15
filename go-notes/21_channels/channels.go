package main

import (
	"fmt"
	"time"
	// "time"
	// "math/rand"
)

// Channels are used for communication between go routines
// it is used for sending data from one go routine to another go routine
// channels are blocking, it is blocked until second side is ready to receive the data

// func processNum(numChan chan int){
// 	// fmt.Println("Processing number: ",<- numChan)
// 	for num := range numChan{
// 		fmt.Println(num)
// 		time.Sleep(time.Second)
// 	}
// }

// func sum(num1 int, num2 int, result chan int){
// 	res := num1 + num2
// 	result <- res
// }

// goroutine synchronyzer
// func task(done chan bool){
// 	defer func(){
// 		done <- true
// 	}()
// 	fmt.Println("Processsing...")
// }

// we can strict this more by saying that done will be used here for sending message and emailChan would be for receiving mail only, it is done by putting sending and receiving arrows
func mailSender(emailChan <-chan string, done chan<- bool){
	defer func(){ done<- true }()

	// When you use a for mail := range emailChan loop, it tells the goroutine: "Keep reading from this channel forever until someone explicitly closes it."
	// Because your main function finishes adding the 5 emails but never closes the channel, the mailSender goroutine processes all 5 emails and then freezes (blocks) on the range line, waiting for a 6th email that will never arrive.Since mailSender is permanently stuck waiting for more emails, it never reaches your defer func() { done <- true }() line. Meanwhile, your main function is stuck waiting at <-done. Both goroutines are asleep forever, causing the deadlock.
	for mail := range emailChan{
		fmt.Println("Sending messages to ",mail)
		time.Sleep(time.Second)
	}
}


func main(){

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func(){
		chan1 <- 1
	}()

	go func(){ chan2<- "golang" }()

	// We used select here because you are dealing with multiple channels and you don't know which one will receive data first.The select statement acts like a switch board for channels. It lets a single goroutine listen to multiple channel operations simultaneously and executes whichever case is ready first.

	// What happens WITH select? (The Solution)By using select, you remove the rigid order. You are telling Go: "Look at both chan1 and chan2 at the same time. Whichever one wakes up first with data, handle that case immediately."Non-blocking choice: It constantly polls all listed cases. If chan2 finishes first, the chan2 case runs immediately.Fairness: If both channels happen to be ready at the exact same fraction of a second, Go picks one at random. This prevents one channel from starving out the other.The Loop Factor: Because you wrapped the select in a for i := 0; i < 2; i++ loop, it executes exactly twice. The first iteration grabs whichever channel finishes first, and the second iteration grabs the remaining channel.Summary AnalogyWithout select (Sequential code): You are standing at a grocery store checkout line. You must wait for the person in front of you to finish paying before you can even look at the next open register.With select: You are standing in the middle of two registers. Whichever cashier yells "Next!" first is the one you walk over to.

	for i:=0; i<2;i++{
		select{
		case chan1Val := <-chan1:
			fmt.Println("Received data from chan1", chan1Val)
		
		case chan2Val := <- chan2:
			fmt.Println("Received data from chan2",chan2Val)
		}
	}




// buffered channels
// ------------------------

// // implementing queue system to send emails to mails stored in db
// 	emailChan := make(chan string, 5)
// 	done := make(chan bool)

// 	go mailSender(emailChan, done)
// 	// email should be loaded from db
// 	for i := 0; i<5; i++ {
// 		emailChan <- fmt.Sprintf("%d@gamil.com",i)
// 	}
// 	fmt.Println("Sending done")

// 	// FIX: Close the channel so the 'for range' loop knows to stop
// 	// You need to add close(emailChan) in your main function right after your loop finishes sending the data. This signals to the for range loop in mailSender that no more data is coming, allowing the loop to terminate gracefully.


// 	close(emailChan)
// 	<- done








// -----------------------
	// emailChan := make(chan string, 100)

	// emailChan <- "1@example.com"
	// emailChan <- "2@example.com"

	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)





// Unbuffered channels
	// --------------------
	// done := make(chan bool)
	// go task(done)

	// <-done  // blocking	



// -------------------------

//  result := make(chan int)
//  go sum(1,2,result)
//  res := <-result  // blocking

//  fmt.Println(res)

// ----------------

	// numChan := make(chan int)
	// go processNum(numChan)  // non blocking

	// for i := 0; i<10; i++ {
	// 	numChan <- rand.Intn(100)
	// }
	

// ------------------

// An unbuffered channel is a channel with zero storage capacity. It acts as a direct, synchronized hand-off between two concurrently running goroutines rather than a mailbox where data can sit and wait.

// Zero Capacity: You initialize it without specifying a size, like make(chan string). It has no internal slots to hold data.
// Synchronous Execution: It forces the sender and the receiver to synchronize their timing. Whichever goroutine arrives at the channel first will pause and wait for the other.
// Guaranteed Delivery: Because it blocks until a hand-off occurs, the sending goroutine gets an immediate guarantee that the receiving goroutine has successfully taken the data the moment the send operation unblocks.

// How the Blocking Behavior Works
// Because there is no storage buffer, operations on the channel dictate the flow of your program:
// 
// Scenario A: The Sender Arrives First
// Goroutine A tries to send data: ch <- "hello"
// Since no one is reading yet, Goroutine A immediately freezes (blocks) on that line.
// It remains frozen until Goroutine B executes <-ch.
// The data transfers directly from A to B, and both goroutines instantly resume working.

// Scenario B: The Receiver Arrives First
// Goroutine B tries to read data: msg := <-ch
// Since no one has sent anything yet, Goroutine B immediately freezes (blocks) on that line.
// It remains frozen until Goroutine A executes ch <- "hello".
// The data transfers directly, and both goroutines move to their next lines of code.



	// messageChan := make(chan string) // channel created

	// messageChan <- "ping"  // blocking, data is being sent to channel

	// msg := <-messageChan   // receiving data from channel in a variable

	// fmt.Println(msg)



// 	This code results in a deadlock because it is trying to send data into an unbuffered channel on a single thread (the main goroutine) without anyone else listening at the exact same moment.
// In Go, an unbuffered channel requires both the sender and the receiver to be ready at the same time to pass the data. It is a hand-off.
// Here is the exact step-by-step breakdown of why your code freezes:
// ## The Step-by-Step Breakdown

//    1. messageChan := make(chan string)
//    * You create an unbuffered channel. Think of this like a telephone call, not a voicemail inbox. For a message to go through, someone must be on the other end answering the call right then.
//    2. messageChan <- "ping"
//    * The main goroutine tries to send "ping".
//       * Because it is unbuffered, the main goroutine pauses (blocks) right here and waits for another goroutine to read the data.
//       * Crucially, execution halts on this line. The code never reaches the next line.
//    3. The Deadlock
//    * The main goroutine is now completely asleep, waiting for a receiver.
//       * But your receiver line (msg := <-messageChan) is on the very next line of the same goroutine. Because the main thread is asleep, it can never execute the code to receive the message.
//       * Go's runtime detects that the only running thread is permanently asleep with no way to wake up, so it crashes with a fatal error: all goroutines are asleep - deadlock!
   
// ------------------------------
// ## How to Fix It
// There are two primary ways to fix this code depending on what you are trying to achieve.
// ## Solution 1: Use a Goroutine (The Idiomatic Way)
// Spin up a separate background thread (goroutine) to handle either the sending or the receiving. This ensures both sides are alive and active simultaneously.

// package main
// import "fmt"
// func main() {
//     messageChan := make(chan string)

//     // Spin up a separate goroutine to send the data
//     go func() {
//         messageChan <- "ping" // This blocks until main reads it below
//     }()

//     // The main thread continues down here and waits to receive
//     msg := <-messageChan 
//     fmt.Println(msg) // Output: ping
// }

// ## Solution 2: Use a Buffered Channel
// If you give the channel a buffer size, it gains a built-in "mailbox" slot. The sender can drop off the message and move on immediately without waiting for a receiver, as long as the buffer isn't full.

// package main
// import "fmt"
// func main() {
//     // Create a buffered channel with room for 1 message
//     messageChan := make(chan string, 1) 

//     // This no longer blocks because there is an open slot in the buffer
//     messageChan <- "ping"  

//     // Read the message out of the buffer
//     msg := <-messageChan   
//     fmt.Println(msg) // Output: ping
// }



}
