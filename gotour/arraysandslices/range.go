package arraysandslices

/*
Range

The range form of the for loop iterates over a slice or map.

When ranging over a slice, two values are returned for each iteration.

The first is the index, and the second is a copy of the element at that index.

 You can skip the index or value by assigning to _.

for i, _ := range pow
for _, value := range pow

If you only want the index, you can omit the second variable.

for i := range pow

*/

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func Range() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
