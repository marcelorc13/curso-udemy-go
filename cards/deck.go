package main

import (
	"fmt"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Espadas", "Ouros", "Copas", "Paus"}
	cardValues := []string{"Ás", "Dois", "Três", "Quatro", "Cinco", "Seis", "Sete", "Oito", "Nove", "Dez", "Valete", "Dama", "Rei"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" de "+suit)
		}
	}

	return cards

}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, "-", card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toBytseSlice() []byte {
	stringDeck := strings.Join([]string(d), ",")
	byteDeck := []byte(stringDeck)
	return byteDeck
}

func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, d.toBytseSlice(), 0666)
}
