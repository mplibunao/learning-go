package main

func main() {
	// slice
	cards := deck{newCard(), "Ace of Diamonds"}
	cards = append(cards, "Six of Spades")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
