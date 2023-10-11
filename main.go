package main

import "fmt"

func main() {
	cards := newDeck()

	cards.print()

	hand, remainingDeck := deal(cards, 5)

	fmt.Println("Deal:")
	fmt.Println(">>> Hand")
	hand.print()
	fmt.Println(">>> Remain")
	remainingDeck.print()

	fmt.Println("to string:")
	fmt.Println(cards.toString())

	cards.saveToFile("mycards.txt")

	fmt.Println("Reading from file...")
	newCards := newDeckFromFile("mycards.txt")
	newCards.print()

	fmt.Println("Shuffled:")
	cards.shuffle()
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
