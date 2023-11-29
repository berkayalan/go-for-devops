package main

import (
	"basics/variables"
	//"basics/conditionals"
	//switchstatement "basics/switch_statement"
	//"basics/loops"
	//"basics/arrays"
	//"basics/slices"
	//"basics/functions"
	//"basics/maps"
	//"basics/structs"
	//"basics/pointers"
	//"basics/goroutines"
	//"basics/interface_module"
	//"basics/defers"
	//"basics/error_handling"
	"basics/strings"
	"fmt"
)

func main() {

	// We can call another functions like that

	//VariablesFunction()

	//conditionals.CheckCondition()

	//switchstatement.CreateSwitch(15)

	//loops.YearCalculator()

	//arrays.CreateArray()

	//arrays.CreateMultidimensionalArray()

	//slices.CreateSlice()

	//slices.CreateSliceFromArray()

	//functions.CreateFunction()

	//name, birthyear, mail, country := functions.CallMultireturn()
	//fmt.Println(name, birthyear, mail, country)

	//functions.CreateVariadic("John", "Mark", "Berkay")

	//names := []string{"Kelvin", "Dimitar", "Berkay", "Alex", "Mauro"}
	//functions.CreateVariadic(names...) // 3 dots are important here.

	//maps.CreateMap()

	//maps.IterateMaps()

	//structs.CreateStruct()

	//num := 169
	//pointers.PointerExtractor(num)

	//go goroutines.Greetings("Berkay") // write go before the function to invoke it as a goroutine.
	//go goroutines.Greetings("Michel")
	// Wait for goroutines to finish before main goroutine ends
	//time.Sleep(time.Second)

	//interface_module.UserInterface()

	//result := defers.AgeCheck(2009)
	//fmt.Println(result)

	//error_handling.OpenFile()
	//error_handling.GetNumber()
	//error_handling.PredictNumber()
	//fmt.Println(error_handling.EnterCheck(2009, 12))

	//strings.CreateString()
	strings.StringAdvanced()

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
