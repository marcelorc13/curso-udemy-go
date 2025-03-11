package main

import "fmt"

type person struct {
	primeiroNome string
	ultimoNome   string
}

func main() {
	marcelo := person{primeiroNome: "Marcelo", ultimoNome: "Carnaúba"}
	fmt.Println(marcelo)
}
