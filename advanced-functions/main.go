package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	sum := sumup(1, 10, 15, 40, -5)
	// Pass values as standalone int values with ...
	anotherSum := sumup(1, numbers...)

	fmt.Println(sum)
	fmt.Println(anotherSum)
}

// sumup as a variadic function: any amount of parameters. Shown as ...
// Place the ... at the end of the function's parameters.
// It can accept another parameter that isn't "variadic".
func sumup(startingValue int, numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
