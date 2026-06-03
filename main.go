package main

import (
	"fmt"
	"strings"
)

var rates = map[string]float64{
	"USD": 1.0,
	"EUR": 0.9,
	"RUB": 80.0,
}

func readCurrency(prompt string) string {
	var cur string

	for {
		fmt.Println(prompt)
		fmt.Println("Доступные валюты: USD, EUR, RUB")
		fmt.Print("Введите валюту: ")

		fmt.Scan(&cur)
		cur = strings.ToUpper(cur)

		if _, ok := rates[cur]; ok {
			return cur
		}
		fmt.Println("Ошибка: некорректная валюта. Попробуйте снова.")
	}
}

func readAmount() float64 {
	var amount float64

	for {
		fmt.Print("Введите сумму: ")

		_, err := fmt.Scan(&amount)
		if err == nil && amount >= 0 {
			return amount
		}
		fmt.Println("Ошибка: введите неотрицательное число.")

		var discard string
		fmt.Scanln(&discard)
	}
}

func convert(amount float64, from, to string) float64 {
	inUSD := amount / rates[from]
	return inUSD * rates[to]
}

func main() {
	fmt.Println("=== Конвертер валют ===")

	from := readCurrency("Выберите исходную валюту")
	amount := readAmount()
	to := readCurrency("Выберите целевую валюту")

	if from == to {
		fmt.Printf("\nРезультат: %.2f %s\n", amount, to)
		return
	}

	result := convert(amount, from, to)

	fmt.Printf(
		"\n%.2f %s = %.2f %s\n",
		amount,
		from,
		result,
		to,
	)
}
