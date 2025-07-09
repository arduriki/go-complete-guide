package main

import (
	"fmt"

	"example.com/practice-arrays-slices/product"
)

func main() {
	// 1) Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line.
	hobbies := [3]string{"Coding", "Listen to music", "Read books"}
	fmt.Println(hobbies)
	// 2) Also output more data about that array:
	//		- The first element (standalone)
	fmt.Println(hobbies[0])
	//		- The second and third element combined as a new list
	fmt.Println(hobbies[1:3])

	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)
	hobbiesSlice := []string{hobbies[0], hobbies[1]}

	hobbiesSliceUpdated := hobbies[:2]
	fmt.Println(hobbiesSliceUpdated)

	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	hobbiesSlice = hobbiesSlice[1:3]
	hobbiesSlice = append(hobbiesSlice, hobbies[2])
	fmt.Println(hobbiesSlice)

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	courseGoals := []string{"Learn Go", "Create APIs"}

	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	courseGoals[1] = "Develop different projects"
	courseGoals = append(courseGoals, "Build fast apps")

	fmt.Println(courseGoals)

	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.
	firstProduct, err := product.New("Tomatoe sauce", "123456", 6.99)
	if err != nil {
		fmt.Println(err)
		return
	}

	secondProduct, err := product.New("Pizza", "654321", 9.99)
	if err != nil {
		fmt.Println(err)
		return
	}

	productsList := []product.Product{*firstProduct, *secondProduct}

	fmt.Println(productsList)

	thirdProduct, err := product.New("Spaghettis", "321456", 3.99)
	if err != nil {
		fmt.Println(err)
		return
	}

	productsList = append(productsList, *thirdProduct)

	fmt.Println(productsList)
}
