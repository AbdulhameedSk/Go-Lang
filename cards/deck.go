package main

import "fmt"

//create a new type called deck
//which is a slice of strings
type deck []string //same as type deck extends []string

// Reciever sets up methods on variable we create

//The actual copy of cards is avalilable in deck as d in the function
//Every variable of type deck will have access to the print method
func (d deck) print() {
	// Loop through the deck and print each card
	for i, card := range d {
		// Print the index and the card
		fmt.Println(i+1, card)
	}
}
