Below are candidate-reported questions, followed by separately labelled preparation questions. No answers or code.

## 1. Machine-coding and practical Go coding: candidate-reported

1. Implement a worker pool where callers enqueue tasks and available workers execute them. **— F5 Networks** [Interview Experience - Discuss - LeetCode](https://leetcode.com/discuss/post/7301168/?utm_source=chatgpt.com)
2. Implement a scheduler that executes N SQL queries while allowing at most five queries to run concurrently. **— DataZip**
3. Implement a document service supporting private/public documents, owner editing, user-specific access grants, and read/write permission checks. **— DataZip**
4. Flatten a deeply nested JSON payload into a structure without nesting. **— DataZip** [LeetCode](https://leetcode.com/discuss/post/6767944/?utm_source=chatgpt.com)
5. Implement a user-management system in Go with MySQL/GORM integration. **— Paytm** [LeetCode](https://leetcode.com/discuss/post/7214484/?utm_source=chatgpt.com)
6. Implement an LRU cache. **— New Relic/Persistent Systems** [Selected - Discuss - LeetCode](https://leetcode.com/discuss/post/6646845/?utm_source=chatgpt.com)
7. Generate Fibonacci numbers using goroutines and channels. **— Cisco** [LeetCode](https://leetcode.com/discuss/post/6323099/?utm_source=chatgpt.com)
8. Use multiple goroutines to print numbers from 1 to 1000 in ascending order. **— MakeMyTrip**
9. Fetch a website’s HTML using Go and count occurrences of a specified word. **— MakeMyTrip** [Rejected - Discuss - LeetCode](https://leetcode.com/discuss/post/1899773/makemytrip-sde-2-bengaluru-march-2022-rejected/?utm_source=chatgpt.com)
10. Coordinate multiple readers and writers accessing a shared byte buffer without races or deadlocks.
11. Implement a producer/consumer queue where consumption begins when queued items reach a specified batch threshold, without channels or WaitGroup.
12. Ensure one function executes only once inside an otherwise continuously running loop. **— Anonymous Go interview report** [reddit.com](https://www.reddit.com/r/golang/comments/176lrjx/struggled_with_advanced_golang_interview/?utm_source=chatgpt.com)

## 2. Additional practical coding: preparation-repository questions

These are repository exercises, not verified company-specific reports.

1. Implement concurrent key lookups and return the first successful result.
2. Implement a concurrent web crawler.
3. Implement a non-blocking cache.
4. Implement a rate limiter.
5. Implement a TTL cache.
6. Implement requests with failover.
7. Merge multiple input channels into one output channel.
8. Implement a concurrency-safe queue.
9. Implement a request-rate tracker. [GitHub](https://github.com/blindlobstar/go-interview-problems?utm_source=chatgpt.com)

## 3. Go technical rounds: candidate-reported questions

### Concurrency

1. How does WaitGroup coordinate goroutine completion?
2. What distinguishes concurrency from parallel execution?
3. How does a worker pool distribute work?
4. How can a done channel signal goroutines to stop? **— DataZip** [LeetCode](https://leetcode.com/discuss/post/6767944/?utm_source=chatgpt.com)
5. How would you collect errors and successful results from goroutines?
6. How would you ingest and process SQS messages using Go workers? **— Anonymous Go interview report** [reddit.com](https://www.reddit.com/r/golang/comments/1nzgkbv?utm_source=chatgpt.com)
7. How would you synchronize concurrent reads and writes to `map[string][]byte` while reducing blocking?
8. How would you coordinate a producer and consumer using a condition variable?
9. How would you ensure initialization happens exactly once? **— Anonymous Go interview report** [reddit.com](https://www.reddit.com/r/golang/comments/176lrjx/struggled_with_advanced_golang_interview/?utm_source=chatgpt.com)

### Go and Kubernetes internals

10. How are Go maps represented and implemented?
11. What gets copied when Go data structures are passed into functions?
12. How does an informer cache work?
13. What happens between a Kubernetes resource update and its reconciler running?
14. How do operators watch resources and reconcile their state?
15. What distinguishes edge-triggered from level-based reconciliation?
16. How do you prevent conflicting reconciliation of the same object?
17. How would you roll out and roll back a new CRD version? **— NVIDIA** [SSE - Discuss - LeetCode](https://leetcode.com/discuss/post/7313633/?utm_source=chatgpt.com)

### Backend and diagnostics

18. How would you investigate a Go memory leak? **— Paytm** [LeetCode](https://leetcode.com/discuss/post/7214484/?utm_source=chatgpt.com)
19. How would you use pprof to investigate memory usage?
20. When would you choose HTTP versus gRPC? **— MakeMyTrip** [Rejected - Discuss - LeetCode](https://leetcode.com/discuss/post/1899773/makemytrip-sde-2-bengaluru-march-2022-rejected/?utm_source=chatgpt.com)

## 4. Additional Go technical preparation questions

### Values, collections and strings

1. Does Go pass arguments by value or by reference?
2. What changes when you pass a pointer?
3. How do arrays and slices differ?
4. What does a slice descriptor contain?
5. How do `len` and `cap` differ?
6. When does `append` allocate another backing array?
7. How can two slices affect each other?
8. How do you make an independent slice copy?
9. How do nil and empty slices differ?
10. What happens when you read from or write to a nil map?
11. Can you rely on map iteration order?
12. Which types can be map keys?
13. How do bytes and runes differ?
14. What does ranging over a string produce? [The Go Programming Language](https://go.dev/ref/spec?utm_source=chatgpt.com)

### Structs, methods, interfaces and generics

15. When should a method use a pointer receiver?
16. How do the method sets of `T` and `*T` differ?
17. How does a type satisfy an interface?
18. Why can an interface containing a nil pointer be non-nil?
19. How do type assertions and type switches differ?
20. How does struct embedding work?
21. How does composition differ from inheritance?
22. What do `any`, `comparable` and type constraints mean?
23. What does `~` mean in a generic constraint?
24. How do loop-variable semantics depend on the module’s Go language version? [The Go Programming Language](https://go.dev/ref/spec?utm_source=chatgpt.com)

### Errors, defer and panic

25. When should you use sentinel errors versus custom error types?
26. How do `errors.Is` and `errors.As` differ?
27. What changes when an error is wrapped using `%w`?
28. How should errors propagate through multiple service layers? [Go Packages](https://pkg.go.dev/errors?utm_source=chatgpt.com)
29. When are deferred-call arguments evaluated?
30. In what order do deferred functions execute?
31. How does panic affect deferred functions?
32. Under what conditions can recover stop a panic?
33. Can one goroutine recover another goroutine’s panic? [The Go Programming Language](https://go.dev/ref/spec?utm_source=chatgpt.com)

### Goroutines and channels

34. What distinguishes a goroutine from an OS thread?
35. When does an unbuffered channel operation block?
36. When does a buffered channel operation block?
37. What happens when you send to a closed channel?
38. What happens when you receive from a closed channel?
39. What happens when you send to or receive from a nil channel?
40. How does select behave when several cases are ready?
41. What effect does a default case have in select?
42. Who should close a channel with multiple producers?
43. How would you prevent leaks when a consumer exits early?
44. How would you preserve result order while processing concurrently?
45. How would you implement backpressure? [The Go Programming Language](https://go.dev/ref/spec?utm_source=chatgpt.com)

### Synchronization

46. What distinguishes a data race, a deadlock and a goroutine leak?
47. When should you use a mutex instead of channel-based ownership?
48. When is RWMutex useful, and when can it hurt?
49. Why must synchronization values not be copied after use?
50. Where should WaitGroup.Add occur relative to starting work?
51. Can a WaitGroup be reused?
52. What happens if a function passed to sync.Once panics?
53. When would you use sync.Cond?
54. When is sync.Map appropriate?
55. When should you use atomic operations?
56. Why can individually synchronized operations still produce an incorrect check-then-act sequence? [Go Packages](https://pkg.go.dev/sync?utm_source=chatgpt.com)

### Context and cancellation

57. How do Background and TODO differ?
58. How do WithCancel, WithTimeout and WithDeadline differ?
59. Why should you invoke the returned cancel function?
60. Does canceling context forcibly stop a goroutine?
61. How does cancellation propagate to child contexts?
62. How should context reach downstream HTTP and database calls?
63. What belongs in context values?
64. Why should context values not carry optional configuration?
65. How would you make channel operations cancellation-aware?
66. How would you distinguish cancellation from deadline expiry? [Go Packages](https://pkg.go.dev/context?utm_source=chatgpt.com)

### HTTP and backend coding

67. How do Handler and HandlerFunc differ?
68. How would you implement HTTP middleware?
69. How would you validate and limit a JSON request body?
70. Why should HTTP clients and transports be reused?
71. Which timeouts would you configure for an HTTP client and server?
72. Why must response bodies be closed?
73. How can unread response bodies affect connection reuse?
74. How would you propagate request cancellation?
75. How would you shut down a server while allowing active requests to finish?
76. Why might main need to wait after ListenAndServe returns during shutdown? [Go Packages](https://pkg.go.dev/net/http?utm_source=chatgpt.com)

### Database access

77. Is sql.DB a connection or a connection pool?
78. How would you configure pool limits and connection lifetimes?
79. How do QueryContext and ExecContext differ?
80. Why must rows be closed and rows.Err checked?
81. How would you implement commit and rollback correctly?
82. What can cause connection-pool exhaustion?
83. What happens when context cancellation occurs during a transaction? [Go Packages](https://pkg.go.dev/database/sql?utm_source=chatgpt.com)

### Testing and debugging

84. How would you write table-driven tests and subtests?
85. How would you test an HTTP handler?
86. How would you test a client against a slow or failing server?
87. How would you test cancellation without arbitrary sleeps?
88. How would you test time-dependent behavior deterministically?
89. What can a fake dependency hide?
90. What can the race detector fail to detect?
91. How would you benchmark an operation and measure allocations?
92. When would you use fuzz testing? [Go Packages](https://pkg.go.dev/testing?utm_source=chatgpt.com)

### Performance and tooling

93. When would you inspect CPU, heap or goroutine profiles?
94. How would you locate a blocked or leaking goroutine?
95. How would you investigate excessive allocations?
96. How would you distinguish high CPU usage from time spent waiting?
97. When would execution tracing be more useful than a CPU profile? [The Go Programming Language](https://go.dev/doc/diagnostics?utm_source=chatgpt.com)
98. What belongs in go.mod versus go.sum?
99. What does go mod tidy change?
100. How would you pin, upgrade or replace a dependency?
101. How does the module’s Go version affect the code you can build? [The Go Programming Language](https://go.dev/doc/modules/managing-dependencies?utm_source=chatgpt.com)