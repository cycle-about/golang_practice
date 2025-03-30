package main

import (
	"fmt"
	"math"
)


// struct with two fields, both type float64
type Vertex struct {
	X, Y float64
}

// receiver can be non-struct type (eg a float type)
type MyFloat float64
func (f MyFloat) AbsMethod() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// receiver: use to define a method on a type (since Go does not have classes)
// receiver here is type Vertex, named v
// receiver must be defined in same package as method
func (v Vertex) AbsMethod() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// function with same effect as receiver: Vertex is arg rather than Receiver
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// receiver can be a pointer (but not a pointer to a pointer)
// modifies the value to which receiver points
func (v *Vertex) ScaleMethodPointer(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// version with value receiver (instead of rather than pointer receiver)
// does NOT change value of v
func (v Vertex) ScaleMethodValue(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// function with pointer as arg
func ScaleFuncPoint(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// function with value as arg
func ScaleFuncValue(v Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Run_tour_4() {
	// 1. Method: a function with a special receiver arg
	// used to define a method on a type (Go does not have classes)
	v0 := Vertex{3, 4}
	fmt.Println(v0.AbsMethod())  // call function on the Vertex object, via receiver in fn signature
	
	// 2. Methods are functions, a function can be made to do same thing
	fmt.Println(AbsFunc(v0))   // version with Vertex as arg
	
	// 3. Methods can be declared on non-struct types ('type MyFloat float64')
	// type must be defined in same package as method
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.AbsMethod())
	
	// 4.a. Pointer receiver: changes the struct
	// receiver of a method can be a pointer, and modifies value to which receiver points
	v1 := Vertex{3, 4}
	fmt.Println("\nv1.AbsMethod() before ScaleMethodPointer: ", v1.AbsMethod())  // 5
	v1.ScaleMethodPointer(10)
	fmt.Println("v1.AbsMethod() after Scale: ", v1.AbsMethod())  // 50

	// 4.b. Value receiver: does NOT modify the struct
	// With a value receiver, the Scale method operates on a copy of the original Vertex value
	// This is the same behavior as for any other function argument
	// **receiver type is passed in SAME WAY, but method defined to handle as either val or pointer**
	v2 := Vertex{3, 4}
	fmt.Println("\nv2.AbsMethod() before ScaleMethodValue: ", v2.AbsMethod())  // 5
	v2.ScaleMethodValue(10)
	fmt.Println("v2.AbsMethod() after: ", v2.AbsMethod())  // 5 **does not change the struct**

	// 5.a. Pointers and functions
	// if function defined to take address as arg, DOES change the object (like pointer receiver in method)
	v3 := Vertex{3, 4}
	fmt.Println("\nAbsFunc(v3) before ScaleFuncPoint(&v3, 10): ", AbsFunc(v3))  // 5
	// ScaleFuncValue(&v3, 10)  // error: cannot use &v3 (value of type *Vertex) as Vertex value in arg
	ScaleFuncPoint(&v3, 10)
	fmt.Println("AbsFunc(v3) after: ", AbsFunc(v3))  // 50

	// 5.b. passing as value to function does NOT change it (like value receiver in method)
	v4 := Vertex{3, 4}
	fmt.Println("\nAbsFunc(v) before ScaleFuncValue: ", AbsFunc(v4))  // 5
	ScaleFuncValue(v4, 10)  // pass the struct
	fmt.Println("AbsFunc(v) after: ", AbsFunc(v4))  // 5

	// 6. Methods and pointer indirection
	
	// a. functions must take pointer or value as declared
    // var v5 Vertex  // created empty, initialize to show effects of function/methods
    v5 := Vertex{3, 4}
	// ScaleFuncPoint(v5, 5) // cannot use v5 (variable of type Vertex) as *Vertex value
    fmt.Println("\nv5 before ScaleFuncValue:", v5)
    ScaleFuncValue(v5, 5)
    fmt.Println("v5 after:", v5)  // function that receives value does NOT change v5
    
    p5 := &v5
    fmt.Println("\np5 before ScaleFuncPoint:", p5)
    ScaleFuncPoint(p5, 5)
    fmt.Println("p5 after:", p5)  // function that receives pointer DOES change v5
    // ScaleFuncValue(p5, 5) // cannot use p5 (variable of type *Vertex) as Vertex value

	// b. methods with pointer receivers take either value OR pointer as receiver
	// as a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5)
	// since the Scale method has a pointer receiver
	// all are allowed, though none 
	// var v6 Vertex
	v6 := Vertex{3, 4}
	fmt.Println("\nv6 before ScaleMethodPointer:", v6)
	v6.ScaleMethodPointer(5)  // OK
	fmt.Println("v6 after:", v6)  // method uses pointer, gets value: DOES change v6

	fmt.Println("\nv6 before ScaleMethodValue:", v6)
	v6.ScaleMethodValue(5)  // OK
	fmt.Println("v6 after:", v6)  // method uses value, gets value: does NOT change v6
	
	p6 := &v6
	fmt.Println("\nv6 before ScaleMethodPointer:", p6)
	p6.ScaleMethodPointer(10) // OK
	fmt.Println("p6 after:", p6)  // method uses pointer, gets pointer: DOES change v6

	fmt.Println("\np6 before ScaleMethodValue:", v6)
	p6.ScaleMethodValue(10) // OK
	fmt.Println("p6 after:", p6)  // method uses value, gets pointer: DOES change v6
}