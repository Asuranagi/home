package main

import (
	"fmt"
)

const usdToEur = 0.85
const usdToRub = 84.1
const eurToRub = usdToRub / usdToEur
const eurToUsd = 1.18
const rubToUsd = 0.012038
const rubToEur = 0.010262

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
		fmt.Scan(&amount)
		if amount > 0 {
			return amount
		}
		fmt.Println("Ошибка! Сумма должна быть больше 0")
	}
}

// func calculation(original string, number float64, target string) (string, error) {
// 	var a string = "Успешно"
// 	if original != "USD" && original != "EUR" && original != "RUB" {
// 		return "Неподходящая валюта", errors.New("Введите правильные значения")
// 	}
// 	if number <= 0 {
// 		return "Неподходящее значение", errors.New("Введите правильные значения")

//		}
//		if target != "USD" && original != "EUR" && original != "RUB" && target == original {
//			return "Неподходящая валюта или одинаковое значение с исходной валютой", errors.New("Введите правильные значения")
//		}
//		return a, nil
//	}
func convert(original string, number float64, target string) float64 {
	switch original + "_" + target {
	case "USD_EUR":
		return number * usdToEur
	case "EUR_USD":
		return number * eurToUsd
	case "USD_RUB":
		return number * usdToRub
	case "RUB_USD":
		return number * rubToUsd
	case "EUR_RUB":
		return number * eurToRub
	case "RUB_EUR":
		return number * rubToEur
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
