// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Generating random text: a Markov chain algorithm

Based on the program presented in the "Design and Implementation" chapter
of The Practice of Programming (Kernighan and Pike, Addison-Wesley 1999).
See also Computer Recreations, Scientific American 260, 122 - 125 (1989).

A Markov chain algorithm generates text by creating a statistical model of
potential textual suffixes for a given prefix. Consider this text:

	I am not a number! I am a free man!

Our Markov chain algorithm would arrange this text into this set of prefixes
and suffixes, or "chain": (This table assumes a prefix length of two words.)

	Prefix       Suffix

	"" ""        I
	"" I         am
	I am         a
	I am         not
	a free       man!
	am a         free
	am not       a
	a number!    I
	number! I    am
	not a        number!

To generate text using this table we select an initial prefix ("I am", for
example), choose one of the suffixes associated with that prefix at random
with probability determined by the input statistics ("a"),
and then create a new prefix by removing the first word from the prefix
and appending the suffix (making the new prefix is "am a"). Repeat this process
until we can't find any suffixes for the current prefix or we exceed the word
limit. (The word limit is necessary as the chain table may contain cycles.)

Our version of this program reads text from standard input, parsing it into a
Markov chain, and writes generated text to standard output.
The prefix and output lengths can be specified using the -prefix and -words
flags on the command-line.
*/
package idiomaticgo

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
)

/*
Modeling Markov chains
======================

A chain consists of a prefix and a suffix. Each prefix is a set number of words, while a suffix is a
single word. A prefix can have an arbitrary number of suffixes. To model this data, we use a
map[string][]string. Each map key is a prefix (a string) and its values are lists of suffixes
(a slice of strings, []string).

Here is the example table from the package comment as modeled by this data structure:

map[string][]string{
	" ":          {"I"},
	" I":         {"am"},
	"I am":       {"a", "not"},
	"a free":     {"man!"},
	"am a":       {"free"},
	"am not":     {"a"},
	"a number!":  {"I"},
	"number! I":  {"am"},
	"not a":      {"number!"},
}

While each prefix consists of multiple words, we store prefixes in the map as a single string.
It would seem more natural to store the prefix as a []string, but we can't do this with a
map because the key type of a map must implement equality (and slices do not).

Therefore, in most of our code we will model prefixes as a []string and join the strings
 together with a space to generate the map key:

Prefix               Map key

[]string{"", ""}     " "
[]string{"", "I"}    " I"
[]string{"I", "am"}  "I am"

*/

/*

The Prefix type
===============

Since we'll be working with prefixes often, we define a Prefix type with the concrete type
[]string. Defining a named type clearly allows us to be explicit when we are working with
a prefix instead of just a []string. Also, in Go we can define methods on any named type
(not just structs), so we can add methods that operate on Prefix if we need to.
*/

// Prefix is a Markov chain prefix of one or more words.
type Prefix []string

// NewPrefix returns a new Prefix length prefixLen.
func NewPrefix(prefixLen int) Prefix {
	return make(Prefix, prefixLen)
}

/*

The String method
=================

The first method we define on Prefix is String. It returns a string representation of a
Prefix by joining the slice elements together with spaces. We will use this method
to generate keys when working with the chain map.

*/
// String returns the Prefix as a string (for use as a map key).
func (p Prefix) String() string {
	return strings.Join(p, " ")
}

// Shift removes the first word from the Prefix and appends the given word.
func (p Prefix) Shift(word string) {
	copy(p, p[1:])
	p[len(p)-1] = word
}

/*

The Chain struct
================
The complete state of the chain table consists of the table itself and the word length
of the prefixes. The Chain struct stores this data.

*/
// Chain contains a map ("chain") of prefixes to a list of suffixes.
// A prefix is a string of prefixLen words joined with spaces.
// A suffix is a single word. A prefix can have multiple suffixes.
type Chain struct {
	chain     map[string][]string
	prefixLen int
}

/*

The NewChain constructor function
=================================

The Chain struct has two unexported fields (those that do not begin with an upper case character),
and so we write a NewChain constructor function that initializes the chain map with make and
sets the prefixLen field.

This is constructor function is not strictly necessary as this entire program is within a single
package (main) and therefore there is little practical difference between exported and unexported
fields. We could just as easily write out the contents of this function when we want to construct
a new Chain. But using these unexported fields is good practice; it clearly denotes that only
methods of Chain and its constructor function should access those fields. Also, structuring Chain
like this means we could easily move it into its own package at some later date.

*/

