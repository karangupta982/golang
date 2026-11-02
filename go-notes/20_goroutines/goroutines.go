package main

import (
	"fmt"
	"sync"
)


// To understand what a "lightweight thread" (a goroutine) is, you first need to look at what a standard Operating System (OS) thread is, and why Go decided to move away from it. [1, 2] 
// Here is the breakdown of how OS threads work, how goroutines differ, and how they interact. [3, 4] 
// ------------------------------
// ## 1. What is a Traditional OS Thread?
// An OS Thread (often called a kernel thread) is a path of execution managed directly by your computer's Operating System (Windows, macOS, or Linux). [5, 6, 7, 8, 9] 

// * The OS Scheduler manages these threads, constantly switching between them so your CPU can execute multiple programs at once. [10, 11, 12, 13] 
// * The Problem with OS Threads: They are resource-heavy.
// * Memory cost: Every OS thread allocated by the system takes up a fixed, relatively large chunk of memory stack space (typically 1 to 2 Megabytes).
//    * Creation cost: Asking the operating system to create a thread requires an expensive system handshake.
//    * Switching cost: Moving the CPU from processing one OS thread to another (a context switch) is slow because the CPU has to save and load massive register states. [14, 15, 16, 17, 18] 

// Because of this, if you try to spin up 10,000 OS threads simultaneously, your computer will quickly run out of memory and crash. [19, 20] 
// ------------------------------
// ## 2. What makes a Goroutine a "Lightweight Thread"?
// A goroutine is a user-space thread. This means the Operating System doesn't even know goroutines exist. They are created, destroyed, and managed entirely by Go's internal runtime engine. [21, 22, 23, 24, 25] 
// Go makes them "lightweight" by radically optimizing memory and execution:

// * Tiny Memory Footprint: A goroutine starts with a dynamic stack size of only 2 Kilobytes (compared to 2 Megabytes for an OS thread). If a goroutine needs more space to run a complex calculation, Go automatically grows and shrinks its memory allocation on the fly. [26, 27, 28, 29, 30] 
// * Incredibly Fast Context Switching: Because Go manages the switching internally without talking to the OS kernel, switching between goroutines takes a fraction of the time it takes to switch OS threads. [31, 32] 

// This efficiency means a standard laptop can easily run hundreds of thousands of goroutines at the exact same time without breaking a sweat. [33, 34] 
// ------------------------------
// ## 3. How Goroutines talk to OS Threads: The M:N Scheduler
// Your CPU can still only execute actual OS threads. So, how do goroutines get executed? [35] 
// Go uses an architectural model called the M:N Scheduler. It maps M goroutines onto N physical OS threads. [36, 37, 38] 

//   [ Goroutines ]         G1   G2   G3   G4   G5   G6   (Thousands of them)
//                            \   |   /      \   |   /
//   [ Go Scheduler ]          ───┬───        ───┬───     (Go manages the queue)
//                                │              │
//   [ OS Threads ]               M1             M2       (Only a few threads)
//                                │              │
//   [ CPU Cores ]             Core 0         Core 1      (Actual physical hardware)


//    1. Go creates a very small number of actual OS threads (usually matching the number of physical CPU cores on your machine).
//    2. The Go runtime multiplexes (shuffles) thousands of active goroutines onto those few OS threads.
//    3. If Goroutine G1 blocks (for example, waiting for a network response or a channel read), Go's scheduler instantly unhooks G1 from the OS thread, moves it to the side, and drops G2 onto the thread so the CPU never sits idle. [39, 40, 41, 42, 43] 

// ## Summary Analogy

// * An OS Thread is like a massive tour bus. It takes a lot of fuel to start, requires a commercial license, and takes up a huge amount of space on the road. You can only fit a few buses on a street.
// * A Goroutine is like a passenger on that bus.
// * Go's Scheduler is the bus driver. Instead of buying 10,000 separate tour buses for 10,000 people, Go buys 4 large buses (OS threads) and manages how passengers (goroutines) swap seats efficiently so the buses are always completely full and moving. [44] 

// Would you like to explore how Go's scheduler uses "Work Stealing" to balance these goroutines across different CPU cores automatically?
// ---
// Thank you so much! I appreciate your feedback. Let’s dive straight into Work Stealing, which is the secret sauce behind how Go’s scheduler maximizes your CPU efficiency.
// To understand work stealing, we need to look at how the Go scheduler is structured. It uses three core entities, commonly referred to as the G-M-P model:

// * G (Goroutine): The lightweight thread of execution (your code).
// * M (Machine): A physical OS thread created and managed by the operating system.
// * P (Processor): A logical context or resource required to execute Go code. The number of Processors exactly matches the number of CPU cores on your computer.

