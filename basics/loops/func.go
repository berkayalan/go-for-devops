package loops

import "fmt"

func YearCalculator() {

	// Let's calculate years for a young person to enter a pub.

	allowed_age := 18

	young_person := 11

	for young_person <= allowed_age {

		if allowed_age == young_person {
			fmt.Println("Yeeey! You can enter our pub anymore. Have fun buddy!")
		} else {
			fmt.Printf("Yo're %v years old, come after %v year(s)\n\nOne Year Later\n---------\n",
				young_person, allowed_age-young_person)
		}

		young_person += 1
	}
}
