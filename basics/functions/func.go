package functions

import "fmt"

func CreateFunction() {

	/*
		A function is a block of statements that can be used repeatedly in a program.
		A function will not execute automatically when a page loads.
		A function will be executed by a call to the function.

	*/

	fmt.Println("This is a main function")

	SubFunctions()

}

func SubFunctions() {

	fmt.Println("This is a sub function")

}
