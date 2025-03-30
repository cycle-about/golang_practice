package main

import (
	"fmt"
	// "strings"
	// "math"
)

// type Vertex struct {
// 	X int
// 	Y int
// }

// // struct literal
// var (
// 	v1 = Vertex{1, 2}  // has type Vertex
// 	// can assign only subset of struct's fields using Name: syntax
// 	v2 = Vertex{X: 1}  // Y:0 is implicit
// 	v3 = Vertex{}      // X:0 and Y:0
// 	// prefix & to a struct returns a pointer to struct value
// 	p  = &Vertex{1, 2} // has type *Vertex
// 	// p prints as: &{1 2}
// )

func Run_tour_3() {
	/*
	// pointers
	i, j := 42, 2701  // i and j are int variables, and have addresses

	// original comment -> my comment
	p := &i         // point to i -> p is pointer to the address of variable i
	fmt.Println(*p) // read i through the pointer -> read the value at address pointed at by p
	*p = 21         // set i through the pointer -> set value at address pointed at by p
	fmt.Println(i)  // see the new value of i -> read the value at the address of var i

	p = &j         // point to j -> p points to address of variable j
	*p = *p / 37   // divide j through the pointer -> set value at address pointed to by p to the value at that address / 37
	fmt.Println(j) // see the new value of j -> read the value at address of var j, which is equal to *p

	// additional test prints
	fmt.Println(*p) // same as j
	fmt.Println(&j) // 0xc00009a050

	// a struct is a collection of fields
	// does not need a toString() method to be printable, like java object
	fmt.Println(Vertex{1, 2}) // {1 2}

	// struct fields accessed using dot
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X) // 4
	fmt.Println(v)   // {4 2}

	// struct fields can be accessed through struct pointer
	// notation does not require explicit dereference, so p.X rather than (*p).X
	p := &v // p is a pointer to address of Vertex v
	p.X = 1e9 // assign field X of the Vertex at address pointed to by p
	fmt.Println(v)  // {1000000000 2}

	fmt.Println(v1, p, v2, v3)

	// arrays cannot be resized
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// primes is an array
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// s is a slice, made from subset of primes
	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	// slices like a and b do not store any data, describes part of an underlying array
	// changing element in a slice changes them in the array also
	// other slices that share the array will also get the changes
	a := names[0:2]  // john, paul
	b := names[1:3]  // paul, george
	fmt.Println(a, b)

	b[0] = "XXX" // element 0 of slice b is element 1 of array names (so paul)
	fmt.Println(a, b) // a: john, XXX   b: XXX, george
	fmt.Println(names) // john, XXX, george, ringo -> predicted all of these correctly

	// page 9. slice literal is like an array literal without length
	// eg array literal: 
	// al := [3]bool{true, true, false}
	// fmt.Printf("Type: %T, Value: %v\n", al, al) // [3]bool
	// eg make that same array, and a slice that references it: 
	// sl := []bool{true, true, false}
	// fmt.Printf("Type: %T, Value: %v\n", sl, sl) // []bool
	
	// q := []int{2, 3, 5, 7, 11, 13}  // q is a slice literal that reference that array of ints
	// fmt.Println(q)

	// r := []bool{true, false, true, true, false, true}
	// fmt.Println(r)

	// make a struct that contains an int i and a bool b
	// ? anonymous struct?
	// make a slice literal of those objects, containing items with those values for i and b
	// s := []struct {
	// 	i int
	// 	b bool
	// }{
	// 	{2, true},
	// 	{3, false},
	// 	{5, true},
	// 	{7, true},
	// 	{11, false},
	// 	{13, true},
	// }
	// fmt.Println(s)

	// // getting and printing items in slices q and r
	// for i := 0; i < len(q); i++ {
	// 	fmt.Println(q[i])
	// }

	// my own idea: try to read from q and r to make struct like s

	// error - []struct{i int; b bool} (type) is not an expression
	// stu := []struct {
	// 	i int
	// 	b bool
	// }

	// this is valid, can assign to empty, and still must be used
	// stu := []struct {
	// 	i int
	// 	b bool
	// } {}

	// try for loop in initializator
	// error
		// ./tour_3_types.go:144:3: syntax error: unexpected for, expected expression
		// ./tour_3_types.go:148:1: syntax error: non-declaration statement outside function body

	// stu := []struct {
	// 	i int
	// 	b bool
	// } {
	// 	for i := 0; i < len(q); i++ {
	// 		stu.put(q[i], r[i])
	// 	}
	// }

	// page 10. slice defaults
	// s := []int{2, 3, 5, 7, 11, 13}  // make a slice

	// s = s[1:4]  // redefine slice
	// fmt.Println(s)  // 3, 5, 7

	// s = s[:2]  // redefine slice of slice
	// fmt.Println(s)  // 3, 5

	// s = s[1:]  // redefine slice of slice of slice
	// fmt.Println(s)  // 5

	// page 11. slice length and capacity: length is ____, capacity is ____
	// length of a slice: number of elements it contains
	// capacity of a slice: number of elements in the underlying array, counting from the first element in the slice
	//     -> this definition seems odd, omitting earlier elements but not later ones
	s := []int{2, 3, 5, 7, 11, 13}
	// len=6 cap=6 [2 3 5 7 11 13]
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s) // []  - len=0 cap=6 []

	// Extend its length.
	s = s[:4]
	printSlice(s)  // 2, 3, 5, 7 - len=4 cap=6 [2 3 5 7]

	// Drop its first two values.
	s = s[2:]
	printSlice(s)  // 5, 7 - len=2 cap=4 [5 7]

	// Try changing one of the slice operations in the example program to extend it beyond its capacity and see what happens
	s = s[4:]
	printSlice(s)  // panic: runtime error: slice bounds out of range [4:2]

	// page 12. nil slice
	// var s []int 					// nil slice -> [] 0 0 nil!
	s := []int{} 					// NOT nil slice, only empty -> [] 0 0
	fmt.Println(s, len(s), cap(s))  
	if s == nil {
		fmt.Println("nil!")
	}

	// page 13. creating slice with make
	// make function creates dynamically-sized arrays: allocates a zeroed array, returns slice that refers to it
	a := make([]int, 5)  // from an array with size 5, capacity default so 5, assign 'a' to slice to it
	// prediction -> len=5 cap=5 [0 0 0 0 0]
	printSlice("a", a)  // correct prediction

	b := make([]int, 0, 5)  // from array with size 0, capacity 5, assign 'b' to slice to it
	// prediction -> len=0 cap=5 []
	printSlice("b", b)  // correct prediction

	c := b[:2]  
	// prediction -> len=0 cap=2 []
	// from slice b indices 0-2, assign 'c' as a new slice
	// slice b is an empty slice, but giving it indices makes its len those 2 indices, and does NOT change its capacity 5
	// cap does not change because it 
	printSlice("c", c)
	// actual: len=2 cap=5 [0 0]
	// DOES NOT change cap b[:2] because 'first element of slice' is same as the array, count there to array end

	d := c[2:5]  // 
	// prediction -> arg is indices 2,3,4 of slice c, out of bounds????
	// revised prediction after knowing answer to c: 
	//     c has len=2 cap=5, this makes a slice with indices 2-5, so initializes 3 empty indices does not change cap
	//     -> len=3 cap=5 [0 0 0]
	printSlice("d", d)
	// actual: d len=3 cap=3 [0 0 0]
	// DOES change cap eg b[2:] because first element in slice is farther back in array, count there to end

	// 'The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice' (page 11)

	// page 14. slices of slices
	// slices can contain any type, including other slices
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// predict: len 3, and each of those items is a slice of strings -> correct
	fmt.Println("len board: ", len(board))

	// The players take turns: these are all reassigned at once
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

		// predicted end result - one not in right place
		// x o x
		// - - x
		// - - o

		// right answer
		// x - x
		// o - x
		// - - o

	// updated board prints once, space separated
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// page 15. append to slice
	// if current underlying array is too small, makes a new larger array pointed to by the returned slice
	var s []int  // nil
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)  // [0]

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)  // [0 1]

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)  // [0 1 2 3 4]

	// page 16. range iterates over slice or map
	// returns the index number, and copy of element at that index
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// page 17. range ignore one return value
	// like in python, ignore one of the returned values by assigning to '_'
	// var names CAN start with '_'
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i  // bit shift once left means multiply by 2
	}
	// for _, value := range pow {
	// 	fmt.Printf("%d\n", value)
	// 	fmt.Printf("%d\n", _)  // cannot use _ as value or type
	// }
		for _i, value := range pow {
		fmt.Printf("%d\n", value)
		fmt.Printf("%d\n", _i)
	}

	// page 19
	var m map[string]Vertex
	m = make(map[string]Vertex)
	fmt.Println(m)  			 // map[]
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])  // {40.68433 -74.39967}

	// page 20. map literal, keys are required
	// var m = map[string]Vertex{
	// 	"Bell Labs": Vertex{
	// 		40.68433, -74.39967,
	// 	},
	// 	"Google": Vertex{
	// 		37.42202, -122.08408,
	// 	},
	// }
	// fmt.Println(m)  // map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]

	// page 21. can omit type name when populating type literal, if just type (eg here Vertex)
	// var m = map[string]Vertex{
	// 	"Bell Labs": {40.68433, -74.39967},
	// 	"Google":    {37.42202, -122.08408},
	// }
	// fmt.Println(m)  // map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]

	// page 22. adding to and removing from maps
	// if value not found, returns default of the type, and second return value is boolean for if found
	// m := make(map[string]int)
	// fmt.Println("map when empty: ", m)
	// fmt.Println("The value before key created:", m["Answer"])  // 0

	// m["Answer"] = 42
	// fmt.Println("The value:", m["Answer"])

	// v, ok := m["Answer"]
	// fmt.Println("The returned value:", v, "Was the value found:", ok)

	// m["Answer"] = 48
	// fmt.Println("The value redefined:", m["Answer"])

	// delete(m, "Answer")
	// fmt.Println("The value after deleted:", m["Answer"])

	// // test that the key is actually present, using second return value of the find
	// v2, ok2 := m["Answer"]
	// fmt.Println("The returned value:", v2, "Was the value found:", ok2)

	// page 23. maps exercise

	// page 24. functional values
	// functions can be assigned variable names, and passed as args
	// assign function to variable
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	// pass value that's a function as an arg
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
	*/

	// page 25. function closures
	/*
	Go functions may be closures. A closure is a function value that references variables from outside its body. 
	The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
	For example, the adder function returns a closure. Each closure is bound to its own sum variable. 
	// original output
		0 0
		1 -2
		3 -6
		6 -12
		10 -20
		15 -30
		21 -42
		28 -56
		36 -72
		45 -90

	// original calls
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	// revised version outputs
	initialized new sum variable
	initialized new sum variable

	i = 0
	pos -> arg: 0, result: 0
	neg -> arg: 0, result: 0

	i = 1
	pos -> arg: 1, result: 1
	neg -> arg: -2, result: -2

	i = 2
	pos -> arg: 2, result: 3
	neg -> arg: -4, result: -6

	i = 3
	pos -> arg: 3, result: 6
	neg -> arg: -6, result: -12

	i = 4
	pos -> arg: 4, result: 10
	neg -> arg: -8, result: -20

	*/

	// my revision to make progress more clear
	pos, neg := adder(), adder()
	for i := 0; i < 5; i++ {
		fmt.Printf("\ni = %d\n", i)
		fmt.Printf("pos -> ")
		pos(i)
		fmt.Printf("neg -> ")
		neg(-2*i)
	}
}

