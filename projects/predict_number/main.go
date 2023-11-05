package main

import "fmt"

func main() {
	PredictNumber()
}

func PredictNumber() {

	essential_number := 113

	var prediction int

	fmt.Println("Please enter your prediction")

	fmt.Scanln(&prediction)

	for essential_number != prediction {
		if prediction < essential_number {
			fmt.Println("Please enter a bigger number")
			fmt.Scanln(&prediction)
		} else if prediction > essential_number {
			fmt.Println("Please enter a smaller number")
			fmt.Scanln(&prediction)
		}
	}

	fmt.Println("Congrats!!! You found the number!")

}
