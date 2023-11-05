package main

import "fmt"

func main() {
	PrimeNumberCheck()
}

func PrimeNumberCheck() {

	var input_number int

	fmt.Println("Please enter a number")

	fmt.Scanln(&input_number)

	is_prime := true

	for num := 2; num < input_number; num++ {

		if input_number%num == 0 { // This is how we get mode of a transaction
			is_prime = false
		}
	}

	if is_prime == true {
		fmt.Println("This number is prime!")
	} else {
		fmt.Println("This number is not a prime!")
	}

}
