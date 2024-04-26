package main

import "fmt"

func main() {
	//Create slice of cards using deck type
	cards := newDeck()
	// hand, remainingDeck := deal(cards, 3)
	fmt.Println("-------------Before shuffle-------------")
	cards.print()

	cards.shuffle()
	fmt.Println("--------------After shuffle-----------")
	cards.print()

}
