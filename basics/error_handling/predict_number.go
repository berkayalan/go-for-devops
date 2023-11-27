package error_handling

import (
	"errors"
	"fmt"
)

func Predictor(prediction int) (string, error) {

	correct_number := 55

	if prediction < 1 || prediction > 1000 {
		return "", errors.New("Please enter a number between 1 and 1000")
	}

	if prediction > correct_number {
		return "Please enter a smaller prediction", nil
	}

	if prediction < correct_number {
		return "Please enter a bigger prediction", nil
	}

	return "You predicted correctly!", nil

}

func PredictNumber() {
	status, err := Predictor(100)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(status)
	}
}
