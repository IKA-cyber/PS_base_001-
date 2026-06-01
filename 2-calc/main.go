package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// parseNumbers принимает строку вида "2, 10, 9" и возвращает срез чисел []float64.
func parseNumbers(input string) ([]float64, error) {
	parts := strings.Split(input, ",")

	numbers := make([]float64, 0, len(parts))

	for _, part := range parts {
		//Убираем пробелы вокруг числа.
		part = strings.TrimSpace(part)

		// Преобразуем строку в число.
		num, err := strconv.ParseFloat(part, 64)
		if err != nil {
			return nil, err
		}

		numbers = append(numbers, num)
	}
	return numbers, nil
}

// calculateSum вычисляет сумму всех чисел.
func calculateSum(numbers []float64) float64 {
	var sum float64

	for _, num := range numbers {
		sum += num
	}
	return sum
}

// calculateAvg вычисляет среднее арифметическое.
func calculateAvg(numbers []float64) float64 {
	sum := calculateSum(numbers)
	return sum / float64(len(numbers))
}

// calculateMedian вычисляет медиану.
func calculateMedian(numbers []float64) float64 {
	// Создаем копию, чтобы не изменять исходный срез.
	sorted := make([]float64, len(numbers))
	copy(sorted, numbers)

	// Сортируем числа по возрастанию.
	sort.Float64s(sorted)

	n := len(sorted)
	middle := n / 2

	// Если количество элементов нечетное
	if n%2 != 0 {
		return sorted[middle]
	}
	// Если количество элементов четное
	return (sorted[middle-1] + sorted[middle]) / 2
}

func main() {
	var operation string
	var input string
	// 1. Считываем операцию.
	fmt.Print("Введите операцию (SUM, AVG, MED): ")
	fmt.Scanln(&operation)

	// 2. Считываем строку с числами.
	fmt.Print("Введите числа через запятую: ")
	fmt.Scanln(&input)

	// 3. Разбираем строку в срез чисел.
	numbers, err := parseNumbers(input)
	if err != nil {
		fmt.Println("Ошибка при разборе чисел:", err)
		return
	}

	// Проверка, что хотя бы одно число введено.
	if len(numbers) == 0 {
		fmt.Println("Список чисел пуст")
		return
	}

	// 4. Выполняем нужную операцию
	switch operation {
	case "SUM":
		fmt.Printf("Сумма: %.2f\n", calculateSum(numbers))
	case "AVG":
		fmt.Printf("Среднее: %.2f\n", calculateAvg(numbers))
	case "MED":
		fmt.Printf("Медиана: %.2f\n", calculateMedian(numbers))
	default:
		fmt.Println("Неизвестная операция")
	}
}
