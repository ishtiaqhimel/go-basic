package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Context")
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	// ctx, cancel := context.WithCancel(ctx)
	// time.AfterFunc(time.Second, cancel)

	// go func() {
	// 	// s := bufio.NewScanner(os.Stdin)
	// 	// s.Scan()

	// 	time.Sleep(time.Second)
	// 	cancel()
	// }()

	sleepAndTalk(ctx, 5*time.Second, "Hello")
}

func sleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	select {
	case <-time.After(d):
		fmt.Println(msg)
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
