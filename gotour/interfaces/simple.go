package interfaces

/*
Interfaces
==========

An interface type is defined as a set of method signatures.

A value of interface type can hold any value that implements those methods.


Interfaces are implemented implicitly

A type implements an interface by implementing its methods. There is no explicit
declaration of intent, no "implements" keyword.

Implicit interfaces decouple the definition of an interface from its implementation,
which could then appear in any package without prearrangement.

*/

import (
	"fmt"
	"math"
)

type myFloat float64

type vertex struct {
	X, Y float64
}

type Abser interface {
	abs() float64
}

func (f myFloat) abs() float64 {
	return float64(f)
}

func (v vertex) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Simple() {
	var a Abser
	f := myFloat(-math.Sqrt2)
	v := vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a Vertex implements Abser

	fmt.Println(a.abs())

	b := vertex{2.0, 5.0}
	fmt.Println(b.abs())
}
