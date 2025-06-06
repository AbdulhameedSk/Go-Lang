package main

import (
	"os"
	"testing"
)

// In Go, *testing.T is a pointer to a testing object provided by Go's testing package
func TestNewDeck(t *testing.T) {
	// Create a new deck of cards
	d := newDeck()

	// Check the length of the deck
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(d))
	}

	// Check the first card in the deck
	if d[0] != "Two of Hearts" {
		t.Errorf("Expected first card to be 'Two of Hearts', but got '%s'", d[0])
	}

	// Check the last card in the deck
	if d[len(d)-1] != "Ace of Spades" {
		t.Errorf("Expected last card to be 'Ace of Spades', but got '%s'", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// Create a new deck of cards
	os.Remove("_decktesting") // Clean up any existing test file
	// This ensures that the test starts with a clean slate
	d := newDeck()

	// Save the deck to a file
	filename := "_decktesting"
	d.saveToFile(filename)

	// Load the deck from the file
	loadedDeck := newDeckFromFile(filename)

	// Check if the loaded deck is equal to the original deck
	if len(loadedDeck) != 52 {
		t.Errorf("Expected loaded deck length of %d, but got %d", len(d), len(loadedDeck))
	}

	for i := range d {
		if d[i] != loadedDeck[i] {
			t.Errorf("Expected card at index %d to be '%s', but got '%s'", i, d[i], loadedDeck[i])
		}
	}

	// Clean up the test file
	os.Remove(filename)
}
