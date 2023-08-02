package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 20, But got %v \n", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("expected last card is `Ace of Spade`,But it is %v \n", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("expected last card is `Ace of Spade`,But it is %v \n", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {

		t.Errorf("Expected 16 cards in the deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
