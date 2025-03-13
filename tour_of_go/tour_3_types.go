package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

// struct literal
var (
	v1 = Vertex{1, 2}  // has type Vertex
	// can assign only subset of struct's fields using Name: syntax
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	// prefix & to a struct returns a pointer to struct value
	p  = &Vertex{1, 2} // has type *Vertex
	// p prints as: &{1 2}
)

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
	*/

	// slice literal is like an array literal without length
	// eg array literal: [3]bool{true, true, false}
	// eg make that same array, and a slice that references it: []bool{true, true, false}
	q := []int{2, 3, 5, 7, 11, 13}  // q is a slice literal that reference that array of ints
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	// make a struct that contains an int i and a bool b
	// ? anonymous struct?
	// make a slice literal of those objects, containing items with those values for i and b
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	// my own idea: can q and r be used to make that slice literal of structs
}