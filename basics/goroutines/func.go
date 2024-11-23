package goroutines

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

/*

A Goroutine is a function or method which executes independently and simultaneously in connection
with any other Goroutines present in our program. The cost of creating Goroutines is very small as
compared to the thread. Every program contains at least a single Goroutine and that
Goroutine is known as the main Goroutine. All the Goroutines are working under the main Goroutines
if the main Goroutine terminated, then all the goroutine present in the program also terminated.

If we do not add waitgroups, function can finish before TEST 2 and TEST 3 executed because
go scheduler only waits for default routine which is TEST 1.

Note: goroutines aren’t able to return values like a standard function would.

*/

func Test() {

	startTime := time.Now()

	MakeUppercase("TEST 1")

	MakeUppercase("TEST 2")

	MakeUppercase("TEST 3")

	fmt.Println("Executed Time without WaitGroup: ", time.Since(startTime))
}

func TestWaitGroup() {

	startTime := time.Now()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() { // running concurrently
		defer wg.Done()
		MakeUppercase("TEST 2 - Another routine")
	}()

	go func() { // running concurrently
		defer wg.Done()
		MakeUppercase("TEST 3 - Another routine")
	}()

	MakeUppercase("TEST 1")

	wg.Wait()

	fmt.Println("Executed Time with WaitGroup: ", time.Since(startTime))
}

func MakeUppercase(word string) {
	fmt.Println(strings.ToUpper(word))
}
