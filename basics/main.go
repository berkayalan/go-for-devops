package main

import (
	"basics/variables"
	// "basics/conditionals"
	//"basics/loops"
	//"basics/arrays"
	//"basics/slices"
	"basics/functions"
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

	//slices.CreateSliceFromArray()

	//functions.CreateFunction()

	//name, birthyear, mail, country := functions.CallMultireturn()
	//fmt.Println(name, birthyear, mail, country)

	functions.CreateVariadic("John", "Mark", "Berkay")
	functions.CreateVariadic("Kelvin", "Dimitar", "Berkay", "Alex", "Mauro")

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
