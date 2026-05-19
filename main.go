package main

import (
	"fmt"
)

const (
	USDToEUR = 0.9
	USDToRUB = 80.0
)

func readCurrency(promt string) string {
	var cur string

	for {
		fmt.Println(cur)
		fmt.Println("Доступные валюты: USD, EUR, RUB")
		fmt.Print("Введите валюту: ")

		fmt.Scan(&cur)

		switch cur {
		case "USD", "EUR", "RUB":
			return cur
		default:
			fmt.Println("Ошибка: некорректная валюта.Попробуйте снова")
		}
	}
}
func readAmount() float64 {
	var amount float64

	for {
		fmt.Print("Введите сумму: ")

		_, err := fmt.Scan(&amount)
		if err == nil {
			return amount
		}

		fmt.Println("Ошибка: введите число.")

		var discard string
		fmt.Scanln(&discard)
	}
}

func convert(amount float64, from string, to string) float64 {
	var inUSD float64

	// сначала в USD
	switch from {
	case "USD":
		inUSD = amount
	case "EUR":
		inUSD = amount / USDToEUR
	case "RUB":
		inUSD = amount / USDToRUB
	}

	// из USD в нужную валюту
	switch to {
	case "USD":
		return inUSD
	case "EUR":
		return inUSD * USDToEUR
	case "RUB":
		return inUSD * USDToRUB
	}

	return 0
}

func main() {
	fmt.Println("=== Конвертер валют ===")

	from := readCurrency("Выберите исходную валюту")
	amount := readAmount()
	to := readCurrency("Выберите целевую валюту")

	result := convert(amount, from, to)

	fmt.Printf("\nРезультат: %.2f %s\n", result, to)
}
