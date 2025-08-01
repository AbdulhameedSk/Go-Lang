# Go Programming Refresher: README

This README is a beginner-friendly guide and Q&A for revisiting Go (Golang), covering the core language concepts and addressing some of the most common questions with code examples. 

## Table of Contents

- Getting Started
- Variables, Constants, and Functions
- Control Structures: If, For, Switch
- Slices, Arrays, and Multi-Return Functions
- Structs
- Maps
- Interfaces & Polymorphism
- Pointers, Pass-by-Value, Pass-by-Reference
- Slices, Maps, Channels: Internal Behavior
- Questions & Answers Section

## Getting Started

Go is a statically typed, compiled language known for simplicity and excellent concurrency. 

**Hello, World! Example:**
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```
- Save as `hello.go` and run: `go run hello.go`

## Variables, Constants, and Functions

- Variables: `var name string = "Go Learner"` or shorthand: `age := 30`
- Constants: `const Pi = 3.14`

**Example:**
```go
func add(x int, y int) int {
    return x + y
}
```

## Control Structures

**If, For, Switch:**
```go
numbers := []int{1, 2, 3, 4}
for i := 0; i < len(numbers); i++ {
    if numbers[i] % 2 == 0 {
        fmt.Println(numbers[i], "is even")
    } else {
        fmt.Println(numbers[i], "is odd")
    }
}
```

**Multi-Return Function:**
```go
func split(sum int) (int, int) {
    return sum / 2, sum % 2
}
```

## Structs

Custom types grouping fields:
```go
type Person struct {
    Name string
    Age  int
}

p := Person{Name: "Alice", Age: 28}
fmt.Println(p.Name, p.Age)
```

## Maps

Key-value stores:
```go
scores := make(map[string]int)
scores["Alice"] = 95
value, exists := scores["Charlie"]
if !exists { fmt.Println("Charlie not found") }
delete(scores, "Bob")
```

## Interfaces & Polymorphism

- An interface is a contract: any type implementing the required methods automatically satisfies the interface.

**Example:**
```go
type Shape interface {
    Area() float64
}

type Rectangle struct{ Width, Height float64 }
func (r Rectangle) Area() float64 { return r.Width * r.Height }

type Circle struct{ Radius float64 }
func (c Circle) Area() float64 { return 3.14 * c.Radius * c.Radius }

func printArea(s Shape) {
    fmt.Println("Area:", s.Area())
}
```

**How it works:**  
- Assign a concrete type (like Rectangle) to an interface variable.
- Calling an interface method (like `Area()`) runs the concrete type’s version, based on the variable's actual type (not function arguments).
- This is known as *dynamic dispatch*.

**Clarification (from Q&A):**  
- The interface does NOT pick a function implementation based on the function arguments. It is based on the *type* held in the interface at runtime.
- In Go, there is **no function overloading**—method selection is type-driven.

## Pointers, Pass-by-Value, Pass-by-Reference

**Go always passes by value!**  
But how the value behaves depends on the type:

- For primitives (int, float, struct, array), the value is copied. Original is not changed.

```go
func modify(n int) {
    n = 100
}

x := 10
modify(x)
fmt.Println(x) // Prints 10, unchanged
```

- Use a pointer to modify the caller’s variable:

```go
func modify(n *int) {
    *n = 100
}

x := 10
modify(&x)
fmt.Println(x) // Prints 100, changed
```

## Slices, Maps, Channels: Internal Behavior

- Slices, maps, channels are special: a small header is copied, but they reference shared backing storage.
- Changing their contents inside a function changes the original.
```go
func modifySlice(s []int) {
    s[0] = 42
}
nums := []int{1, 2, 3}
modifySlice(nums)
fmt.Println(nums) // Prints [42 2 3]
```

## Q&A from Session

### Q: Why is "x" unchanged after passing it to a function?

**A:** Because Go passes arguments by value—function gets a copy, not the original.

### Q: What’s the difference between pass-by-value and pass-by-reference?

- Pass-by-value (default): Function works on a copy.
- To modify original, pass a pointer (“reference” the variable’s address).
- Slices, maps, channels allow modifying their underlying data even when passed by value.

### Table Summarizing Behavior

| Type                         | Passed By    | Modify Original? | Example Needed?        |
|------------------------------|--------------|------------------|------------------------|
| int, float, struct, array    | Value        | No (unless ptr)  | Use pointer to alter   |
| slice, map, channel          | Value (head) | Yes (data shared)| No pointer needed      |

```https://gabrieleromanato.name/go-data-types-passed-by-value-and-data-types-passed-by-reference#:~:text=integers%20(int)%2C%20floats%20(float64)%2C%20strings%20(string)%2C%20and%20booleans%20(bool)%2C%20are%20passed%20by%20value```

