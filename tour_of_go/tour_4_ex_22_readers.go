package main

import (
	"fmt"
	"golang.org/x/tour/reader"
	"io"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader
// Implement a Reader type that emits an infinite stream of the ASCII character 'A'
// for any array of bytes that's the arg, modify it to contain all 'A'
// return value: how many bytes read, and error
func (MyReader) Read(b []byte) (int, error) {
	for i := range b {
        b[i] = 'A'
    }
    // this ALWAYS uses the entire array, never reaches EOF, so 'bytes read' is always len of arg array
    return len(b), nil
}

func tour_4_ex_reader() {	
    buf := make([]byte, 10)
    fmt.Printf("Before-read buffer len %d and contents: %s\n", len(buf), string(buf))  // Before-read buffer len 10 and contents: 
    
    r := &MyReader{}
    n, err := r.Read(buf)
    fmt.Printf("After-read buffer contents: %s\n", string(buf))  // After-read buffer contents: AAAAAAAAAA

    if err != nil && err != io.EOF {
        fmt.Println("Error:", err)
        return
    }

    fmt.Printf("Read %d bytes: %s\n", n, string(buf))  // Read 10 bytes: AAAAAAAAAA

	reader.Validate(MyReader{})
}


/*
MESSAGE 3.
read zero bytes after 1048576 Read calls

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (int, error) {
	return 0, nil
}

func tour_4_ex_reader() {
	reader.Validate(MyReader{})
}

-------------------------

ERROR 2. 
./tour_4_ex_22_readers.go:8:46: missing return

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read([]byte) (int, error) {}

func tour_4_ex_reader() {
	reader.Validate(MyReader{})
}

-------------------------

ERROR 1.
cannot use MyReader{} (value of type MyReader) as io.Reader value in argument to reader.Validate: 
MyReader does not implement io.Reader (missing method Read)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func tour_4_ex_reader() {
	reader.Validate(MyReader{})
}


*/