package main

import (
	"fmt"
	"math"
	"time"
	// "strconv"
	// "io"
	// "strings"
)

// interface type: a set of method signatures
// value of interface (Abser) can hold any value that implements method Abs()
// defines function name and return type but NOT its arg or implementation (that's done by implementers)
type Abser interface {
	Abs() float64
}

type MyFloat float64

// method Abs() has type MyFloat as receiver
// so type MyFloat implements interface Abser: matches method name and return type
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

// differently implementedmethod Abs() has pointer receiver *Vertex
// so *Vertex implements interface Abser: matches method name and return type
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
// version that does not handle nil receiver
/*
func (t T) M() {
	fmt.Println(t.S)
}
*/

type F float64

func (f F) M() {
	fmt.Println(f)
}

// %v: Prints the value in a default format
// %T: Prints the type of the value
/*
func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
*/

// surprising: get error about redeclaration of describe unless above version commented out
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// method version of M with pointer receiver to a type T
// handles case where the receiver it was called on is nil
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

// empty interface as arg means it can be a value of any type
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

// implemented in fmt package: any value that implements 'String()' implements interface Stringer
/*
type Stringer interface {
    String() string
}
*/

type Person struct {
	Name string
	Age  int
}

// struct person implements interface 'Stringer' because it implements Stringer's method 'String()'
// Sprintf: formats according to a format specifier and returns the resulting string (https://pkg.go.dev/fmt#Sprintf)
// like overriding java toString()
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

// built-in interface, make types that implement it to customize error handling
// type error interface {
//     Error() string
// }

// struct for customized error handling, must have it implement Error() so implements interface 'error'
type MyError struct {
	When time.Time
	What string
}

// MyError implements interface 'error'
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

