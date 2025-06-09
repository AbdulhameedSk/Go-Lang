// File: main.go
package main

import "fmt"

/////////////////////////////////////////////////////////
// 1. Interface definition
/////////////////////////////////////////////////////////

// bot is an *implicit* contract: any value that has a
//   method   greet() string
// automatically satisfies the interface.  
// No extra “implements” keyword is required.
type bot interface {
    greet() string
    // ← uncomment an alternative signature if you ever
    //    want to return (string, error) instead.
}

/////////////////////////////////////////////////////////
// 2. Concrete types that satisfy the contract
/////////////////////////////////////////////////////////

type EnglishBot struct{}
type SpanishBot struct{}

// NOTE: Because both receivers are *values* (not pointers),
//       a pointer OR a value of the concrete type will work
//       when passed to printGreeting. If you ever need to
//       mutate the receiver, switch to pointer receivers.

func (EnglishBot) greet() string { return "Hello" }
func (SpanishBot) greet() string { return "Hola" }

/////////////////////////////////////////////////////////
// 3. Reusable code that works with **any** bot
/////////////////////////////////////////////////////////

// printGreeting is blissfully unaware of the concrete type.
// It only knows the caller satisfies the bot interface.
func printGreeting(b bot) {
    fmt.Println(b.greet())
}

/////////////////////////////////////////////////////////
// 4. Demo
/////////////////////////////////////////////////////////

func main() {
    eb := EnglishBot{}
    sb := SpanishBot{}

    printGreeting(eb) // → Hello
    printGreeting(sb) // → Hola
}
