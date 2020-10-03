package main

import "fmt"

func countAndSay(n int) string {
	ans := []byte{'1'}
	if n == 1 {
		return string(ans)
	}
	for i := 2; i <= n; i++ {
		next := make([]byte, 0, 2*len(ans))
		c, count := ans[0], 1
		for _, v := range ans[1:] {
			if v == c {
				count++
			} else {
				t := fmt.Sprintf("%d%c", count, c)
				next = append(next, []byte(t)...)
				c, count = v, 1
			} //if-else
		} //for
		t := fmt.Sprintf("%d%c", count, c)
		next = append(next, []byte(t)...)
		ans = next
	} //for i
	return string(ans)
}
