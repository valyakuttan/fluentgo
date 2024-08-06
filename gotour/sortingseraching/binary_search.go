package sortingseraching

import (
	"fmt"
	"sort"
)

func BinarySearchExample() {

	// binary search
	data := []int{1, 2, 3, 4, 23}
	x := 21
	i := sort.Search(len(data), func(i int) bool { return data[i] >= x })
	if i < len(data) && data[i] == x {
		fmt.Println(x, " is present at data[", i, "]")
	} else {
		fmt.Println(x, " is not present in data,but", i, "is the index where it would be inserted.")
	}

	names := []string{"Alice", "Bob", "Eve", "Sam"}
	target := "Charlie"
	i = sort.Search(len(names), func(i int) bool { return names[i] >= target })
	if i < len(names) && names[i] == target {
		fmt.Println(target, "is present at data[", i, "]")
	} else {
		fmt.Println(target, "is not present in data,but", i, "is the index where it would be inserted.")
	}

}
