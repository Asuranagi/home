package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	for {
		fmt.Println("Калькулятор")
		userIn := inputUserStr()
		inputUsNu := inputUserNumber()
		result := calculate(userIn, inputUsNu)
		fmt.Println(result)
	}
}
func inputUserStr() string {
	inputUs := []string{"AVG", "SUM", "MED"}
	for {
		var term string
		fmt.Print("Введите операцию которую хотите обработать: ")
		fmt.Scan(&term)
		for _, valid := range inputUs {
			if term == valid {
				return term
			}
		}
		fmt.Println("Введите правильно операцию")
	}
}
func inputUserNumber() []int {
	inputUsNu := []int{}
	for {
		var number string
		fmt.Print("Введите числа которую хотите обработать: ")
		fmt.Scan(&number)
		if number == "end" {
			break
		}
		num, err := strconv.Atoi(number)
		if err != nil {
			fmt.Println("Введите правильно число")
			continue
		}
		inputUsNu = append(inputUsNu, num)
	}
	return inputUsNu
}
func calculate(operation string, number []int) float64 {
	for {
		switch operation {
		case "SUM":
			sum := 0
			for _, n := range number {
				sum += n
			}
			return float64(sum)
		case "AVG":
			if len(number) == 0 {
				return 0
			}
			sum := 0
			for _, n := range number {
				sum += n
			}
			return float64(sum) / float64(len(number))
		case "MED":
			if len(number) == 0 {
				return 0
			}
			nums := make([]int, len(number))
			copy(nums, number)
			sort.Ints(nums)
			mid := len(nums) / 2
			if len(nums)%2 == 0 {
				return float64(nums[mid-1]+nums[mid]) / 2.0
			} else {
				return float64(nums[mid])
			}
		default:
			fmt.Println("Неизвестная операция")
			return 0
		}
	}
}
