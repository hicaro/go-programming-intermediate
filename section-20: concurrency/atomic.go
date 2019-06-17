package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func atomic_func() {
	const gs = 100

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	var counter int64

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			// yields the processor, allowing other goroutines to run. It does not suspend the current goroutine, so execution resumes automatically.
			runtime.Gosched()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)
}
