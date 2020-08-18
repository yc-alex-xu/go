package main

func romanToInt(s string) int {
	m := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	ans, valPre := 0, 0
	for _, v := range s {
		valCur := m[v]
		if valPre < valCur {
			//因外上一个loop中，valPre被加到ans了
			ans = ans - 2*valPre + valCur
		} else {
			ans = ans + valCur
		}
		valPre = valCur
	}
	return ans
}
