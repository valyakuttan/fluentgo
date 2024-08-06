package methods

/*

Pointer receivers
=================

You can declare methods with pointer receivers.

This means the receiver type has the literal syntax *T for some type T. (Also, T cannot itself be a pointer such as *int.)

Methods with pointer receivers can modify the value to which the receiver points.
Since methods often need to modify their receiver, pointer receivers are more common than value receivers.

Note:
-----

1. A functions with a pointer argument must take a pointer.

2. But methods with pointer receivers take either a value
   or a pointer as the receiver when they are called.

   var v Vertex
   v.Scale(5)  // OK
   p := &v
   p.Scale(10) // OK

   For the statement v.Scale(5), even though v is a value and not a pointer,
   the method with the pointer receiver is called automatically. That is, as
   a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5)
   since the Scale method has a pointer receiver.

3. Functions that take a value argument must take a value of that specific type.

4. While methods with value receivers take either a value or a pointer as the
   receiver when they are called.

   var v Vertex
   fmt.Println(v.Abs()) // OK
   p := &v
   fmt.Println(p.Abs()) // OK

   In this case, the method call p.Abs() is interpreted as (*p).Abs().

5. In general, all methods on a given type should have either value or pointer receivers,
   but not a mixture of both.

*/

import (
	"fmt"
	"math"
)

type vertex1 struct {
	X, Y float64
}

func (v vertex1) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *vertex1) scale(f float64) {
	v.X *= f
	v.Y *= f
}

func PointerReceiver() {
	v := vertex1{3, 4}
	v.scale(10)
	fmt.Println(v.abs())
}
