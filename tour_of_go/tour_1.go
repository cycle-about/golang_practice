// Packages, variables, and functions

package main

import (
    "fmt"
    // "math/rand"  // it's an error to import but not use
    // "math"
    // "math/cmplx"
)

const (
    // Create a huge number by shifting a 1 bit left 100 places.
    // In other words, the binary number that is 1 followed by 100 zeroes.
    Big = 1 << 100
    // Shift it right again 99 places, so we end up with 1<<1, or 2.
    Small = Big >> 99
)

func needInt(x int) int {
    return x*10 + 1
}

func needFloat(x float64) float64 {
    return x * 0.1
}

// allowed assignments, outside of function
var i, j int = 1, 2  // works
// var i, j = 1, 2  // can omit type if initializing

// error out assignments, outside of function
// var i, j := 1, 2  // syntax error: unexpected :=, expected = 
// var i, j = 1  // assignment mismatch: 2 variables but 1 value
// i, j := 1, 2  // syntax error: non-declaration statement outside function body

// named return values: 'naked' return
// use only for short functions, or harms readability
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

// surprising: it's NOT an error to initialize class variable or function but not use it (unlike function var)
// var (
//     ToBe   bool       = false
//     MaxInt uint64     = 1<<64 - 1
//     z      complex128 = cmplx.Sqrt(-5 + 12i)
// )

func Run_tour_1() {
    /*
    // numbers can be printed with strings comma-separated
    fmt.Println("My favorite number is", rand.Intn(10))
    fmt.Println("My favorite number is", 3)
    fmt.Println("My favorite number is" + " three")
    // strings can be concatenated with +, but not numbers
    

    fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
    fmt.Printf("Now you have %d problems.\n", math.Sqrt(7))
    fmt.Printf("Now you have %s problems.\n", math.Sqrt(7))
    // interesting: wrong number type in % is warning but not error
    //    Now you have 2.6457513110645907 problems.
    //    Now you have %!d(float64=2.6457513110645907) problems.
    //    Now you have %!s(float64=2.6457513110645907) problems.

    // also interesting: going back to 'Exported names', local-socket verison saved my edit to code window, still running in another terminal

    // multiple return values are in (), single return values don't have to be
    // func swap(x, y string) (string, string) {
    //     return y, x

    fmt.Println(split(17))

    // allowed assignments, in function
    // var c, python, java = true, false, "no!"  // can omit type if initializing
    // c, python, java := true, false, "no!"  // don't use keyword 'var' with := syntax
    // i, j := 1, 2  // this COULDN'T run outside function

    // error-out assignments, in function
    // var c, python, java := true, false, "no!" // syntax error: unexpected :=, expected =

    // fmt.Println(i, j, c, python, java)

    // surprising it IS an error to assign var in function but not use it
    // var (
    //     ToBe   bool       = false
    //     MaxInt uint64     = 1<<64 - 1
    //     z      complex128 = cmplx.Sqrt(-5 + 12i)
    // )

    // fmt.Printf("Type: %T, Value: %v\n", ToBe, ToBe)
    // fmt.Printf("Type: %T, Value: %v\n", MaxInt, MaxInt)
    // fmt.Printf("Type: %T, Value: %v\n", z, z)
        // Type: bool Value: false
        // Type: uint64 Value: 18446744073709551615
        // Type: complex128 Value: (2+3i)

    // type conversion
    i := 42
    f := float64(i)
    u := uint(f)
    fmt.Printf("Type: %T Value: %v\n", i, i)
    fmt.Printf("Type: %T Value: %v\n", f, f)
    fmt.Printf("Type: %T Value: %v\n", u, u)
    // Type: int Value: 42
    // Type: float64 Value: 42
    // Type: uint Value: 42

    fmt.Println(needInt(Small))
    fmt.Println(needFloat(Small))
    fmt.Println(needFloat(Big))
    */

    fmt.Println(needInt(Small))
    fmt.Println(needFloat(Small))
    fmt.Println(needFloat(Big))

    // can reassign class var
    // i = 5

    // can't reassign class const
    // Small = 4  // cannot assign to Small (neither addressable nor a map index expression)
}
