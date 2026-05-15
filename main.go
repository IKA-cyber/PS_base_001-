package main

import (
	"fmt"
)

const (
	USDToEUR = 0.9
	USDToRUB = 80.0
)

func readAmount() float64 {
	var amount float64
	fmt.Print("Введите сумму: ")
	fmt.Scan(&amount)
	return amount
}

// convertCurrency — заготовка функции для конвертации.
// amount      - сумма для перевода
// fromCurrency - исходная валюта (например "USD")
// toCurrency   - целевая валюта (например "EUR")
//
// Пока функция ничего не вычисляет и просто возвращает 0.
func convertCurrency(amount float64, fromCurrency string, toCurrency string) float64 {
	return 0
}

func main() {
	amount := readAmount()
	result := convertCurrency(amount, "USD", "EUR")
	fmt.Printf("Результат конвертации: %.2f\n", result)
}
