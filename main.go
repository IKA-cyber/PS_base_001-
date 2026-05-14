package main

import (
	"fmt"
)

const (
	USDToEUR = 0.9
	USDToRUB = 80.0
)

func main() {
	EURToRUB := USDToRUB / USDToEUR

	fmt.Printf("1 USD = %.2f EUR\n", USDToEUR)
	fmt.Printf("1 USD = %.2f RUB\n", USDToRUB)
	fmt.Printf("1 EUR = %.2f RUB\n", EURToRUB)
}
