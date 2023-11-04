package main

import (
	"basics/variables"
	"fmt"
)

func main() {

	// This is how we invoke a function in a package

	fmt.Println("\nString Variables")
	variables.StringCreator()

	fmt.Println("\nInteger Variables")
	variables.IntegerCreator()

	fmt.Println("\nFloat Variables")
	variables.FloatCreator()

	fmt.Println("\nBoolean Variables")
	variables.BooleanCreator()

}
