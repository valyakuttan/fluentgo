package arraysandslices

import "fmt"

func SimpleSliceMain() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// The type []T is a slice with elements of type T.

	//  A slice is formed by specifying two indices, a low and high bound, separated by a colon:
	// a[low : high]
	var s []int = primes[:]

	fmt.Println(s)
}
