package bufferedchannels

import (
	"fmt"
	"time"
)

/*

Channels are a type-safe` message that helps to communicate between goroutines.
A channel connects two goroutines and synchronizes the exchange of any information that passes through it.
Each channel value has a capacity, which will be explained in the section after next. A channel value with a
zero capacity is called unbuffered channel and a channel value with a non-zero capacity is called buffered channel.


*/

func BufferedChannel() {

	buffered := make(chan int, 4)

	go func() {
		for i := 0; i <= 10; i++ {
			buffered <- i
			fmt.Println("Sent data: ", i)
			time.Sleep(time.Second * 1)
		}
		close(buffered) // We will get deadlock error if we don't close this.
	}()

	for data := range buffered {
		fmt.Println("Received Data: ", data)
		time.Sleep(time.Second * 5)
	}

}

func ReadOnlyChannel() {

	mychan := make(chan string, 3)

	mychan <- "TEST"
	OnlyRead(mychan)

}

func WriteOnlyChannel() {

	mychan := make(chan string, 3)

	OnlyWrite(mychan)
	val := <-mychan
	fmt.Println(val)

}

func OnlyRead(mychan <-chan string) {
	val := <-mychan
	fmt.Println(val)
	//mychan <- "example"  -- cannot send to receive-only channel mychan
}

func OnlyWrite(mychan chan<- string) {
	mychan <- "example"
	//val := <-mychan -- cannot send to sent-only channel mychan
}
