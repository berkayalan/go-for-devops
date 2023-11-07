package main

import (
	"basics/variables"
	// "basics/conditionals"
	//"basics/loops"
	//"basics/arrays"
	"basics/slices"
	"fmt"
)

func main() {

	// We can call another functions like that

	//VariablesFunction()

	//conditionals.CheckCondition()

	//loops.YearCalculator()

	//arrays.CreateArray()

	//arrays.CreateMultidimensionalArray()

	//slices.CreateSlice()

	slices.CreateSliceFromArray()

}

func VariablesFunction() {

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
