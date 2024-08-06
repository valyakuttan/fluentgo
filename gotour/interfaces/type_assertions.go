package interfaces

/*


Type assertions
===============

A type assertion provides access to an interface value's underlying concrete value.

t := i.(T)

This statement asserts that the interface value i holds the concrete type T
and assigns the underlying T value to the variable t.

If i does not hold a T, the statement will trigger a panic.

To test whether an interface value holds a specific type, a type assertion can
return two values: the underlying value and a boolean value that reports whether
the assertion succeeded.

t, ok := i.(T)

If i holds a T, then t will be the underlying value and ok will be true.

If not, ok will be false and t will be the zero value of type T, and no panic occurs.

Note the similarity between this syntax and that of reading from a map.


Type switches
=============

A type switch is a construct that permits several type assertions in series.

A type switch is like a regular switch statement, but the cases in a type switch specify
types (not values), and those values are compared against the type of the value held by the given interface value.

switch v := i.(type) {
case T:
    // here v has type T
case S:
    // here v has type S
default:
    // no match; here v has the same type as i
}

The declaration in a type switch has the same syntax as a type assertion i.(T), but the
 specific type T is replaced with the keyword type.

This switch statement tests whether the interface value i holds a value of type T or S.
In each of the T and S cases, the variable v will be of type T or S respectively and
hold the value held by i. In the default case (where there is no match), the variable
v is of the same interface type and value as i.


*/

import "fmt"

func TypeAssertion() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)

	// type switch
	do(21)
	do("hello")
	do(true)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
