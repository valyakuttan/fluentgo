package concurrency

import (
	"fmt"
	"slices"
	"time"
)

func BasicSync() {

	c := make(chan int) // Allocate a channel.

	list := []int{1, 2, 5, 3}
	// Start the sort in a goroutine; when it completes, signal on the channel.
	go func() {
		slices.Sort(list)
		c <- 1 // Send a signal; value does not matter.
	}()
	
	fmt.Println("going to sleep")
	time.Sleep(10 * time.Millisecond)
	<-c // Wait for sort to finish; discard sent value.

	fmt.Println(list)
}
