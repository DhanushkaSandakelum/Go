package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of "deck"
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// Use _ to skip the variables that are not going to use
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	// Seperate the deck into two parts
	return d[:handSize], d[handSize:]
}

// Receiver Function to print cards
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	byte_slice, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	slice_of_strings := strings.Split(string(byte_slice), ",")

	return deck(slice_of_strings)
}

func (d deck) shuffle() {
	// Generate a new source with a seed on nano time of now
	source := rand.NewSource(time.Now().UnixNano())
	// Use than random number generator to generate new random number
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		// One line swap
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
