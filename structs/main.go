package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	primeiroNome string
	ultimoNome   string
	contactInfo
}

func main() {
	jim := person{
		primeiroNome: "Jim",
		ultimoNome:   "Halpert",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipcode: 94000,
		},
	}

	jim.updateName("Jimer")
	jim.print()
}

func (pointerToPerson *person) updateName(novoPrimeiroNome string) {
	(*pointerToPerson).primeiroNome = novoPrimeiroNome
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
