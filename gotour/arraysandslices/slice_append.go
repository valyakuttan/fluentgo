package arraysandslices

/*
Appending to a slice

It is common to append new elements to a slice, and so Go provides a built-in append function.
The documentation of the built-in package describes append.

func append(s []T, vs ...T) []T

The first parameter s of append is a slice of type T, and the rest are T values to append to the slice.

The resulting value of append is a slice containing all the elements of the original slice plus the provided values.

If the backing array of s is too small to fit all the given values a bigger array will be allocated.
The returned slice will point to the newly allocated array.
*/

import "fmt"

func SliceAppend() {
	var s []int
	printSlice1(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice1(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice1(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice1(s)
}

func printSlice1(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
