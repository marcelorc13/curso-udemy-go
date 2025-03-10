package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	// cards.saveToFile("minhas_cartas")
}
