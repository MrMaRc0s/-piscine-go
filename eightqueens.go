// Quest4
package piscine

import "github.com/01-edu/z01"

const N = 8

func EightQueens() {
	printSolution := func(solution []int) {
		for _, val := range solution {
			z01.PrintRune(rune(val + '1'))
		}
		z01.PrintRune('\n')
	}

	isSafe := func(board []int, row, col int) bool {
		for i := 0; i < col; i++ {
			if board[i] == row || board[i]-i == row-col || board[i]+i == row+col {
				return false
			}
		}
		return true
	}

	var solveNQueens func(board []int, col int)
	solveNQueens = func(board []int, col int) {
		if col == N {
			printSolution(board)
			return
		}
		for i := 0; i < N; i++ {
			if isSafe(board, i, col) {
				board[col] = i
				solveNQueens(board, col+1)
			}
		}
	}

	board := make([]int, N)
	solveNQueens(board, 0)
}
