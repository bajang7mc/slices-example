package main

import (
	"fmt"
)

func main() {

	cards := newDeck()
	//cards.print()

	hand, remainingCards := deal(cards, 5)

	fmt.Println("= On Hand =")
	hand.print()
	fmt.Println("= Remaing =")
	remainingCards.print()
}
