package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("O deck deve ter 52 cartas, porem ele possui %v", len(d))
	}

	if d[0] != "Ás de Espadas" {
		t.Errorf("É esperado que primeira carta seja um Ás de Espadas, porem temos %v", d[0])
	}

	if d[len(d)-1] != "Rei de Paus" {
		t.Errorf("É esperado que ultima carta seja um Rei de Paus, porem temos %v", d[0])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loaddedDeck := newDeckFromFile("_decktesting")

	if len(loaddedDeck) != 52 {
		t.Errorf("O deck deve ter 52 cartas, porem ele possui %v", len(d))
	}

	os.Remove("_decktesting")
}
