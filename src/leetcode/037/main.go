package main

import "fmt"

func solveSudoku(board [][]byte) {
	b := backtrace(board, 0, 0)
	fmt.Println(b)
}

//从当前board出发，是否有解
func solved(board [][]byte, row int, col int) bool {
	if row == 8 && col == 8 {
		return true
	}
	if col == 8 {
		return backtrace(board, row+1, 0)
	}
	return backtrace(board, row, col+1)

}

//先col,后row的尝试
func backtrace(board [][]byte, row int, col int) bool {
	if board[row][col] != '.' {
		return solved(board, row, col)
	}
	for c := byte('1'); c <= byte('9'); c++ {
		{ //试着放入c
			if isOK(board, row, col, c) && solved(board, row, col) {
				return true
			}
			//回嗍
			board[row][col] = '.'
		} //if
	} //for
	return false
}

//判断c与已有的board是否冲突
func isOK(board [][]byte, row int, col int, c byte) bool {
	for i := 0; i < 9; i++ {
		// 判断行是否存在重复
		if board[row][i] == c {
			return false
		}
		// 判断列是否存在重复
		if board[i][col] == c {
			return false
		}
		// 判断单元格是否存在重复,将3*3格子编号了
		if board[(row/3)*3+i/3][(col/3)*3+i%3] == c {
			return false
		}
	}
	board[row][col] = c
	return true
}
