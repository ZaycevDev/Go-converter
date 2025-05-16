package main

import "fmt"

const USDtoEUR = 0.89
const USDtoRUB = 80.85

func main() {
	EURtoRUB := USDtoRUB / USDtoEUR
	fmt.Printf("1 EUR = %.2f RUB\n", EURtoRUB)
}
