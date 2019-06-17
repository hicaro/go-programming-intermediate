package main

import (
	"fmt"
	"sync"
)

func exercise_1() {

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		fmt.Println("Gorouting #1")
		wg.Done()
	}()

	go func() {
		fmt.Println("Gorouting #2")
		wg.Done()
	}()

	wg.Wait()
}
