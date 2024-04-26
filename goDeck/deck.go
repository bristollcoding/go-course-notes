package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts"}
	cardNumbers := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardNumbers {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]

}

// print function
// func (receiver) make print available for receiver type
func (d deck) print() {
	//iterate over all cards inside cards slice
	for _, card := range d {
		fmt.Println(card)
	}
}

// Conver a deck into a string containing all cards info
func (d deck) toString() string {
	//Convert back to []string from deck type
	return strings.Join([]string(d), ",")

}

func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), os.ModePerm)
}

func deckFromFile(filePath string) (deck, error) {
	bytes, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Println("Error reading deck from file: "+filePath, err)
		return nil, err
	}

	return deck(strings.Split(string(bytes), ",")), nil
}

func (d deck) shuffle() {
	for i := range d {
		//Generate random index
		randIndex := rand.Intn(len(d))
		// //get card from index
		// randCard := d[randIndex]
		// //put current card into random card position
		// d[randIndex] = card
		// //put random card into current card place
		// d[i] = randCard

		//reduced using go multivariable asigment
		d[i], d[randIndex] = d[randIndex], d[i]
	}
}
