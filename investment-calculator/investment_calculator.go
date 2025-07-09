package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	// fmt.Print("Investment amount: ")
	outputText("Investment amount: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Expected return rate: ")
	outputText("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)
	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFutureValue := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRealFutureValue := fmt.Sprintf("Future Value (adjusted for Inflation): %.1f\n", futureRealValue)
// 	fmt.Printf(
// 		`Future value: %.1f
// Future Value (adjusted for Inflation): %.1f
// `, futureValue, futureRealValue)
	// fmt.Println("Future Value (adjusted for Inflation):", futureRealValue)
	fmt.Print(formattedFutureValue, formattedRealFutureValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (futureValue float64, realFutureValue float64) {
	// Because the variables are created in the return values it isn't
	// necessary to declare them again.
	futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	realFutureValue = futureValue / math.Pow(1+inflationRate/100, years)
	// return futureValue, realFutureValue
	return
}