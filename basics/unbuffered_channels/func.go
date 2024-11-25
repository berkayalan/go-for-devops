package unbufferedchannels

import (
	"fmt"
	"time"
)

/*

Channel mainly acts as a concurrency synchronization technique. It helps to manage data between
different goroutines. We can view a channel as an internal FIFO (first in, first out) queue within a program.

A channel value with a zero capacity is called unbuffered channel and a channel value with a non-zero capacity
is called buffered channel. Unbuffered channels don't have a buffer to store data. This means that when a goroutine
sends data to an unbuffered channel, the data is not stored in a buffer, but is instead immediately passed to
the goroutine that is trying to receive data from the channel. Similarly, when a goroutine receives data from
an unbuffered channel, it blocks until data is available to be received.
*/

func UnbufferedChannel() {
	mychannel := make(chan string)

	done := make(chan bool) // This is like wait group, it will check if goroutines finished.

	go func() {
		message := "This is a message from the first goroutine."
		mychannel <- message
	}()

	go func() {
		message := <-mychannel
		fmt.Println("This message is taken from the first goroutine: ", message)
		// NOTE: If we don't read message from channel, it throws an error.
		done <- true
	}()

	<-done // This will check if goroutines ended.

	fmt.Println("End of the function.")

}

func UnbufferedChannelMultiple() {
	mychan := make(chan int)

	go func() {
		for num := 1; num <= 5; num++ {
			mychan <- num
			fmt.Println("Sent data: ", num)
		}
		close(mychan) // It helps to close channel.
	}()

	for {
		data, isOpen := <-mychan
		if isOpen == false {
			break
		}
		fmt.Println("Received data: ", data)
		time.Sleep(1 * time.Second)
	}

}
