package main

import "fmt"

func main() {
	intSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, numero := range intSlice {
		if numero%2 == 0 {
			fmt.Println(numero, "é Par")
		} else {
			fmt.Println(numero, "é Ímpar")
		}
	}
}
