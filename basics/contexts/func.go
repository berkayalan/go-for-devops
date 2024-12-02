package contexts

import (
	"context"
	"fmt"
	"time"
)

/*

Context is a built-in package in the Go standard library that provides a powerful toolset
for managing concurrent operations. It enables the propagation of cancellation signals, deadlines,
and values across goroutines, ensuring that related operations can gracefully terminate when necessary.
With context, you can create a hierarchy of goroutines and pass important information down the chain.

*/

func CreateUser(UserName string) {
	ctx := context.Background()

	ctx = context.WithValue(ctx, "Username", UserName)

	StringUser(ctx)
}

func StringUser(ctx context.Context) {
	fmt.Printf("User created with username %s at %s", ctx.Value("Username"), time.Now().String())

}

func CreateUserActivity(UserName string) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	go StringUserActivity(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("Task timed out") // the operation is terminated prematurely, resulting in a timeout.
	}
}

func StringUserActivity(ctx context.Context) {

	time.Sleep(time.Second * 3)

	fmt.Printf("User made an activity at %s", time.Now().String())

}
