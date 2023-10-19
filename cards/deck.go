package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create new type of deck
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(fmt.Sprintf("[%d]", i+1), card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"♠️", "♥️", "♦️", "♣️"}

	cardValues := []string{"1"}
	for i := 1; i <= 12; i++ {
		cardValues = append(cardValues, fmt.Sprint(i))
	}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, fmt.Sprintf("%s %s", suit, value))

		}
	}

	return cards
}

func deal(cards deck, handSize int) (deck, deck) {
	return cards[:handSize], cards[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	cards := strings.Split(string(bs), ",")

	return deck(cards)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
