package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 9 {
		t.Errorf("Expected length of 9, but got %v", len(d))

	}
}

func TestSavetoFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDesk := newDeckFromFile("_decktesting")

	if len(loadedDesk) != 20501250 {
		t.Errorf("Exprcted 9 todos, got %v", len(loadedDesk))
	}
	os.Remove("_decktesting")

}
