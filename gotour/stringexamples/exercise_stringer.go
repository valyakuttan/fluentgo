package stringexamples

/*

Exercise: Stringers
===================

Make the IPAddr type implement fmt.Stringer to print the address as a dotted quad.

For instance, IPAddr{1, 2, 3, 4} should print as "1.2.3.4".

*/

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	return fmt.Sprintf("%v:%v:%v:%v", ip[0], ip[1], ip[2], ip[3])
}

func ExerciseStringer() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
