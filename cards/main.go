package main

import "fmt"

func main() {
	// Variable
	card := newCard()
	fmt.Println(card)

	// Slice
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")
	
	
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Ace of Spades"
}
