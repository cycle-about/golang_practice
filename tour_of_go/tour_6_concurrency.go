package main

import (
	"fmt"
	"time"
	"sync"
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


// orginal functions
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


// with print statements, comments, experiments
func fibonacci_6(c, quit chan int) {
	fmt.Println("in fibonacci_6()")
	x, y := 0, 1
	for {
		// select blocks until one of its cases can run, then executes it
		select {
		case c <- x: 			// send_1 x to channel 'c'
			fmt.Println("value of x is ", x)
			x, y = y, x+y
		case <-quit: 			// receive_2 from channel 'quit', without assigning value
			fmt.Println("\nquit")
			return
		}
	}
}

func Run_tour_6() {
	fmt.Println("in main()")
	c := make(chan int)      // unbuffered channel of ints
	quit := make(chan int) 	 // also unbuffered channel of ints
	// execute this funtion literal in a new goroutine
	// must have another goroutine ready to receive from unbuffered channel before main goroutine sends to it
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("\nloop", i)
			v := <-c			// receive from channel c, assign to v
			fmt.Printf("received %d in channel c\n", v)
		}
		quit <- 0 				// send to channel 'quit'
	}()
	// main goroutine can start using the channels, after other goroutine is ready as send/receive counterpart
	fibonacci_6(c, quit)  // this is on main thread, and prints right after "in main()"
}
*/

////////////////////////////////////

/*
6. Default Selection
default case in a select runs if no other case is ready

Exercise: use a default case to try a send or receive without blocking
	select {
	case i := <-c:
	    // use i
	default:
	    // receiving from c would block
	}

Guess about what this will do: print '.' every 50 ms, and print tick every 100 ms (so two . between tick), then after 5 ticks with dots, print boom
It did that


// original sample code
func Run_tour_6() {
	tick := time.Tick(100 * time.Millisecond)
	fmt.Println(tick) 						// 0xc000134000
	fmt.Printf("tick is type %T\n", tick) 	// tick is type <-chan time.Time
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

// Exercise: use a default case to try a send or receive without blocking
// idea: try to receive from an empty buffer?

func fibonacci_6(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:   			// sends to channel
			fmt.Println("sent to channel c: ", x)
			x, y = y, x+y
		case <-quit: 			// never receives
			fmt.Println("quit")
			return
		// without default: deadlock error from 'goroutine 1 [select]'
		default:
			// result: infinite loop
			// fmt.Println("waiting")
			
			// result: "closing channels" panic: send on closed channel goroutine 1 [running]
			// fmt.Println("closing channels")
			// close(c)
			// close(quit)

			// result: does not do anything else, only "closing channel quit \n quit"
			// strange: is calling close on a channel like reaching that case? somehow it prints quit
			fmt.Println("closing channel quit")
			close(quit)
		}
	}
}

func Run_tour_6() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)  // receives 10 times from channel
		}
		// quit <- 0  // never send the condition that ends other goroutine
	}()
	// without the default that makes an infinite loop, both channels are closed when
	// the other goroutine in function literal completes, because then there is no send/receive counterpart
	fibonacci_6(c, quit)
}

// from chatgpt: non-blocking receive example
// this does NOT create an infinite loop because there is only one goroutine, which ends
// buffered channel can be used with only a single goroutine
// result without the default: all goroutines are asleep - deadlock! goroutine 1 [chan receive]:
func Run_tour_6() {
	c := make(chan int, 1) // buffered channel, capacity 1
    // No value sent yet, so c is empty
    select {
    case i := <-c:
        fmt.Println("Received:", i)
    // default:
    //     fmt.Println("No value to receive (channel is empty)")
    }
}

// my own non-blocking send example
// result: "no other cases can run"
func Run_tour_6() {
	c := make(chan int, 1) // buffered channel, capacity 1
	c <- 5  // goes into channel buffer, which is now full
    select {
    case c <- 3:
        fmt.Println("sent to channel")
    default:
        fmt.Println("no other cases can run")
    }
}
*/

////////////////////////////////////

/*
7. sync.Mutex
Ensures only one goroutine can access a variable at a time
Block of code can be defined to execute in mutual exclusion by surrouding with call to Lock and Unlock
defer can also ensure mutex will be unlocked
*/

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func Run_tour_6() {
	fmt.Println("in tour 6")
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}