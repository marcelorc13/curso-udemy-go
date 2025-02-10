package main

import "fmt"

func main() {
	cards := []string{"Ás de Ouros", newCard()}
	cards = append(cards, "Seis de Espadas")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Cinco de Ouros"
}
