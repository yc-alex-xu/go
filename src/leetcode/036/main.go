package main

func isValidSudoku(board [][]byte) bool {
	m := len(board)
	if m == 0 {
		return false
	}
	n := len(board[0])
	if n == 0 {
		return false
	}
	for i := 0; i < m; i++ {
		//index 0 不用
		row := make([]byte, 10)
		for j := 0; j < n; j++ {
			t := board[i][j]
			if t != '.' {
				t = t - '0'
				//同一行出现重复 3*3方块出现重复
				if row[t] > 0 {
					return false
				}
				row[t]++
			} //if
		} //for j
	} //for i

	for j := 0; j < n; j++ {
		col := make([]byte, 10)
		for i := 0; i < m; i++ {
			t := board[i][j]
			if t != '.' {
				t = t - '0'
				if col[t] > 0 {
					return false
				}
				col[t]++
			} //if
		} //for i
	} //for j

	for i := 0; i < m/3; i++ {
		for j := 0; j < n/3; j++ {
			cube := make([]byte, 10)
			for k := 3 * i; k < 3*i+3; k++ {
				for l := 3 * j; l < 3*j+3; l++ {
					t := board[k][l]
					if t != '.' {
						t = t - '0'
						if cube[t] > 0 {
							return false
						}
						cube[t]++
					} //if
				} //for l
			} //for k
		} //for j
	} //for i

	return true
}
