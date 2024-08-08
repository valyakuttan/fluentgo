package arraysandslices

/*
 A slice has both a length and a capacity.

The length of a slice is the number of elements it contains.

The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).

You can extend a slice's length by re-slicing it, provided it has sufficient capacity.
*/
import "fmt"

func SliceLenCapacity() {
	s := []int{2, 3, 5, 7, 11, 13}
	print_slice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	print_slice(s)

	// Extend its length.
	s = s[:4]
	print_slice(s)

	// Drop its first two values.
	s = s[2:]
	print_slice(s)
}

func print_slice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
