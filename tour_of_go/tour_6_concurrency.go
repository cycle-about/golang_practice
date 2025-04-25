package main

import (
	"fmt"
	// "time"
)

/* 
1. Goroutine: lightweight thread managed by Go runtime
Evaluation of statement: in current goroutine
Execution of statement: in new goroutine
Run in same address space, so must synchronize access to shared memory

Output: hello usually prints first, and last, order varies

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func Run_tour_6() {
	go say("world")  // executed in new goroutine
	say("hello")
}
*/

////////////////////////////////////

/* 
2. Channel: typed conduit to send and receive values with operator <-

ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and assign value to v

Data flows in direction of arrow
By default, sends and receives block until other side is ready
	For unbuffered channel, sender and receiver must be ready simultaneously
Default blocking allows goroutines to synchronize without explicit locks or condition variables
Must be created before use, like other variables (can use :=)
Channels are variables and can be function args
Example distributes work between two goroutines
Channel is modified by the function (no return value)

These goroutines print sequencially: making first half (x) finishes before second half (y) starts
That is unlike the example in 1 with only one goroutine that mixed with current thread

Behavior of print statements with goroutines and channels:
	1. prints statements on the main thread (up until receiving from the goroutines)
	2. prints all 3 items handled in first goroutine started (second half of list)
	3. prints all 3 items handled in second goroutine started (first half of list)
	4. then statements on main thread after the goroutines have send values to the channel

	about to start first goroutine
	about to start second goroutine
	about to assign variables to x and y
	index: 0, adding: -9
	index: 1, adding: 4
	index: 2, adding: 0
	index: 0, adding: 7
	index: 1, adding: 2
	index: 2, adding: 8
	assigned variables to x and y

	Final Sums:  -5 17 12

	Result if change line order of the goroutines: print statements now go 7 2 8 -9 4 0


func sum(s []int, c chan int) {
	sum := 0
	for i, v := range s {
		fmt.Printf("index: %d, adding: %d\n", i, v)
		sum += v
	}
	c <- sum // send sum to channel c
}

func Run_tour_6() {
	// s := []int{7, 2, 8, -9, 4, 0}
	// c := make(chan int) // make channel of type int

	// fmt.Println("about to start first goroutine")
	// go sum(s[:len(s)/2], c)  // goroutine computing sum of first half of slice
	// fmt.Println("about to start second goroutine")  // this prints before
	// go sum(s[len(s)/2:], c)
	// fmt.Println("about to assign variables to x and y")
	// x, y := <-c, <-c // receive from c, twice
	// fmt.Println("assigned variables to x and y")

	// without extra print statements
	// go sum(s[len(s)/2:], c)
	// go sum(s[:len(s)/2], c)  // goroutine computing sum of first half of slice
	// x, y := <-c, <-c // receive from c, twice

	// fmt.Println("\nFinal Sums: ", x, y, x+y)

	// try out getting the code snippet in description to work
	// an unbuffered channel must be used between multiple goroutines, or else deadlocks
		ch := make(chan string)
		go func() {  // launch in new goroutine
		    ch <- "hello"
		}()  // function literal: parentheses mean call it right away after defining
		value := <-ch
		fmt.Println(value)
	}
*/

////////////////////////////////////

/* 
3. Buffered channels
Sends to a buffered channel block when buffer is full
Receives block when buffer is empty
Same error message when buffered channel is at capacity as when it's empty, except for chan receive vs chan send
	

func Run_tour_6() {
	ch := make(chan int, 2)
	// no other goroutine is already running that could send a value
	// empty buffer: receive blocks
	// fmt.Println(<-ch)  // fatal error: all goroutines are asleep - deadlock! goroutine 1 [chan receive]

	// ch <- 4
	// ch <- 7
	// overfill buffer: send blocks
	// // ch <- 3  // fatal error: all goroutines are asleep - deadlock!
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)

	// try alternating sends and receives: same result as two sends then two receives
	ch <- 4
	fmt.Println(<-ch)
	ch <- 7
	fmt.Println(<-ch)
}
*/

////////////////////////////////////

/* 
4. Range and Close
Sender can close channel to indicate no more values will be sent
Receivers can check if channel closed with second, boolean value in receive expression
Channels only need to be closed to communicate to receiver that no more values are coming
	eg if the receiver is running a range loop on the channel

cap: The cap built-in function returns the capacity of v, according to its type
	For channel: the channel buffer capacity, in units of elements; if v is nil, cap(v) is zero

Result if remove close(c) by sender:
	"fatal error: all goroutines are asleep - deadlock! goroutine 1 [chan receive]"
	infer: the goroutine function has ended, no more goroutines to send anything, but receiver still waiting


// sender to channel
func fibonacci_6(n int, c chan int) {
	x, y := 0, 1  				// initialize x and y
	for i := 0; i < n; i++ {
		c <- x  				// push current x into channel
		x, y = y, x+y 			// assign x to value of y, and y to value of x+y
	}
	close(c)  // call close on channel (no more values will be added) so range loop receiver's running on channel ends
}

func Run_tour_6() {
	c := make(chan int, 10)  // make a channel of ints with buffer capacity 6
	go fibonacci_6(cap(c), c)  // pass the channel buffer size and channel
	for i := range c {  // range is the receiver, which receives until channel is closed
		fmt.Println(i)
	}

	// use the channel close check
	ch := make(chan int, 2)
	ch <- 4
	v0, ok0 := <-ch
	fmt.Println(v0, ok0)  // 4 true
	ch <- 7
	v1, ok1 := <-ch
	fmt.Println(v1, ok1)  // 7 true
	// v4, ok4 := <-ch
	// fmt.Println(v4, ok4)  // try to receive on an open, empty channel: receiver lock error
	close(ch)
	v2, ok2 := <-ch
	fmt.Println(v2, ok2)  // 0 false  // receiving an a closed channel returns 0
	v3, ok3 := <-ch
	fmt.Println(v3, ok3)  // 0 false
}
*/

////////////////////////////////////

/*
5. Select
Select statement lets goroutine wait on multiple communication operations
Blocks until one of its cases can run, then executes that case
If multiple are ready chooses one at random
*/

func fibonacci_6(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func Run_tour_6() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci_6(c, quit)
}