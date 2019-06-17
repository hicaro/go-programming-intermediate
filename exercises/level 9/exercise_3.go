package main

import (
	"fmt"
	"runtime"
	"sync"
)

func exercise_3() {
	const gs = 50

	counter := 0
	var wg sync.WaitGroup

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
