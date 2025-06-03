package main

import "fmt"

//create a new type called deck
//which is a slice of strings
type deck []string //same as type deck extends []string

type card struct {
	// Define a struct to represent a card
	suit string
	rank string
}

//newDeck function creates and returns a new deck of cards essentially a Array of strings
func newDeck() deck {
	// Create a new deck of cards with some initial values
	cardSuits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	cardValues := []string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}
	// Initialize an empty deck
	cards := deck{}
	for _, s := range cardSuits {
		for _, v := range cardValues {
			cards = append(cards, v+" of "+s) // Append each card to the deck})
		}
	}
	// Return the newly created deck
	return cards
}

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
