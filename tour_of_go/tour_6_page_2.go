package main

import (
	"fmt"
)

func sum_orig(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func sum_prints(s []int, c chan int) {
	sum := 0
	for i, v := range s {
		fmt.Printf("index: %d, adding: %d\n", i, v)
		sum += v
	}
	c <- sum // send sum to channel c
}

func Run_tour_6_page_2() {
	/* 
	2. Channel: typed conduit to send and receive values with operator <-
	Data flows in direction of arrow
	Must be created before use, like other variables (can use :=)
	Channels are variables and can be function args
	Example distributes work between two goroutines
	Channel is modified by the function (no return value)

	These goroutines run sequencially: first making x finishes before y starts
	That is unlike the example in 1 with only one goroutine that mixed with current thread
	*/

	// 1. start, working version
	/*
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum_prints(s[:len(s)/2], c)
	go sum_prints(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x+y)

	output 1.
		orig: -5 17 12
		prints:
			index: 0, adding: -9
			index: 1, adding: 4
			index: 2, adding: 0
			index: 0, adding: 7
			index: 1, adding: 2
			index: 2, adding: 8
			-5 17 12
	*/

	// 2. call second function not in goroutine
	/*
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum_prints(s[:len(s)/2], c)
	sum_prints(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x+y)

	output 2.
		index: 0, adding: -9
		index: 1, adding: 4
		index: 2, adding: 0
		index: 0, adding: 7
		index: 1, adding: 2
		index: 2, adding: 8
		fatal error: all goroutines are asleep - deadlock!
	*/

	// 3. call both functions not in goroutine
	/*
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	sum_prints(s[:len(s)/2], c)
	sum_prints(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

	output 3.
		index: 0, adding: 7
		index: 1, adding: 2
		index: 2, adding: 8
		fatal error: all goroutines are asleep - deadlock!
	*/

	// 4. call first function not in goroutine
	/* 
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	sum_prints(s[:len(s)/2], c)
	go sum_prints(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

	output 4. 
	index: 0, adding: 7
	index: 1, adding: 2
	index: 2, adding: 8
	fatal error: all goroutines are asleep - deadlock!
	*/

	// 5. first function not in goroutine, and only receive from c once
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	sum_prints(s[:len(s)/2], c)
	go sum_prints(s[len(s)/2:], c)
	x := <-c
	fmt.Println(x)

	/*
	output 5.
	index: 0, adding: 7
	index: 1, adding: 2
	index: 2, adding: 8
	fatal error: all goroutines are asleep - deadlock!
	*/
	
	// QUESTION
	// 	Why not able to complete printout of final sum?
	// 	Does a channel have to be used in a goroutine?

	// Experiments about this question
	// 	1. 
}