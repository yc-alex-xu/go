package main

func convert(s string, numRows int) string {
	return fillByRow(s, numRows)
}

//模仿人工逐行填写过程
func fillByRow(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	n := min(len(s), numRows)
	rows := make([]string, n)
	lineNo, inc := 0, 1
	for _, ch := range s {
		rows[lineNo] += string(ch)
		//到头转向
		if lineNo == numRows-1 {
			inc = -1
		} else if lineNo == 0 {
			inc = 1
		}
		lineNo += inc
	}
	ans := ""
	for _, v := range rows {
		ans += v
	}
	return ans

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
