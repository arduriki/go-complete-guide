package main

import "fmt"

func main() {
	age := 32 // regular variable

	var agePointer *int // declare a pointer
	agePointer = &age // initialize it

	fmt.Println("Age:", agePointer) // prints the memory address
	fmt.Println("Age:", *agePointer) // dereference the pointer (get value)
	
	// adultYears := getAdultYears(agePointer) // or (&age)
	editAgeToAdultYears(agePointer)
	// fmt.Println(adultYears)

	// Prints the changed original value
	fmt.Println(age)
}

// dereference in the function
func editAgeToAdultYears(age *int) {
	// return *age - 18

	// put a new value to the variable
	*age = *age - 18
}