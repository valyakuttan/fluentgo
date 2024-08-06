package interfaces

/*

Interface values

Under the hood, interface values can be thought of as a tuple of a value and a concrete type:

(value, type)

An interface value holds a value of a specific underlying concrete type.

Calling a method on an interface value executes the method of the same name on its underlying type.


Interface values with nil underlying values
-------------------------------------------

If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.

In some languages this would trigger a null pointer exception, but in Go it is common to write methods
that gracefully handle being called with a nil receiver (as with the method M in this example.)

Note that an interface value that holds a nil concrete value is itself non-nil.


Nil interface values
--------------------

A nil interface value holds neither value nor concrete type.

Calling a method on a nil interface is a run-time error because there is no type inside the interface
tuple to indicate which concrete method to call.


The empty interface
-------------------

The interface type that specifies zero methods is known as the empty interface:

interface{}

An empty interface may hold values of any type. (Every type implements at least zero methods.)

Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print
takes any number of arguments of type interface{}.

*/

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func InterfaceValues() {
	var i I // nil interface value
	describe(i)
	// i.M() error, call method on nil interface

	i = &T{"Hello"}
	describe(i)
	i.M()

	var t *T
	i = t
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
