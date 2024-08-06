package function

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func Multi_return_function() {
	var a, b = "hello", "world"
	a, b = swap(a, b)

	fmt.Println(a, b)
}
