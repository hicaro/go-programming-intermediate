package main

import (
	"fmt"
	"sync"
)

func exercise_7() {
	c := make(chan string)

	go func(c chan string) {
		var wg sync.WaitGroup

		wg.Add(10)

		for i := 0; i < 10; i++ {
			go func(ch chan<- string, index int) {
				for j := 0; j < 10; j++ {
					ch <- fmt.Sprintf("Position %d, %d", index, j)
				}
				wg.Done()
			}(c, i)
		}
		wg.Wait()
		close(c)
	}(c)

	for value := range c {
		fmt.Println(value)
	}
}