## Extra Examples: Error Handling, Pointer Basics

**Error Handling:**
```go
func divide(x, y float64) (float64, error) {
    if y == 0 {
        return 0, errors.New("division by zero")
    }
    return x / y, nil
}
```

**Pointer:**
```go
x := 42
ptr := &x
*ptr = 100
fmt.Println(x) // 100
```

# Go Routines and Channels Explained

## What are Goroutines?

**Goroutines** are lightweight threads managed by the Go runtime. They allow you to run functions concurrently without the overhead of traditional operating system threads. Think of them as very cheap "mini-threads" that can run thousands simultaneously.

### Basic Goroutine Example

```go
package main

import (
    "fmt"
    "time"
)

func sayHello() {
    for i := 0; i < 5; i++ {
        fmt.Println("Hello from goroutine!")
        time.Sleep(time.Millisecond * 100)
    }
}

func main() {
    // Start a goroutine
    go sayHello()  // The 'go' keyword makes this run concurrently
    
    // Main function continues
    for i := 0; i < 5; i++ {
        fmt.Println("Hello from main!")
        time.Sleep(time.Millisecond * 100)
    }
    
    // Wait a bit to let the goroutine finish
    time.Sleep(time.Second)
}
```

**Output:** You'll see messages from both the main function and the goroutine mixed together, showing they're running concurrently.

## What are Channels?

**Channels** are Go's way to communicate between goroutines safely. Think of them as pipes that can send data from one goroutine to another. They solve the problem of sharing data between concurrent processes.

### Basic Channel Example

```go
package main

import "fmt"

func main() {
    // Create a channel that can send/receive strings
    ch := make(chan string)
    
    // Start a goroutine that sends data to the channel
    go func() {
        ch <- "Hello from goroutine!"  // Send data to channel
    }()
    
    // Receive data from the channel
    message := <-ch  // Receive data from channel
    fmt.Println(message)
}
```

## Combining Goroutines and Channels

Here's where it gets powerful - using them together:

```go
package main

import (
    "fmt"
    "time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
        time.Sleep(time.Second) // Simulate work
        results <- job * 2      // Send result back
    }
}

func main() {
    jobs := make(chan int, 5)
    results := make(chan int, 5)
    
    // Start 3 workers
    for i := 1; i <= 3; i++ {
        go worker(i, jobs, results)
    }
    
    // Send 5 jobs
    for i := 1; i <= 5; i++ {
        jobs <- i
    }
    close(jobs) // Important: close the channel when done sending
    
    // Collect results
    for i := 1; i <= 5; i++ {
        result := <-results
        fmt.Printf("Result: %d\n", result)
    }
}
```

## Channel Direction and Buffering

### Buffered vs Unbuffered Channels

```go
package main

import "fmt"

func main() {
    // Unbuffered channel - blocks until someone receives
    unbuffered := make(chan string)
    
    // Buffered channel - can hold 2 values before blocking
    buffered := make(chan string, 2)
    
    // This would block forever (deadlock) with unbuffered
    // unbuffered <- "Hello" // DON'T DO THIS
    
    // This works fine with buffered
    buffered <- "Hello"
    buffered <- "World"
    
    fmt.Println(<-buffered) // "Hello"
    fmt.Println(<-buffered) // "World"
}
```

### Channel Directions

```go
package main

import "fmt"

// send-only channel
func sender(ch chan<- string) {
    ch <- "Hello"
    ch <- "World"
    close(ch)
}

// receive-only channel
func receiver(ch <-chan string) {
    for msg := range ch {
        fmt.Println("Received:", msg)
    }
}

func main() {
    ch := make(chan string)
    
    go sender(ch)
    receiver(ch)
}
```

