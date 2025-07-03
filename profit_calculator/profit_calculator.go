package main

import "fmt"

func main() {
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax rate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	showResults(ebt, profit, ratio)
}

func getUserInput(text string) float64 {
	var data float64
	fmt.Print(text)
	fmt.Scan(&data)
	return data
}

func calculateFinancials(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}

func showResults(ebt, profit, ratio float64) {
	fmt.Printf("Earnings before taxes: %.2f\n", ebt)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)
}
