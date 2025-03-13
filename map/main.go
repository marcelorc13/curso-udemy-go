package main

import "fmt"

func main() {
	colors := map[string]string{
		"vermelho": "#ff000",
		"verde":    "#4vf745",
		"branco":   "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("%v: %v", color, hex)
		fmt.Println()
	}

}
