package main

import (
	"os"
	"testing"
	"github.com/josephpenafiel/cards/deck"
)

func TestNewDeck(t *testing.T) {
	d := deck.NewDeck()

	if len(d) != 52 {
		t.Errorf("Expected 52 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spaces, but got %v", d[0])
	}

	if d[len(d) - 1] != "K of Clubs" {
		t.Errorf("Expected last card to be K of Clubs, but got %v", d[0])
	}
}

func TestSaveToFileAndLoadFromFile(t *testing.T) {
	os.Remove("_decktest.txt")
	d := deck.NewDeck()
	d.SaveToFile("_decktest.txt")

	loadedDeck := deck.NewDeckFromFile("_decktest.txt")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards, but got %v", len(loadedDeck))
	}

	os.Remove("_decktest.txt")
}