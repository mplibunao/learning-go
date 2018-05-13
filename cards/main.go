package main

func main() {
	// cards := newDeck()
	// cards.saveToFile("deck")

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
