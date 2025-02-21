package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    const oneRune = "廈"
    const nihongo = "日本語"

    input := nihongo
    rev := Reverse(input)
    doubleRev := Reverse(rev)
    fmt.Printf("original: %q\n", input)
    fmt.Printf("reversed: %q\n", rev)
    fmt.Printf("reversed again: %q\n", doubleRev)

    /*
    // ORIGINAL
    // initialize a string, call Reverse, print result, call Reverse again, print result again
    input := "The quick brown fox jumped over the lazy dog"
    rev := Reverse(input)
    doubleRev := Reverse(rev)
    fmt.Printf("original: %q\n", input)
    fmt.Printf("reversed: %q\n", rev)
    fmt.Printf("reversed again: %q\n", doubleRev)
    */
}


// accepts a string, loops over it one byte at a time, returns string with bytes reversed
func Reverse_original(s string) string {
    b := []byte(s)
    for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
        b[i], b[j] = b[j], b[i]
    }
    return string(b)
}

func Reverse(s string) string {
    var runes []rune
	
	for i, w := 0, 0; i < len(s); i += w {
        runeValue, width := utf8.DecodeRuneInString(s[i:])
        //fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
        runes = append(runes, runeValue)
        w = width
    }
    fmt.Printf("runes: %q\n", runes)

    for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func PrintBytes(s string) {
    b := []byte(s)
    for i := 0; i < len(b); i = i+1 {
        fmt.Printf("byte: %q\n", b[i])
    }
}

func PrintRunes(s string) {
	fmt.Printf("\nstring: %q\n", s)
	for i, w := 0, 0; i < len(s); i += w {
        runeValue, width := utf8.DecodeRuneInString(s[i:])
        fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
        w = width
    }
}
