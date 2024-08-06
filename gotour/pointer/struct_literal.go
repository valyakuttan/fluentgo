package pointer

import "fmt"

type vertex struct {
	X, Y int
}

var (
	v1 = vertex{1, 2}  // has type Vertex
	v2 = vertex{X: 1}  // Y:0 is implicit
	v3 = vertex{}      // X:0 and Y:0
	p  = &vertex{1, 2} // has type *Vertex
)

func StructLiteralMain() {
	fmt.Println(v1, p, v2, v3)
}
