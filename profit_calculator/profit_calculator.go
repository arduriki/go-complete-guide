package main

import (
	"errors"
	"fmt"
	"os"
)

const resultsFile = "results.txt"

func main() {
	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		fmt.Printf("[ERROR] %v\n", err)
		fmt.Println("-----")
		return
	}
	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		fmt.Printf("[ERROR] %v\n", err)
		fmt.Println("-----")
		return
	}
	taxRate, err := getUserInput("Tax rate: ")
	if err != nil {
		fmt.Printf("[ERROR] %v\n", err)
		fmt.Println("-----")
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	showResults(ebt, profit, ratio)
}

func getUserInput(text string) (float64, error) {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)
	if userInput < 1 {
		return 0, errors.New("sorry, can't be equal or less than 0")
	}
	return userInput, nil
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f", ebt, profit, ratio)
	os.WriteFile(resultsFile, []byte(results), 0644)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}

func showResults(ebt, profit, ratio float64) {
	fmt.Printf("Earnings before taxes: %.2f\n", ebt)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)
	storeResults(ebt, profit, ratio)
}
