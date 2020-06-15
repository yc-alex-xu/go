// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, v := range os.Args[1:] {
		fmt.Printf("  %s\n", comma(v))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer

	var lenSign int
	if s[0] == '+' || s[0] == '-' {
		lenSign = 1
	}
	idxDot := strings.LastIndex(s, ".")
	var lenInt int
	if idxDot != -1 {
		lenInt = idxDot - lenSign
	} else {
		lenInt = len(s) - lenSign
	}

	buf.WriteString(s[:lenInt%3+lenSign])
	for i := 0; i < lenInt/3; i++ {
		if lenInt%3 > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(s[i*3+lenInt%3+lenSign : (i+1)*3+lenInt%3+lenSign])
	}
	if idxDot != -1 {
		buf.WriteString(s[idxDot:])
	}
	return buf.String()
}

//!-
