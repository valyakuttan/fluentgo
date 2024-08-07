package concurrency

import "fmt"

/*
Exercise: Equivalent Binary Trees
=================================

There can be many different binary trees with the same sequence of values
stored in it.

A function to check whether two binary trees store the same sequence is quite
complex in most languages. We'll use Go's concurrency and channels to write a simple solution.

This example uses the tree package, which defines the type:

type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}

1. Implement the Walk function.

2. Test the Walk function.

The function New(k) constructs a randomly-structured (but always sorted) binary
tree holding the values k, 2k, 3k, ..., 10k.

Create a new channel ch and kick off the walker:

go Walk(tree.New(1), ch)

Then read and print 10 values from the channel. It should be the numbers 1, 2, 3, ..., 10.

3. Implement the Same function using Walk to determine whether t1 and t2 store the same values.

4. Test the Same function.

Same(tree.New(1), tree.New(1)) should return true, and Same(tree.New(1), tree.New(2)) should return false.
*/

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *Tree, ch chan int) {
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}

}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *Tree, size int) bool {
	ch1, ch2 := make(chan int, size), make(chan int, size)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	result := make(chan bool, 1)

	go func() {
		for i := 0; i < size; i++ {
			if <-ch1 != <-ch2 {
				result <- false
				return
			}
		}
		result <- true
	}()

	return <-result
}

func ExerciseEqBTree() {

	x55 := Same(New(5), New(5), 10)
	x56 := Same(New(5), New(6), 10)
	x65 := Same(New(6), New(5), 10)
	x66 := Same(New(6), New(6), 10)

	fmt.Println(x55, x56, x65, x66)
}
