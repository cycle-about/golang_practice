package main

// https://go.dev/blog/strings

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    
    const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

    fmt.Println("---- EXAMPLE 1 ----")
    fmt.Println("string:")
    fmt.Println(sample)

    // explicitly loop over every byte
    fmt.Println("\nByte loop:")
    for i := 0; i < len(sample); i++ {
        fmt.Printf("%x ", sample[i])
    }
    fmt.Printf("\n")

    // A shorter way to generate presentable output for a messy string is to use the %x (hexadecimal) format verb of fmt.Printf
    fmt.Println("\nPrintf with %x:")
    fmt.Printf("%x\n", sample)

    // add a space between each hex value
    fmt.Println("\nPrintf with % x:")
    fmt.Printf("% x\n", sample)

    // The %q (quoted) verb will escape any non-printable byte sequences in a string so the output is unambiguous.
    fmt.Println("\nPrintf with %q:")
    fmt.Printf("%q\n", sample)

    // the “plus” flag to the %q verb causes the output to escape not only non-printable sequences, but also any 
    // non-ASCII bytes, all while interpreting UTF-8. The result is that it exposes the Unicode values of properly 
    // formatted UTF-8 that represents non-ASCII data in the string:
    fmt.Println("\nPrintf with %+q:")
    fmt.Printf("%+q\n", sample)
    

    //////////////////

    //  To avoid any confusion, we create a “raw string”, enclosed by back quotes, so it can contain 
    // only literal text. (Regular strings, enclosed by double quotes, can contain escape sequences
    const placeOfInterest = `⌘`

    fmt.Println("\n---- EXAMPLE 2 ----")
    fmt.Printf("plain string: ")
    fmt.Printf("%s", placeOfInterest)
    fmt.Printf("\n")

    fmt.Printf("quoted string: ")
    fmt.Printf("%+q", placeOfInterest)
    fmt.Printf("\n")

    fmt.Printf("hex bytes: ")
    for i := 0; i < len(placeOfInterest); i++ {
        fmt.Printf("%x ", placeOfInterest[i])
    }
    fmt.Printf("\n")
    

    //////////////////

    fmt.Println("\n---- EXAMPLE 3 ----")
    const nihongo = "日本語"
    for i, w := 0, 0; i < len(nihongo); i += w {
        runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
        fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
        w = width
    }
}