package main

// User defined function
// func newCard() string {
// 	return "Five of Diamonds"
// }

func main() {
	// Decalring a variable
	// var card string = "Ace of Spades"

	// Calling a Function
	// card := newCard()
	// fmt.Println(card)

	// Slices
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")

	// For loop over cards
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	cards := newDeck()
	cards.shuffle()
	cards.print()

	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// remainingDeck.print()
}
