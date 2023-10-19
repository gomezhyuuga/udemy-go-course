package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	expectedLength := 52

	if len(d) != expectedLength {
		t.Errorf("Expected deck length of %v but got %v", expectedLength, len(d))
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	d := newDeck()
	testFilename := "_decktesting"

	// Remove old files
	os.Remove(testFilename)

	d.saveToFile(testFilename)

	d2 := newDeckFromFile(testFilename)

	if len(d2) != 52 {
		t.Errorf("File contained a different deck, got %v cards", len(d2))
	}

	// Remove old files
	os.Remove(testFilename)
}
