package main

import (
	"fmt"
	"os"
)

func main() {
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)
	printResult(ebt, profit, ratio)
	saveToFile(ebt, profit, ratio)

}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	validator(userInput)
	fmt.Println(userInput)
	return userInput
}

func printResult(ebt, profit, ratio float64) {

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func saveToFile(ebt, profit, ratio float64) {
	stringResult := fmt.Sprintf("Ebt: %.1f\nProfit: %.2f\nRatio: %.4f\n", ebt, profit, ratio)
	os.WriteFile("balance.txt", []byte(stringResult), 0644)

}

func validator(input float64) {
	if input < 0 {
		panic("cannot be negative")
	} else if input == 0 {
		panic("cannot be zero")
	}

}
