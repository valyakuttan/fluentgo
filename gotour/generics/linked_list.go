package generics

import (
	"fmt"
)

/*
Generic types
=============

In addition to generic functions, Go also supports generic types. A type
can be parameterized with a type parameter, which could be useful
for implementing generic data structures.

This example demonstrates a simple type declaration for a singly-linked
list holding any type of value.

*/

// Node represents a singly-linked list that holds
// values of any type.
type Node[T comparable] struct {
	next *List[T]
	elem T
}

type List[T comparable] struct {
	head *Node[T]
}

func NewList[T comparable]() *List[T] {
	return new(List[T])
}

func NewNode[T comparable](val T) *Node[T] {
	n := new(Node[T])
	n.elem = val
	return n
}

func (lst *List[T]) Push(elem T) {
	l := NewList[T]()
	l.head = lst.head

	n := NewNode(elem)
	n.next = l
	lst.head = n
}

func (lst *List[T]) Pop() (T, error) {
	if lst.head == nil {
		var t T
		return t, fmt.Errorf("can't pop from an empty list")
	}

	n := lst.head
	lst.head = n.next.head

	return n.elem, nil
}

func (lst *List[T]) Print() {
	if lst.head != nil {
		fmt.Print(lst.head.elem, " -> ")
		lst.head.next.Print()
	} else {
		fmt.Println(nil)
	}
}
func GenericTypeExample() {
	lst := NewList[int]()

	lst.Push(1)
	lst.Push(2)
	lst.Push(3)

	lst.Print()
	x, _ := lst.Pop()
	lst.Print()
	fmt.Println(x)

	lst.Pop()
	lst.Print()

	lst.Pop()
	lst.Print()

	x, ok := lst.Pop()
	fmt.Println(x, ok)
}