// ------------------------------
// ## The Architecture: Local Run Queues
// Every logical processor (P) maintains its own private checklist of tasks called a Local Run Queue. This queue holds the goroutines (G) that are waiting to be executed on that specific CPU core. There is also a single Global Run Queue that holds leftover goroutines.

//   [Core 0 / P1] ───> [ Local Queue: G1, G2, G3 ] ───> Executing on Thread M1
//   [Core 1 / P2] ───> [ Local Queue: G4, G5, G6 ] ───> Executing on Thread M2

// In a perfect world, every core would finish its tasks at the exact same time. But in reality, some tasks take longer than others.
// ------------------------------
// ## The Problem: Thread Starvation
// Imagine Core 0 gets stuck executing a complex mathematical formula that takes a long time. Meanwhile, Core 1 finishes its 3 small tasks instantly.
// Without work stealing, Core 1 would run out of work, its OS thread would go to sleep, and half of your laptop's computing power would sit completely idle while Core 0 struggles to finish its heavy backlog.
// ------------------------------
// ## The Solution: How Work Stealing Saves the Day
// When a logical processor (P) completely finishes all the goroutines in its local run queue, it doesn't go to sleep. Instead, it enters Work Stealing Mode and executes a strict, aggressive search algorithm to find work elsewhere:

// Step 1: Check Local Queue (Empty)
//         ↓
// Step 2: Check Global Queue (Empty)
//         ↓
// Step 3: Check Network Poller (Empty)
//         ↓
// Step 4: STEAL FROM PEERS! ───> Steals half the queue from a busy processor


//    1. Stealing from Peer Cores (The 51% Rule): The idle processor picks another random processor (P) that is currently slammed with work. It reaches into that busy processor's local queue, steals exactly half of its waiting goroutines, and moves them to its own queue.
//    2. Immediate Resumption: The idle CPU core instantly begins executing those stolen goroutines.

// [Core 0 (Busy)]  ───> [ G1, G2, G3, G4 ]  
//                          |   │
//                          └─  ┴─> [Core 1 (Idle)] steals G3 and G4!

// ------------------------------
// ## Why is this design so efficient?

// * No Central Lock (Locks are Slow): If thousands of goroutines sat in a single central line, every CPU core would have to fight for a "lock" to grab the next task. This friction slows down multi-core processors. By giving each core its own local queue, they can work independently 99% of the time without interfering with each other.
// * Zero Hardware Waste: It ensures that as long as there is even one task left to do in your program, every single core on your computer will be running at 100% capacity to get it done.




// import (
// 	"fmt"
// 	"time"
// )

// // lightweight threads, multi threading, concurrent threads

// func task(num int){
// 	 fmt.Print("Doing task",num)
// }

// func main(){
// 	// It is a blocking nature, execution of each task is blocking, so multiple task would not be happening at the same time but it will be sequential
// 	// for i := 0; i<=10; i++{
// 	// 	task(i)
// 	// }

// 	// These tasks are non blocking but each call would be scheduled by scheduler to run as a thread, here the order would not be sequential but it would be random
// 	for j := 0; j<10; j++ {
// 		// go task(j)

// 		go func(j int){
// 			fmt.Println(j)
// 		}(j)
// 	}

// 	// since all the calls are being made simulataneously and for loops finishes very fast and then main functiona get end so each thread got successful but since main function got end so nothing printed on console by background threads
// 	time.Sleep((time.Second * 2))

// }

// --------
// Wait Group



func task(task_id int, wg *sync.WaitGroup){
	defer wg.Done() // defer runs as a cleanup function like in react after this function gets completed, here it will decrease wg count by 1 when function gets complete
	fmt.Println("Doing task", task_id)
}

func main(){
	var wg sync.WaitGroup

	for i := 0; i<10; i++{
		wg.Add(1)
		go task(i, &wg)
	}

	wg.Wait() // will wait untill all wg get 0 value
}



// Direct Comparison (Concurrency vs Parallelism)

// Concurrency 🔀																													Parallelism 🏃‍♂️🏃‍♂️
// Composition & Design: Structuring your program into independent tasks that can overlap in time.									Hardware Execution: Literally running multiple tasks at the exact same physical millisecond.

// Can run on a single CPU core (or multiple).																						Requires multiple CPU cores or processors.

// The system switches back and forth between tasks rapidly (context switching), creating the illusion of simultaneous execution.	Tasks run independently on separate hardware tracks simultaneously.

// Efficiency & Structure: Managing many things (like I/O waits, user requests) without getting blocked.							Speed & Performance: Heavy, computational throughput by spreading the workload.



