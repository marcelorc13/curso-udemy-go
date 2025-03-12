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
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
