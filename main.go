package main

import "fmt"

func main() {
	const usdToEur = 0.85
	const usdToRub = 84.1
	const eurToRub = usdToRub / usdToEur
	fmt.Printf("1 Usd = %.2f Eur\n", usdToEur)
	fmt.Printf("1 Usd = %.2f Rub\n", usdToRub)
	fmt.Printf("1 Eur = %.2f Rub\n", eurToRub)
}
