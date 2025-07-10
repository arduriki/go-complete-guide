package lists

import "fmt"

func main() {
	// dynamic slices
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[1])
	prices[1] = 9.99
	prices = append(prices, 5.99)
	prices = prices[1:]

	fmt.Println(prices)

	discountPrices := []float64{10.99, 8.99, 20.59}
	// Merge slices
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}

// func main() {
// 	var productNames [4]string = [4]string{"A Book"}
// 	// store 4 float data into a list = they describe the same thing
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
// 	// Print the prices list.
// 	fmt.Println(prices)
// 	// Assign a certain value to a list's item.
// 	productNames[2] = "A Carpet"
// 	// Print the productNames list with zero values in them.
// 	fmt.Println(productNames)
// 	// Print a specific item in the list of prices.
// 	fmt.Println(prices[0])

// 	// From position 1 up to 2 - 3 is EXCLUDED
// 	// featuredPrices := prices[1:4]
// 	// From the begining to position 2
// 	// featuredPrices := prices[:3]
// 	// From position 1 until the end of the slice
// 	featuredPrices := prices[1:]
// 	featuredPrices[0] = 199.99
// 	highlightedPrices := featuredPrices[:1]
// 	fmt.Println(highlightedPrices) // 9.99
// 	// Arrays aren't copied, its edited
// 	fmt.Println(prices) // [1] item = 199.99
// 	// cap: is the capacity that a slice can get
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

// 	// "Reslized" the slice to the right, not to the left.
// 	highlightedPrices = highlightedPrices[:3]
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
// }

