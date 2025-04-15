package main

import (
	"fmt"
)

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

// List represents a singly-linked list that holds values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) getNext() *List[T] {
	return l.next
}

func Run_tour_5() {
	/* 
	1. Type parameters
	
	func Index[T comparable](s []T, x T) int
	
	This declaration means that s is a slice of any type T that fulfills the built-in constraint comparable. 
	x is also a value of the same type.

	comparable is a useful constraint that makes it possible to use the == and != operators on values of the type. 
	In this example, we use it to compare a value to all slice elements until a match is found. This Index function 
	works for any type that supports comparison. 
	*/

	// Index works on a slice of ints (because they fulfill 'comparable' constraint)
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings (because they also fulfill 'comparable' constraint)
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))

	// my question: are there types that can't be made into Index, because not comparable?
	// slice of maps, and a map
	ms := make([]map[string]int, 3)  // create slice of 3 maps

	for i := range ms {  // initialize each map in slice
        ms[i] = make(map[string]int)
    }

    ms[0]["Answer 1"] = 42
    ms[0]["Question 1"] = 41 
    ms[0]["Thing 1"] = 40
    ms[1]["Answer 2"] = 43
    ms[1]["Question 2"] = 44
    ms[1]["Thing 2"] = 45
    ms[2]["Answer 3"] = 46

    // single map to be the 'comparison'
    m := make(map[string]int)
    m["Another Thing"] = 1

    fmt.Println(ms)
    fmt.Println(m)

    // put these into Index
    // error: map[string]int does not satisfy comparable
    // fmt.Println(Index(ms, m))

    // try with an uninitialized slice?
    var s []int
    fmt.Println(s)  // []
    // fmt.Println(Index(s, nil))  // error: cannot use nil as int value in argument to Index
    fmt.Println(Index(s, 0))  // -1

    // slice with multiple matches?
    os := make([]int, 5)
	for i := range os {
	    os[i] = 1
	}
	fmt.Println(Index(os, 1))  // 0, so first match

	/*
	2. Generic types
	In addition to generic functions, Go also supports generic types ("T any"). A type can be parameterized with a type parameter, 
	which could be useful for implementing generic data structures.

	This example demonstrates a simple type declaration for a singly-linked list holding any type of value.

	As an exercise, add some functionality to this list implementation. 

	// List represents a singly-linked list that holds values of any type
	type List[T any] struct {
		next *List[T]
		val  T
	}

	*/

	// var ints List[int]
	// fmt.Println("ints: ", ints)  // {<nil> 0}   'next' is nil, and val is 0

	// var sts List[string]
	// fmt.Println("strings: ", sts)  // {<nil> }

	// // var lsts List[List]  // cannot use generic type List[T any] without instantiation
	
	// var byts List[byte]
	// fmt.Println("byts: ", byts)  // {<nil> 0}

	node3 := &List[int]{val: 126}
	node2 := &List[int]{val: 125, next: node3}
	node1 := &List[int]{val: 123, next: node2}
	list  := &List[int]{val: 120, next: node1}

	for node := list; node != nil; node = node.getNext() {
		fmt.Println(node.val)
	}

}