package maps

/*

Exercise: Maps

Implement WordCount. It should return a map of the counts of each “word” in the string s.
The wc.Test function runs a test suite against the provided function and prints success or failure.

You might find strings.Fields helpful.

*/

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wc := make(map[string]int, 0)

	for _, w := range strings.Fields(s) {
		wc[w]++
	}

	return wc
}

func Exercise() {
	wc.Test(WordCount)
}
