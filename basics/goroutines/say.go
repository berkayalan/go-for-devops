package goroutines

import (
	"fmt"
	"sync"
	"time"
)

func Printer(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		if s == "hello" {
			time.Sleep(100 * time.Millisecond)
			fmt.Println(s)
		} else {
			time.Sleep(500 * time.Millisecond) // It will wait until this ends.
			fmt.Println(s)
		}
	}
}

func Say() {
	var wg sync.WaitGroup
	wg.Add(2)
	go Printer("world", &wg)
	go Printer("hello", &wg)
	wg.Wait()
}
