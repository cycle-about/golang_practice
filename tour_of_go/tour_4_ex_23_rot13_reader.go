package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

/*
A common pattern is an io.Reader that wraps another io.Reader, modifying the stream in some way.

For example, the gzip.NewReader function takes an io.Reader (a stream of compressed data) and 
returns a *gzip.Reader that also implements io.Reader (a stream of the decompressed data).

Implement a rot13Reader that implements io.Reader and reads from an io.Reader, 
modifying the stream by applying the rot13 substitution cipher to all alphabetical characters.

The rot13Reader type is provided for you. Make it an io.Reader by implementing its Read method. 
*/

type rot13Reader struct {
	r io.Reader
}

/*
Definition from slide 21: "Read populates the given byte slice with data and 
returns the number of bytes populated and an error value. 
It returns an io.EOF error when the stream ends."
*/
func (rot13 rot13Reader) Read(b []byte) (int, error) {
	reader_field := rot13.r
	n, err := reader_field.Read(b)
	// TODO substitute bytes in the array
	for i := 0; i < n; i++ {
        b[i] = char_rot13(b[i])
    }
	return n, err
}

func char_rot13(c byte) byte {
    switch {
    case 'A' <= c && c <= 'Z':
        return 'A' + (c-'A'+13)%26
    case 'a' <= c && c <= 'z':
        return 'a' + (c-'a'+13)%26
    default:
        return c // leave non-letters unchanged
    }
}

func tour_4_ex_rot13reader() {
	input := "Lbh penpxrq gur pbqr!"
    s := strings.NewReader(input)
	r := rot13Reader{s}
	fmt.Println("Output from main:")
	io.Copy(os.Stdout, &r)
}

/*

----------------

OUTPUT 4: gets reader field from rot13Reader, prints input string without substitution
	Lbh penpxrq gur pbqr!


func (rot13 rot13Reader) Read(b []byte) (int, error) {
	regReader := rot13.r
	n, err := regReader.Read(b)
	// TODO substitute bytes in the array
	return n, err
}


func tour_4_ex_rot13reader() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

----------------

OUTPUT 3: minimal return values
	[hangs until ctrl-c]

type rot13Reader struct {
	r io.Reader
}


func (rot13Reader) Read(b []byte) (int, error) {
	return len(b), nil
}


func tour_4_ex_rot13reader() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

----------------

OUTPUT 2: un-implemented Read method defined for rot13Reader receiver
	missing return

----------------

OUTPUT 1: template code, before implementing reader interface
	cannot use &r (value of type *rot13Reader) as io.Reader value in argument to io.Copy: 
*rot13Reader does not implement io.Reader (missing method Read)
*/