package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func fan_out() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go fanOutInWithThrottle(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}
}

func populate(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup

	// unlimite number of workers (as many as the number of values from c1 channel)
	for v := range c1 {
		wg.Add(1)
		// asynchrounous func to perform task
		go func(v2 int) { // value is passed as parameter to avoid closure
			c2 <- timeConsumingWork(v2)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(c2)
}

func fanOutInWithThrottle(c1, c2 chan int) {
	var wg sync.WaitGroup
	const workers = 10
	wg.Add(workers)

	// create 10 workers
	for i := 0; i < workers; i++ {
		// asynchrounous func to launch task
		go func() {
			// each worker ranges on c1 channel concurrently
			for v := range c1 {
				// synchronous func to perform task
				func(v2 int) { // value is passed as parameter to avoid closure
					c2 <- timeConsumingWork(v2)
				}(v)
			}

			wg.Done()
		}()
	}

	wg.Wait()
	close(c2)
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
