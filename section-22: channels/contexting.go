package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func contexting() {
	contexting_example()
}

func contexting_example() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("Error check 1:", ctx.Err())
	fmt.Println("# goroutines 1:", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("Working", n)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("Error check 2:", ctx.Err())
	fmt.Println("# goroutines 2:", runtime.NumGoroutine())

	time.Sleep(time.Second * 2)

	fmt.Println("About to cancel context")
	cancel()
	fmt.Println("Canceled context")

	time.Sleep(time.Second * 2)
	fmt.Println("Error check 3:", ctx.Err())
	fmt.Println("# goroutines 3:", runtime.NumGoroutine())
}

func contexting_print() {
	ctx := context.Background()

	fmt.Println("Context:\t", ctx)
	fmt.Println("Context err:\t", ctx.Err())
	fmt.Printf("Context type:\t%T\n", ctx)

	fmt.Println("--------------------------")

	ctx, cancel := context.WithCancel(ctx)

	fmt.Println("Context:\t", ctx)
	fmt.Println("Context err:\t", ctx.Err())
	fmt.Printf("Context type:\t%T\n", ctx)
	fmt.Println("Context cancel:\t", cancel)
	fmt.Printf("Context cancel type:\t%T\n", cancel)

	fmt.Println("--------------------------")

	cancel()

	fmt.Println("Context:\t", ctx)
	fmt.Println("Context err:\t", ctx.Err())
	fmt.Printf("Context type:\t%T\n", ctx)
	fmt.Println("Context cancel:\t", cancel)
	fmt.Printf("Context cancel type:\t%T\n", cancel)

	fmt.Println("--------------------------")
}
