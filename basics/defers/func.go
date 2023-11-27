package defers

import (
	"fmt"
	"time"
)

func AgeCheck(birthyear int) string {

	/*

		Defer statements delay the execution of the function or method or an anonymous method until
		the nearby functions returns. In other words, defer function or method call arguments evaluate
		instantly, but they don’t execute until the nearby functions returns.
		We can create a deferred method, or function, or anonymous function by using the defer keyword.

		Defer Functions works with Last In First Out order.

	*/

	defer WelcomeMessage()     // run in every situation
	defer GreetingsForTheDay() // run in every situation

	year, _, _ := time.Now().Date()

	if year-birthyear >= 18 {
		return "You're allowed to enter. Have Fun!"
	} else {
		return "Sorry, You're not allowed to enter."
	}

}

func WelcomeMessage() {
	fmt.Println("Welcome to our pub!")
}

func GreetingsForTheDay() {
	fmt.Println("Have a nice day!")
}
