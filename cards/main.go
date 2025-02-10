package main

func main() {
	cards := deck{"Ás de Ouros", newCard()}
	cards = append(cards, "Seis de Espadas")

	cards.print()
}

func newCard() string {
	return "Cinco de Ouros"
}