// returns pointer to new MyError type, with values assigned to its When and What fields
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func Run_tour_4() {
	/*
	// 9. Interface type: a set of method signatures
	// value declared to be an instance of that interface can hold any value that implements those methods
	var a Abser
	// fmt.Println("a and a.Abs() after declared: ", a, a.Abs())
	//   -> panic: runtime error: invalid memory address or nil pointer dereference
	
	f := MyFloat(-math.Sqrt2)
	a = f  // a MyFloat implements Abser
	fmt.Println("a and a.Abs() after defined as a float: ", a, a.Abs())
	// executes the verion of Abs() for its type
	//   -> -1.4142135623730951 1.4142135623730951

	v := Vertex{3, 4}
	a = &v // a *Vertex implements Abser
	fmt.Println("a and a.Abs() after defined as *Vertex: ", a, a.Abs())
	//   -> &{3 4} 5

	// In the following line, v is a Vertex (not *Vertex) and does NOT implement Abser.
	// a = v  // this statement fails
	// -> cannot use v (variable of type Vertex) as Abser value in assignment: Vertex does not implement Abser (method Abs has pointer receiver)

	// 10. interfaces implemented implicitly
	// A type implements an interface by implementing its methods, no explicit declaration
	// expect this should print "hello" - yep
	var i I = T{"hello"}
	i.M()

	// 11. Interface values
	// Under the hood, interface values can be thought of as a tuple of a value and a concrete type
	//     (value, type)
	// An interface value holds a value of a specific underlying concrete type
	// Calling a method on an interface value executes the method of the same name on its underlying type
	// here, types T and I have their own implementations of 'M()'

	var i I  // declare i. Type I is interface type, must implement 'M()' which has no return value

	// i = &T{"Hello"}  // define i as a pointer to struct type T, with its field S assigned to "hello"
	// describe(i)  // expect: type T -> no, actually a pointer to type T
	i.M()  // expect: hello
	// -> (&{Hello}, *main.T)
	// -> Hello

	// i = F(math.Pi)  // define i as fload value F, with numeric assignment to pi
	// describe(i)  // expect: type F
	i.M()  // expect: 3.14
	// -> (3.141592653589793, main.F)
	// -> 3.141592653589793

	// 12. Interface values with nil underlying values
	// if concrete value inside interface itself is nil, method called with nil receiver
	// common in go to write methods that gracefully handle being called with nil receiver
	// an interface value that holds a nil concrete value is itself non-nil
	var i I

	var t *T
	i = t  // assign i to be a pointer to T, though what it points to is nil
	describe(i)  // expect: type *T
	i.M() // expect: <nil>

	i = &T{"hello"}
	describe(i)  // expect: type *T, value hello
	i.M()  // expect: hello

	// 13. Nil interface value
	// holds neither value nor concrete type
	// calling method on nil interface is runtime error because no type inside the interface
	//     tuple to indicate which concrete method to call

	var i I  // declared but not assigned
	describe(i)
	i.M()
	// -> (<nil>, <nil>)
	//     panic: runtime error: invalid memory address or nil pointer dereference	

	// 14. Empty interface
	// interface that specifies no methods
	// it can hold values of any type
	// use for code that handles values of unknown type, eg fmt.Print()
	// needed a new version of describe() that accepts empty interface as arg
	var i interface{}
	describe(i)  // (<nil>, <nil>)
	fmt.Println(i)  // <nil>

	i = 42
	describe(i)  // (42, int)

	i = "hello"
	describe(i)  // (hello, string)

	// 15. Type assertion
	// provides access to interface value's underlying concrete value
	// type assertion: triggers a panic if incorrect type
	// type test: return the value, and a boolean for whether assertion is correct
	// syntax for type assertion of an interface's value is similar to reading from a map
	var i interface{} = "hello"  // i implements empty interface, and has underlying type string, value "hello"

	s := i.(string)  // if i is not a string, would trigger panic
	fmt.Println(s)  // expect: hello

	s, ok := i.(string)
	fmt.Println(s, ok)  // expect: hello, true

	f, ok := i.(float64)
	fmt.Println(f, ok)  // expect: hello, false
	// actually: 0, false. Seems if assertion fails it returns the default value for the type

	var t interface{} = 4
	q, ok := t.(string)
	fmt.Println(q, ok)  // " false"
	fmt.Println(q == "")  // true
	// yes, if the test of match was false, returns default value for asserted type (eg empty string)

	// f = i.(float64) // panic
	// fmt.Println(f)  // expect: not reached

	// 16. Type switches
	// permits several type assertions in series
	// like regular switch statement but cases specify types (not values)
	do(21)  // Twice 21 is 42
	do("hello")  // "hello" is 5 bytes long
	do(true)  // I don't know about type bool!
	do(nil)  // I don't know about type <nil>!

	// 17. Stringers
	// ubiquitous interface, defined in fmt package, which uses it to print values
	// a type that can describe itself as a string
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)

	// 18. Stringer ex

	// 19. Errors
	// error type is a built-in interface (similar to fmt.Stringer)
	// functions often return an error value, calling code should check for whether it is nil (meaning success)
	
	// return value of method run() is a value that implements interface error
	if err := run(); err != nil {
		fmt.Println(err)
	}

	// i, err := strconv.Atoi("42")  // Converted integer: 42
	i, err := strconv.Atoi("hi")	 // couldn't convert number: strconv.Atoi: parsing "hi": invalid syntax
	// i, err := strconv.Atoi(44) // does not allow this to be called at all 'cannot use 44 (untyped int constant) as string value'
	if err != nil {
	    fmt.Printf("couldn't convert number: %v\n", err)
	    return
	}
	fmt.Println("Converted integer:", i)

	// 20. Errors ex

	// 21. Readers
	// read end of a stream of data; implementatinos for files, network connections, compressors, ciphers, etc
	// Read populates given byte slice with data, returns the number of bytes populated, and error value
	// returns io.EOF when streams ends

	r := strings.NewReader("Hello, Reader!")

	// create slice of bytes with length 8
	b := make([]byte, 8)  // consumes output 8 bytes at a time (so 2 runes)
	for {
		n, err := r.Read(b)
		fmt.Printf("read %v bytes\n", n)  // added to confirm what return value means: how many bytes read, which may be LESS than size of input array
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	} 

	// PREDICTION
	// this string has 14 runes, rune is an alias for int32 (4 bytes)
	// 14 runes * 4 bytes per rune = 64 bytes
	// 64 bytes / 8 bytes per loop = 8 loops
	// expect it will print something like this for loops 1 and 2
	//     n = He err = <nil> b = [a pointer]
	//     b[:n] = llo, Reader!
	//     n = ll err = <nil> b = [a pointer]
	//     b[:n] = o, Reader!
	// last loop, since will have only 4 bytes
	//     n = ! err = io.EOF b = [a pointer]
	//     b[:n] = [empty string]

	// ACTUAL
	// 	n = 8 err = <nil> b = [72 101 108 108 111 44 32 82]
	// 	b[:n] = "Hello, R"
	// 	n = 6 err = <nil> b = [101 97 100 101 114 33 32 82]
	// 	b[:n] = "eader!"
	// 	n = 0 err = EOF b = [101 97 100 101 114 33 32 82]
	// 	b[:n] = ""

	// CORRECTIONS
	// 	[:n] prints from start of string, to where reader stopped [I had it backwards]
	// 	each loop actually read in equivalent of 8 letters each time, not 2 as expected for rune
	// 	1 letter per 8 bytes indicates UTF-8 encoding
	// 	Forgot that by defauts utf8 is used, not runes

	// 22. Exercise: Readers

	// 23. Exercise: rot13Reader
	*/
}
