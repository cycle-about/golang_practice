package main

import (
    "fmt"
    "unicode/utf8"
)



func main() {
    const oneRune = "廈"
    const nihongo = "日本語"
    const c1 = "\xe8"
    const hello = "hello"

    // PrintRunes(input)
    input := "The quick brown fox jumped over the lazy dog"
    rev, revErr := Reverse(input)
    doubleRev, doubleRevErr := Reverse(rev)
    fmt.Printf("original: %q\n", input)
    fmt.Printf("reversed: %q, err: %v\n", rev, revErr)
    fmt.Printf("reversed again: %q, err: %v\n", doubleRev, doubleRevErr)

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

// my answer that handles reversing by rune
// func Reverse(s string) string {
//     var runes []rune
	
// 	for i, w := 0, 0; i < len(s); i += w {
//         runeValue, width := utf8.DecodeRuneInString(s[i:])
//         //fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
//         runes = append(runes, runeValue)
//         w = width
//     }
//     // fmt.Printf("runes: %q\n", runes)

//     for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
//         runes[i], runes[j] = runes[j], runes[i]
//     }
//     return string(runes)
// }

// from tutorial, to debug the "\xe8" double reverse error
// func Reverse(s string) string {
//     fmt.Printf("input: %q\n", s)
//     r := []rune(s)
//     fmt.Printf("runes: %q\n", r)
//     for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
//         r[i], r[j] = r[j], r[i]
//     }
//     return string(r)
// }

func PrintBytes(s string) {
    b := []byte(s)
    for i := 0; i < len(b); i = i+1 {
        fmt.Printf("byte: %q\n", b[i])
    }
}

// version to accept only valid utf8
func Reverse(s string) (string, error) {
    if !utf8.ValidString(s) {
        return s, errors.New("input is not valid UTF-8")
    }
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r), nil
}

func PrintRunes(s string) {

	fmt.Printf("runes in string: %q\n", s)
	for i, w := 0, 0; i < len(s); i += w {
        runeValue, width := utf8.DecodeRuneInString(s[i:])
        fmt.Printf("    %#U starts at byte position %d\n", runeValue, i)
        w = width
    }
}



