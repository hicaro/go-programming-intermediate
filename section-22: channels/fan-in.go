package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func fan_in() {
	fan_in_pike()
}

func fan_in_pike() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring; I'm leaving.")
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

// FAN IN
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

/*
code source:
Rob Pike
https://talks.golang.org/2012/concurrency.slide#25

source:
https://blog.golang.org/pipelines
*/

func fan_in_todd() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go func(e, o chan<- int) {
		for i := 0; i < 10; i++ {
			if i%2 == 0 {
				e <- i
			} else {
				o <- i
			}
		}
		close(e)
		close(o)
	}(even, odd)

	go func(e, o <-chan int, f chan<- int) {
		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			for v := range e {
				f <- v
			}
			wg.Done()
		}()

		go func() {
			for v := range o {
				f <- v
			}
			wg.Done()
		}()

		wg.Wait()
		close(fanin)
	}(even, odd, fanin)

	for v := range fanin {
		fmt.Println(v)
	}
}
