package methods

/*

Methods
=======

Go does not have classes. However, you can define methods on types.

A method is a function with a special receiver argument.

The receiver appears in its own argument list between the func keyword and the method name.

In this example, the Abs method has a receiver of type Vertex named v.

Note: A method is just a function with a receiver argument.

You can only declare a method with a receiver whose type is defined in the same package
as the method. You cannot declare a method with a receiver whose type is defined in another package
*/

import (
	"fmt"
	"math"
)

type vertex struct {
	X, Y float64
}

type myFloat float64

func (f myFloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v vertex) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Simple() {
	v := vertex{3, 4}
	fmt.Println(v.abs())

	f := myFloat(-math.Sqrt2)
	fmt.Println(f.abs())
}
