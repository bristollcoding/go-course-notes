package main

import (
	"os"
	"reflect"
	"testing"
)

// test functions must start with uppercase
// test for newDeck function
func TestNewDeck(t *testing.T) {
	expectedLength := 27
	cards := newDeck()
	//check length
	if len(cards) != expectedLength {
		//Error description
		t.Errorf("Expected deck length of %v, but got %v", expectedLength, len(cards))
		// t.Fail()
	}

	//check first item
	firstCard := "Ace of Spades"
	if cards[0] != firstCard {
		t.Errorf("Expected %s but got %s", firstCard, cards[0])
	}

}

func TestSaveAndReadDeckFromFile(t *testing.T) {
	deckFileName := "_decktest"
	//delete any file that contain _decktest
	os.Remove(deckFileName)
	//create deck and save to _decktest file
	writeDeck := newDeck()

	writeDeck.saveToFile(deckFileName)
	//read file _decktest to create a deck
	readDeck, err := deckFromFile(deckFileName)
	if err != nil {
		t.Errorf("Error reading file: %s", err)
	}

	if !reflect.DeepEqual(writeDeck, readDeck) {
		t.Errorf("Deck saved and read expected to be equial but are not. \nwriteDeck: %s \nreadDeck: %s", writeDeck, readDeck)
	}
	//delete any file that contain _decktest
	os.Remove(deckFileName)
}

func TestDeal(t *testing.T) {
	d := newDeck()
	handSize := 3

	hand, _ := deal(d, handSize)

	if len(hand) != handSize {
		t.Errorf("Expected size %v but was %v", handSize, len(hand))
	}

}

func TestShuffle(t *testing.T) {
	d := newDeck()

	d.shuffle()

	if reflect.DeepEqual(d, newDeck()) {
		t.Errorf("Deck expected to be different after shuffle but was equal")
	}
}
