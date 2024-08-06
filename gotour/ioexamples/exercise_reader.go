package ioexamples

/*

Exercise: Readers
=================

Implement a Reader type that emits an infinite stream of the ASCII character 'A'.

*/

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (n int, err error) {
	n, err = len(b), nil
	for i := 0; i < n; i++ {
		b[i] = 'A'
	}

	return
}

func ExerciseReader() {
	reader.Validate(MyReader{})
}