## Practical Example: Web Scraper

Here's a real-world example showing the power of goroutines and channels:

```go
package main

import (
    "fmt"
    "time"
)

// Simulate fetching a URL
func fetchURL(url string, ch chan<- string) {
    // Simulate network delay
    time.Sleep(time.Millisecond * 500)
    ch <- fmt.Sprintf("Fetched: %s", url)
}

func main() {
    urls := []string{
        "https://example.com",
        "https://google.com",
        "https://github.com",
        "https://stackoverflow.com",
    }
    
    ch := make(chan string)
    
    // Start all fetches concurrently
    for _, url := range urls {
        go fetchURL(url, ch)
    }
    
    // Collect all results
    for i := 0; i < len(urls); i++ {
        result := <-ch
        fmt.Println(result)
    }
}
```

## Select Statement for Multiple Channels

The `select` statement lets you wait on multiple channels:

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    go func() {
        time.Sleep(1 * time.Second)
        ch1 <- "Channel 1"
    }()
    
    go func() {
        time.Sleep(2 * time.Second)
        ch2 <- "Channel 2"
    }()
    
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Println("Received:", msg1)
        case msg2 := <-ch2:
            fmt.Println("Received:", msg2)
        case <-time.After(3 * time.Second):
            fmt.Println("Timeout!")
        }
    }
}
```

## Key Points to Remember

- **Goroutines** are started with the `go` keyword
- **Channels** are created with `make(chan Type)` or `make(chan Type, bufferSize)`
- Use `<-` to send (`ch <- value`) and receive (`value := <-ch`)
- Always `close()` channels when you're done sending
- Use `select` for handling multiple channels
- Buffered channels don't block until full; unbuffered channels block immediately

## Common Patterns

1. **Fan-out/Fan-in**: Distribute work to multiple workers, collect results
2. **Pipeline**: Chain goroutines together with channels
3. **Worker Pool**: Fixed number of workers processing jobs from a queue
4. **Timeout**: Use `time.After()` with `select` for timeouts

This concurrent programming model makes Go excellent for building scalable, concurrent applications like web servers, data pipelines, and distributed systems.


## Testing in Go

Testing is built into Go’s standard library, making it easy to write and run tests for your code. Here’s everything you need to get started:

### Basic Test Structure

- **Test files** end with `_test.go`.
- Tests use the `testing` package.
- Each test function starts with `Test`, takes a `*testing.T`, and does not return anything.

**Example: Testing an Add Function**

Suppose you have this function in `math.go`:

```go
package mymath

func Add(x, y int) int {
    return x + y
}
```

Write a test in `math_test.go`:

```go
package mymath

import "testing"

func TestAdd(t *testing.T) {
    got := Add(2, 3)
    want := 5
    if got != want {
        t.Errorf("Add(2, 3) = %d; want %d", got, want)
    }
}
```

- If `got != want`, the test fails and prints the message.

### Running Tests

- Run all tests in the current directory:
  ```
  go test
  ```
- To see more output, use:
  ```
  go test -v
  ```

### Table-Driven Tests

For multiple cases, use a loop and a slice of structs:

```go
func TestAdd(t *testing.T) {
    tests := []struct{
        x, y, want int
    }{
        {1, 2, 3},
        {0, 0, 0},
        {-1, -1, -2},
    }
    for _, tc := range tests {
        got := Add(tc.x, tc.y)
        if got != tc.want {
            t.Errorf("Add(%d, %d) = %d; want %d", tc.x, tc.y, got, tc.want)
        }
    }
}
```

### Test Coverage & More

- Measure test coverage with:
  ```
  go test -cover
  ```
- Benchmark performance with functions starting `Benchmark` and using `*testing.B`.

### Example Directory Structure

```
mymodule/
  math.go         // your functions
  math_test.go    // tests
```

### Mocking and Interfaces

When testing code that depends on external systems, define interfaces and provide test implementations (mocks) in your test files.

**Tip:** Keep your functions pure (return results from inputs, avoid side effects) for easiest testing.

Let me know if you want to see table-driven examples for structs, want to test methods, or use advanced features like setup/teardown!