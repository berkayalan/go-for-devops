package bufferedchannels

import (
	"fmt"
	"time"
)

/*

The select statement lets a goroutine wait on multiple communication operations.
A select blocks until one of its cases can run, then it executes that case.
It chooses one at random if multiple are ready.

*/

func SelectStatement() {

	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 5)
		chan1 <- "Berkay"
	}()

	go func() {
		time.Sleep(time.Second * 10)
		chan2 <- "Alan"
	}()

	var data1 string
	var data2 string

	for len(data1) == 0 || len(data2) == 0 {
		select {
		case data1 := <-chan1:
			fmt.Println("Data taken from channel 1: ", data1)
		case data2 := <-chan2:
			fmt.Println("Data taken from channel 2: ", data2)
		default:
			fmt.Println("No data yet")
		}
		time.Sleep(time.Second * 1)
	}

}
