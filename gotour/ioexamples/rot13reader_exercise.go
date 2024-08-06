package ioexamples

/*

Exercise: rot13Reader
=====================

A common pattern is an io.Reader that wraps another io.Reader, modifying the stream in some way.

For example, the gzip.NewReader function takes an io.Reader (a stream of compressed data) and
returns a *gzip.Reader that also implements io.Reader (a stream of the decompressed data).

Implement a rot13Reader that implements io.Reader and reads from an io.Reader, modifying
the stream by applying the rot13 substitution cipher to all alphabetical characters.

The rot13Reader type is provided for you. Make it an io.Reader by implementing its Read method.

*/

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13r rot13Reader) Read(b []byte) (n int, err error) {
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}

	n, err = rot13r.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = byte(rot13(rune(b[i])))
	}

	return

}

func ExerciseRot13Reader() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
