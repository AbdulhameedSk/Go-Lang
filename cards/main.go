//Arrays have a fixed size, while slices are dynamic and can grow or shrink

package main

func main() {
	//Static type Language variables declared using data types
	// cards:=[]string{"Ace of Spades", "Two of Hearts", "Three of Clubs", newCard()}
	cards := deck{"Ace of Spades", "Two of Hearts", "Three of Clubs", newCard()}
	cards = append(cards, "Six of Spades") // Adding a new card to the slice
	cards.print()                          // now d is cards

}
func newCard() string {
	// var card string = "Ace of Spades"
	return "Five of Diamonds" // Using short variable declaration
}
