package functionsarevalues

import "fmt"

// Custom types.
type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 1, 2}
	// In the arguments, just enter the function name that is going to be used
	// inside of the the other function.
	// If using func(var) it'll pass just the result of the function.
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	transformerFn1 := getTransformerFunction(&numbers)
	transformerFn2 := getTransformerFunction(&moreNumbers)

	transformedNumbers := transformNumbers(&numbers, transformerFn1)
	moreTransformedNumbers := transformNumbers(&moreNumbers, transformerFn2)

	fmt.Println(transformedNumbers)
	fmt.Println(moreTransformedNumbers)
}

// Insert a function firm in the parameters.
func transformNumbers(numbers *[]int, transform transformFn) []int {
	var dNumbers []int

	for _, val := range *numbers {
		// Using a utility function in another function.
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
	// Return the function, not the result of executing it.
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
