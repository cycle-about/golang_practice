package main

import (
    "testing"
    "unicode/utf8"
)

// testing.T is a type provided by Go's built-in testing package, which is used for 
// writing and running tests. It represents the state of a single test

// creates a map of inputs and expected outputs for Reverse
// if result of Reverse does not match, returns an error
// param is testing.T
/*
func TestReverse(t *testing.T) {
    testcases := []struct {
        in, want string
    }{
        {"Hello, world", "dlrow ,olleH"},
        {" ", " "},
        {"!12345", "54321!"},
    }
    for _, tc := range testcases {
        rev := Reverse(tc.in)
        if rev != tc.want {
                t.Errorf("Reverse: %q, want %q", rev, tc.want)
        }
    }
}
*/

// param is Testing.F
// method f.Fuzz() takes as params: testing.T and the string to test
// raises error if reversing it twice is not the original, or not valid utf8
// provides three initial tests
func FuzzReverse(f *testing.F) {
    // this is what it runs without fuzzing
    testcases := []string{"Hello, world", " ", "!12345"}
    for _, tc := range testcases {
        f.Add(tc)  // Use f.Add to provide a seed corpus
    }
    // this is what it runs when fuzzing is called
    f.Fuzz(func(t *testing.T, orig string) {
        rev := Reverse(orig)
        doubleRev := Reverse(rev)
        if orig != doubleRev {
            t.Errorf("Before: %q, after: %q", orig, doubleRev)
        }
        if utf8.ValidString(orig) && !utf8.ValidString(rev) {
            t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
        }
    })
}