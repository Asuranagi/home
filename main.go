package main

import (
	"errors"
	"fmt"
)

const usdToEur = 0.85
const usdToRub = 84.1
const eurToRub = usdToRub / usdToEur
const eurToUsd = 1.18
const rubToUsd = 0.012038
const rubToEur = 0.010262

func main() {

	// fmt.Printf("1 Usd = %.2f Eur\n", usdToEur)
	// fmt.Printf("1 Usd = %.2f Rub\n", usdToRub)
	// fmt.Printf("1 Eur = %.2f Rub\n", eurToRub)

	for {
		fmt.Println("Конвертация")
		original, number, target := userMessage()
		a, err := calculation(original, number, target)
		if err != nil {
			panic("Error")
		}
		fmt.Println(a)
		result := convert(original, number, target)
		fmt.Println(result)
		checkUserQuestion := questions()
		if !checkUserQuestion {
			break
		}
	}
}
func userMessage() (string, float64, string) {
	var original string
	var number float64
	var target string
	fmt.Print("Введите исходную валюту(USD, EUR, RUB): ")
	fmt.Scan(&original)
	fmt.Print("Введите число: ")
	fmt.Scan(&number)
	fmt.Print("Введите цулевую валюту: ")
	fmt.Scan(&target)
	return original, number, target
}

func calculation(original string, number float64, target string) (string, error) {
	var a string = "Успешно"
	if original != "USD" && original != "EUR" && original != "RUB" {
		return "Неподходящая валюта", errors.New("Введите правильные значения")
	}
	if number <= 0 {
		return "Неподходящее значение", errors.New("Введите правильные значения")

	}
	if target != "USD" && original != "EUR" && original != "RUB" && target == original {
		return "Неподходящая валюта или одинаковое значение с исходной валютой", errors.New("Введите правильные значения")
	}
	return a, nil
}
func convert(original string, number float64, target string) float64 {
	switch original + "_" + target {
	case "USD_EUR":
		return number * usdToEur
	case "EUR_USD":
		return number * eurToUsd
	case "USD_RUB":
		return number * usdToRub
	case "RUB_USD":
		return number / rubToUsd
	case "EUR_RUB":
		return number * eurToRub
	case "RUB_EUR":
		return number / rubToEur
	default:
		return number
	}
}
func questions() bool {
	var question string
	fmt.Print("Желаете ли повторить программу?")
	fmt.Scan(&question)
	if question == "y" || question == "Y" {
		return true
	}
	return false
}