// returns a closure; each closure bound to its own sum variable
func adder() func(int) int {
	// initialize 'sum' once when the function variable is first created
	fmt.Println("initialized new sum variable")
	sum := 0  // each gets its own sum variable, looks like new instance made with each function call
	// call this function each time the function is called with a new arg value
	return func(x int) int {  // 
		fmt.Printf("arg: %d, ", x)
		sum += x
		fmt.Printf("result: %d\n", sum)
		return sum
	}
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

type Vertex struct {
	Lat, Long float64
}

/*
	Behavior with methods with same name
	1. Two methods with same name, second one has correct args for its call
		2 errors: 
			- too many arguments in call to printSlice, have (string, []int), want ([]int)
			- printSlice redeclared in this block, ./tour_3_types.go:234:6: other declaration of printSlice
	2. Two methods with same name, first one has correct args for its call
		1 error:  printSlice redeclared in this block, ./tour_3_types.go:236:6: other declaration of printSlice
	3. Two methods with same name and same args
		1 error: printSlice redeclared in this block, ./tour_3_types.go:237:6: other declaration of printSlice
*/

// test with duplicate method signature
// func printSlice(s string, x []int) {
// 	fmt.Println("hello")
// }

// same method name, different signature due to args
func printSlice(s []uint8) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSliceShort(s []uint8) {
	fmt.Println(s)
}

// same method name, different signature due to args
// func print2DSlice(s [][]uint8) {
// 	fmt.Println("2D slice")
// 	for _, value := range s {
// 		fmt.Printf("%d\n", value)
// 	}
// }

func print2DSlice(s [][]uint8) {
	for i, value := range s {
		fmt.Printf("row=%d ", i)
		printSlice(value)
	}
	fmt.Printf("\n")
}

func print2DSliceShort(s [][]uint8) {
	for _, value := range s {
		printSliceShort(value)
	}
	fmt.Printf("\n")
}

// func printSlice(s string, x []int) {
// 	fmt.Printf("%s len=%d cap=%d %v\n",
// 		s, len(x), cap(x), x)
// }


