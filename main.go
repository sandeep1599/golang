package main

import "fmt"

func main() {

	cards := newDeck()

	// cards.shuffle()
	cards.Print()

	fmt.Println(deal(cards,4))
}
