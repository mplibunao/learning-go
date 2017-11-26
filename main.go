package main

func main() {
	// cards := newDeck()
	// cards.saveToFile("deck")

	cards := newDeckFromFile("deck1")
	cards.print()
}
