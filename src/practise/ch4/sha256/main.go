// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import (
	"crypto/sha256" //!+
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {
	var sz = flag.Int("SHA", 256, "The size [256,384] for sha generation")
	flag.Parse()
	switch *sz {
	case 256:
		c1 := sha256.Sum256([]byte("x"))
		c2 := sha256.Sum256([]byte("X"))
		fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
		// Output:
		// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
		// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
		// false
		// [32]uint8
		fmt.Println("the number of bits different between 2 sha256:", bitCount(&c1, &c2))
	case 384:
		c1 := sha512.Sum384([]byte("x"))
		c2 := sha512.Sum384([]byte("X"))
		fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
		// Output:
		// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
		// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
		// false
		// [32]uint8
		fmt.Println("the number of bits different between 2 sha384:", bitCount48(&c1, &c2))
	default:
		fmt.Println("wrong size")

	}
}

func bitCount(p1 *[32]byte, p2 *[32]byte) (n int) {
	for i := 0; i < len(*p1); i++ {
		n += popCount(p1[i] ^ p2[i])

	}
	return
}

func bitCount48(p1 *[48]byte, p2 *[48]byte) (n int) {
	for i := 0; i < len(*p1); i++ {
		n += popCount(p1[i] ^ p2[i])

	}
	return
}

func popCount(b byte) (c int) {
	for i := 0; i < 8; i++ {
		if (b>>i)&1 != 0 {
			c++
		}
	}
	return
}

//!-
