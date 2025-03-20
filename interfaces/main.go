package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type portugueseBot struct{}

func main() {
	eb := englishBot{}
	pb := portugueseBot{}

	printGreeting(eb)
	printGreeting(pb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// logica costumizada para gerar saudações em ingles
	return "Hi there!"
}

func (portugueseBot) getGreeting() string {
	return "Olá!"
}
