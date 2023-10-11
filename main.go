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

}

func newCard() string {
	return "Five of Diamonds"
}
