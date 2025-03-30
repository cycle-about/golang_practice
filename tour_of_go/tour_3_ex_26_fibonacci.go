package main

import "fmt"

/*
Implement a fibonacci function that returns a function (a closure) that returns
successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...). 

TEMPLATE

	func fibonacci() func() int {
	}

	func main() {
		f := fibonacci()
		for i := 0; i < 10; i++ {
			fmt.Println(f())
		}
	}
*/

// fibonacci is a function, that returns
// a function that returns an int.
func fibonacci() func() int {
	seed_0 := 0
	seed_1 := 2
	nums := make([]int, 12)
	next_sum := 0
	index := 0  // index in nums
	return func() int {
		if (index == 0) {
			next_sum = seed_0
		} else if (index == 1) {
			next_sum = nums[0] + seed_1
		} else {
			next_sum = nums[index] + nums[index-1]
		}
		index++
		nums[index] = next_sum
		return next_sum
	}
}

func Run_tour_3_ex_fibonacci() {
	fmt.Println("hello from Fibonacci")
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}