// NewChain returns a new Chain with prefixes of prefixLen words.
func NewChain(prefixLen int) *Chain {
	return &Chain{make(map[string][]string), prefixLen}
}

/*

Building the chain
==================

The Build method reads text from an io.Reader and parses it into prefixes and suffixes that
are stored in the Chain.

The io.Reader is an interface type that is widely used by the standard library and other Go
code. Our code uses the fmt.Fscan function, which reads space-separated values from an io.Reader.

The Build method returns once the Reader's Read method returns io.EOF (end of file) or some
other read error occurs.

Buffering the input
-------------------

This function does many small reads, which can be inefficient for some Readers. For efficiency
we wrap the provided io.Reader with bufio.NewReader to create a new io.Reader that provides buffering.

Scanning words
--------------

In our loop we read words from the Reader into a string variable s using fmt.Fscan. Since Fscan uses
space to separate each input value, each call will yield just one word (including punctuation),
which is exactly what we need.

Fscan returns an error if it encounters a read error (io.EOF, for example) or if it can't scan the
requested value (in our case, a single string). In either case we just want to stop scanning,
so we break out of the loop.

Adding a prefix and suffix to the chain
---------------------------------------

The word stored in s is a new suffix. We add the new prefix/suffix combination to the chain map
by computing the map key with p.String and appending the suffix to the slice stored under that key.

The built-in append function appends elements to a slice and allocates new storage when necessary.
 When the provided slice is nil, append allocates a new slice. This behavior conveniently ties
 in with the semantics of our map: retrieving an unset key returns the zero value of the value
 type and the zero value of []string is nil. When our program encounters a new prefix
 (yielding a nil value in the map) append will allocate a new slice.

For more information about the append function and slices in general see the Slices: usage
and internals article.

Pushing the suffix onto the prefix
----------------------------------

Before reading the next word our algorithm requires us to drop the first word from
the prefix and push the current suffix onto the prefix.

When in this state

p == Prefix{"I", "am"}
s == "not"

the new value for p would be

p == Prefix{"am", "not"}

This operation is also required during text generation so we put the code to perform
this mutation of the slice inside a method on Prefix named Shift.

*/

// Build reads text from the provided Reader and
// parses it into prefixes and suffixes that are stored in Chain.
func (c *Chain) Build(r io.Reader) {
	br := bufio.NewReader(r)
	p := NewPrefix(c.prefixLen)
	for {
		var s string
		if _, err := fmt.Fscan(br, &s); err != nil {
			break
		}
		key := p.String()
		c.chain[key] = append(c.chain[key], s)
		p.Shift(s)
	}
}

/*

Generating Text
===============

The Generate method is similar to Build except that instead of reading words
from a Reader and storing them in a map, it reads words from the map and appends
them to a slice (words).

Generate uses a conditional for loop to generate up to n words.

Getting potential suffixes
===========================

At each iteration of the loop we retrieve a list of potential suffixes for the current
prefix. We access the chain map at key p.String() and assign its contents to choices.

If len(choices) is zero we break out of the loop as there are no potential suffixes
for that prefix. This test also works if the key isn't present in the
map at all: in that case, choices will be nil and the length of a nil slice is zero.

Choosing a suffix at random
===========================

To choose a suffix we use the rand.Intn function. It returns a random integer up
to (but not including) the provided value. Passing in len(choices) gives us a
random index into the full length of the list.

We use that index to pick our new suffix, assign it to next and append it
to the words slice.

Next, we Shift the new suffix onto the prefix just as we did in the Build method.

Returning the generated text
============================

Before returning the generated text as a string, we use the strings.Join function
to join the elements of the words slice together, separated by spaces.

*/

// Generate returns a string of at most n words generated from Chain.
func (c *Chain) Generate(n int) string {
	p := make(Prefix, c.prefixLen)
	var words []string
	for i := 0; i < n; i++ {
		choices := c.chain[p.String()]
		if len(choices) == 0 {
			break
		}
		next := choices[rand.Intn(len(choices))]
		words = append(words, next)
		p.Shift(next)
	}
	return strings.Join(words, " ")
}

func MarkovTextGenerator() {
	// Register command-line flags.
	numWords := flag.Int("words", 100, "maximum number of words to print")
	prefixLen := flag.Int("prefix", 2, "prefix length in words")

	flag.Parse() // Parse command-line flags.

	c := NewChain(*prefixLen)     // Initialize a new Chain.
	c.Build(os.Stdin)             // Build chains from standard input.
	text := c.Generate(*numWords) // Generate text.
	fmt.Println(text)             // Write text to standard output.
}
