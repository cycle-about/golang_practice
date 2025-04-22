package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for i, v := range s {
		fmt.Printf("index: %d, adding: %d\n", i, v)
		sum += v
	}
	c <- sum // send sum to channel c
}

func fibonacci_6(n int, c chan int) {
	x, y := 0, 1  				// initialize x and y
	for i := 0; i < n; i++ {
		c <- x  				// push current x into channel
		x, y = y, x+y 			// assign x to value of y, and y to value of x+y
	}
	close(c)  // call close on channel (no more values will be added) so range loop receiver's running on channel ends
}

func Run_tour_6() {
	/* 
	1. Goroutine: lightweight thread managed by Go runtime
	Evaluation of statement: in current goroutine
	Execution of statement: in new goroutine
	Run in same address space, so must synchronize access to shared memory
	Output of below: hello usually prints first, and last, order varies
	*/
	// go say("world")  // executed in new goroutine
	// say("hello")

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
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int) // make channel of type int
	go sum(s[:len(s)/2], c)  // goroutine computing sum of first half of slice
	fmt.Println("starting second slice")  // this prints before
	go sum(s[len(s)/2:], c)
	// sum(s[len(s)/2:], c)
	fmt.Println("about to assign variables to x and y")
	x, y := <-c, <-c // receive from c
	fmt.Println("assigned variables to x and y")

	fmt.Println("\nFinal Sums: ", x, y, x+y)

	/* what if call one as not a goroutine?
		index: 0, adding: -9
		index: 1, adding: 4
		index: 2, adding: 0
		index: 0, adding: 7
		index: 1, adding: 2
		index: 2, adding: 8
		fatal error: all goroutines are asleep - deadlock!
	
	QUESTION
		Why not able to complete printout of final sum?
		Does a channel have to be used in a goroutine?

	Experiments about this question
		1. 
	*/

	/* 
	3. Buffered channels
	Surprising, same error message when buffer is full as when it's empty
	*/

	// ch := make(chan int, 2)
	// // fmt.Println(<-ch)  // fatal error: all goroutines are asleep - deadlock!

	// ch <- 1
	// ch <- 2
	// // ch <- 3  // fatal error: all goroutines are asleep - deadlock!
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)

	/* 
	4. Range and Close
	Sender can close channel to indicate no more values will be sent
	Receivers can check if channel closed with second, boolean value in receive expression
	Channels only need to be closed to communicate to receiver that no more values are coming
		eg if the receiver is running a range loop on the channel

	cap: The cap built-in function returns the capacity of v, according to its type
		For channel: the channel buffer capacity, in units of elements; if v is nil, cap(v) is zero
	*/

	// c := make(chan int, 10)  // make a channel of ints with buffer capacity 6
	// go fibonacci_6(cap(c), c)  // pass the channel buffer size and channel
	// for i := range c {
	// 	fmt.Println(i)
	// }
}