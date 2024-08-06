package pointer

import "fmt"

type Vertex struct {
	X int
	Y int
}

func StructMain() {
	v := vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
