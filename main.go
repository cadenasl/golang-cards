package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	 cards:= newDeck()
	 cards.shuffleDeck()
	// hand,remainingDeck:=deal (cards,5)

	// cards.print()

	//  fmt.Println(cards)
	//  remainingDeck.print()
	//  hand.print()
	// cards := newDeckFromFile("my_cards")
	fmt.Println(cards)
	cards.print()
}
