package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// create a new type called deck
// which is a slice of strings
type deck []string //same as type deck extends []string

type card struct {
	// Define a struct to represent a card
	suit string
	rank string
}

// newDeck function creates and returns a new deck of cards essentially a Array of strings
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
// The actual copy of cards is avalilable in deck as d in the function
// Every variable of type deck will have access to the print method
func (d deck) print() {
	// Loop through the deck and print each card
	for i, card := range d {
		// Print the index and the card
		fmt.Println(i+1, card)
	}
}

// deck is a slice of strings
func deal(d deck, handSize int) (deck, deck) {
	// Split the deck into two parts: hand and remaining cards
	return d[:handSize], d[handSize:]

}

// function to convert a deck to a string to single string
func (d deck) toString() string {
	// Convert the deck to a single string with each card separated by a comma
	// return fmt.Sprint([]string(d))
	str := strings.Join([]string(d), ", ")
	return str
}

// 0 6 6 6
// │ │ │ └── Others' permissions
// │ │ └──── Group's permissions
// │ └────── Owner's permissions
// └──────── Special bits (like setuid, setgid, sticky bit)
// Each digit is a combination of:
// Octal	Binary	Meaning
// 0	000	No permissions
// 1	001	Execute only
// 2	010	Write only
// 4	100	Read only
// 6	110	Read + Write
// 7	111	Read + Write + Execute

// So 666 is:

// Owner → Read + Write

// Group → Read + Write

// Others → Read + Write

// That means anyone can read and write the file, but no one can execute it.

func (d deck) saveToFile(filename string) error {
	// Save the deck to a file
	// Convert the deck to a string
	data := d.toString()
	// Write the string to the file
	return os.WriteFile(filename, []byte(data), 0666)
}

// loadDeckFromFile function reads a deck from a file and returns it
func newDeckFromFile(filename string) deck {
	// Read the file content
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1) // Exit the program if there's an error
	}
	// Split the file content into individual cards
	s := strings.Split(string(bs), ", ")
	// Return the new deck
	return deck(s)
}

// ... rest of your code ...

func (d deck) shuffle() {
	// Shuffle the deck of cards
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i] // Swap cards
	}
}
