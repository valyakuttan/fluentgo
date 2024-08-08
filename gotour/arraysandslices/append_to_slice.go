package arraysandslices

import "fmt"

func AppendSlice() {
	// append values to a slice
	x := []int{1, 2, 3}
	x = append(x, 4, 5, 6)
	fmt.Println(x)

	// append another slice
	y := []int{4, 5, 6}
	x = append(x, y...)
	fmt.Println(x)
}
