package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    /*
    const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

    fmt.Println("Println:")
    fmt.Println(sample)

    fmt.Println("\nByte loop:")
    for i := 0; i < len(sample); i++ {
        fmt.Printf("%x ", sample[i])
    }
    fmt.Printf("\n")

    fmt.Println("\nPrintf with %x:")
    fmt.Printf("%x\n", sample)

    fmt.Println("\nPrintf with % x:")
    fmt.Printf("% x\n", sample)

    fmt.Println("\nPrintf with %q:")
    fmt.Printf("%q\n", sample)

    fmt.Println("\nPrintf with %+q:")
    fmt.Printf("%+q\n", sample)
    

    //////////////////

    const placeOfInterest = `⌘`

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
    */

    //////////////////

    const nihongo = "日本語"
    for i, w := 0, 0; i < len(nihongo); i += w {
        runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
        fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
        w = width
    }
}