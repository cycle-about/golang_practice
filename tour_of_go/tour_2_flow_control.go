package main

import (
	"fmt"
	"math"
	// "runtime"
	// "time"
)


// if statements do not need parens
func sqrt(x float64) string {
	if x < 0 {
		fmt.Println("handling negative number:", x)
		return sqrt(-x) + "i"
	}
	fmt.Println("handling positive number:", x)
	return fmt.Sprint(math.Sqrt(x))
}


// if can also have a statement before the condition, only exists in the if/else
// note: both print statements IN method print BEFORE the return values
func pow_1(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		fmt.Println("Result is within limit")
		return v
	}
	// v no longer exists here
	fmt.Println("Result is larger than limit")
	return lim
}


// can use vars in statement before if in else alse
func pow_2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		fmt.Println("true clause, returning value")
		return v
	} else {
		fmt.Println("false clause, printing statement")
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}


func Sqrt(x float64) float64 {
	// z := 1.0  // changes only slightly after around 7 loops
	// z := x/2  // also converges quickly after around 7 loops
	z := x
	for i := 0; i < 10; i++ {
		fmt.Println(i, z)
		z -= (z*z - x) / (2*z)
	}
	return z
}


func Run_tour_2() {
	/*
	// with init, condition, and post
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)  // 45
	
	// without init or post
	sum = 1
	for ; sum < 1000; {
		fmt.Print(sum, " ")
		sum += sum
	}
	fmt.Println(sum)  // 1 2 4 8 16 32 64 128 256 512 1024

	// without init, post, or semicolons: same as 'while'
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// omit condition for infinite loop
	for {
	}

	fmt.Println(sqrt(2), sqrt(-4))
		// handling positive number: 2
		// handling negative number: -4
		// handling positive number: 4
		// 1.4142135623730951 2i

	fmt.Println(
		pow_1(3, 2, 10),
		pow_1(3, 3, 20),
	)
		// Result is within limit
		// Result is larger than limit
		// 9 20

		fmt.Println(
		pow_2(3, 2, 10),  // 3^2 = 9, so under 10, so true: return 9
		pow_2(3, 3, 20),  // 3^3 = 27, so over 20, so false: print statement and return 20
	)

	x := 8.0
	fmt.Println(Sqrt(x))
	fmt.Println("math.Sqrt:", math.Sqrt(x))
	

	// switch statement: cases don't need to be constants, no 'break'
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// switch without a condition, becomes like long if/else
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
	*/

	// defer: delays execution of function until surrounding function returns
	defer fmt.Println("world")
	fmt.Println("hello")

	// defers are pushed onto stack, then executed LIFO
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")

	// result with running BOTH defer snippets (on newlines)
	// printed immediately: hello counting done
	// stack push order: world 0 1 2 ... 8 9
	// -> hello counting done 9 8 ... 1 0 world
}
