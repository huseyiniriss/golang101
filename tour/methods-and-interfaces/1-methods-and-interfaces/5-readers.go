package main

import (
	"fmt"
	"io"
	"strings"
)

//type rot13Reader struct {
//	r io.Reader
//}

// type MyReader struct{}
func main() {
	// Readers
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	// Exercise: Readers
	//import "golang.org/x/tour/reader"
	//reader.Validate(MyReader{})

	// Exercise: rot13Reader
	//s := strings.NewReader("Lbh penpxrq gur pbqr!")
	//r = rot13Reader{s}
	//io.Copy(r, s)
}
