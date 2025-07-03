// Ask fo revenue, expenses & tax rate
// Calculate earnings before taxes (EBT) and earnings after tax (profit)
// Calculate ratio (EBT / profit)
// Output EBT, profit and the ratio
package main

import "fmt"

func main() {
	revenue := outputText("Revenue: ")
	expenses := outputText("Expenses: ")
	taxRate := outputText("Tax rate: ")
	
	outputText("Expenses: ")
	fmt.Scan(&expenses)

	outputText("Tax rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Printf("Earnings before taxes: %v\n", ebt)
	fmt.Printf("Profit: %v\n", profit)
	fmt.Printf("Ratio: %v\n", ratio)
}

func outputText(text string) (data float64) {
	fmt.Print(text)
	fmt.Scan(data)
	return
}