package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func exercise_5() {
	const gs = 50

	var counter int64
	var wg sync.WaitGroup

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
