package timebasics

import (
	"fmt"
	"time"
)

func CallNow() {
	fmt.Println("Now Function: ", time.Now())
	fmt.Println("Now Function's converted to UTC: ", time.Now().UTC())
	fmt.Println("Now Function's weekday ", time.Now().Weekday())
	fmt.Println("Now Function's year: ", time.Now().Year())
	fmt.Println("Now Function's day: ", time.Now().Day())
	fmt.Println("Now Function's hour: ", time.Now().Hour())

}
