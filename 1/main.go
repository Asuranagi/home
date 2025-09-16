package main

import (
	"fmt"
	"strings"
)

//	var m = map[string]float64{
//		"USD_EUR": 0.85,
//		"USD_RUB": 84.1,
//		"EUR_RUB": 98,
//		"EUR_USD": 1.18,
//		"RUB_USD": 0.012038,
//		"RUB_EUR": 0.010262,
//	}
var m = map[string]map[string]float64{
	"USD": {"EUR": 0.85, "RUB": 84.1},
	"EUR": {"USD": 1.18, "RUB": 98},
	"RUB": {"USD": 0.012038, "EUR": 0.010262},
}

func main() {
	for {
		fmt.Println("`===` Конвертер валют `===`")
		original := inputCurrency("исходную")
		amount := inputAmount()
		target := inputCurrency("целевую")
		result := convert(original, amount, target)
		fmt.Printf("Результат: %.2f %s\n", result, target)
		ascQ := questions()
		if !ascQ {
			break
		}
	}
}
func inputCurrency(currencyType string) string {
	validCurrencies := []string{"USD", "EUR", "RUB"}
	for {
		fmt.Printf("Введите %s валюту (USD, EUR, RUB): ", currencyType)
		var currency string
		fmt.Scan(&currency)
		currency = strings.ToUpper(currency)
		for _, valid := range validCurrencies {
			if currency == valid {
				return currency
			}
		}
		fmt.Println("Ошибка! Доступные валюты: USD, EUR, RUB")
	}
}
func inputAmount() float64 {
	for {
		fmt.Print("Введите сумму: ")
		var amount float64
		n, err := fmt.Scan(&amount)
		if n == 1 && err == nil && amount > 0 {
			return amount
		}
		fmt.Println("Ошибка! Сумма должна быть больше 0")
	}
}

//	func convert(original string, amount float64, target string) float64 {
//		if original == target {
//			return amount
//		}
//		key := original + "_" + target
//		rate, ok := m[key]
//		if !ok {
//			fmt.Println("Такой валютной пары нет в базе. Конвертация невозможна!")
//			return 0
//		}
//		return amount * rate
//	}
func convert(original string, amount float64, target string) float64 {
	if original == target {
		return amount
	}
	if targets, ok := m[original]; ok {
		if rate, ok := targets[target]; ok {
			return amount * rate
		}
	}
	fmt.Println("Нет курса для выбранной пары валют!")
	return 0
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
