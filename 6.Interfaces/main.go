// Interfaces in Go helps to define a contract that a type must adhere to.
//Interfaces are generic, meaning they can be used with any type that implements the methods defined in the interface.
//Interface are implicit, which means that you don't have to explicitly declare that a type implements an interface.

package main

import "fmt"

//we have a type called bot which is interface.
//if you are type in this program with function called greet, and return string,you are member of type bot.
//since both EnglishBot and SpanishBot implement the greet method, they are considered to be of type bot.
// The bot interface defines a method for greeting.
type bot interface{
	greet() string
	// greet() (string) (string, error)
}
type EnglishBot struct{}
type SpanishBot struct{}

func main(){
	eb:= EnglishBot{}
	sb := SpanishBot{}
	printGreeting(eb) // Output: Hello
	printGreeting(sb) // Output: Hola

}

func (EnglishBot) greet() string {
	return "Hello"
}
func (SpanishBot) greet() string {
	return "Hola"
}
// Greet is an interface that defines a method for greeting.

// func printGreeting(e EnglishBot) {
// 	// The function accepts any type that implements the Greet interface.
// 	fmt.Println(e.greet())
// }

// func printGreeting(s SpanishBot){
// 	fmt.Println(s.greet())
// }

func printGreeting(b bot) {
	// The function accepts any type that implements the bot interface.
	fmt.Println(b.greet())
}