package sprint

import (
	"strings"
)

func EightQueensSolver() string {
	size := 4
	var solutions []string
	board := make([]int, size) // Resizeable board

	var placeQueens func(row int)
	placeQueens = func(row int) {
		if row == size {
			solutions = append(solutions, formatBoard(board, size))
			return
		}
		for col := 0; col < size; col++ {
			if isSafe(board, row, col) {
				board[row] = col
				placeQueens(row + 1)
			}
		}
	}

	placeQueens(0)
	return strings.Join(solutions, "\n\n")
}

func isSafe(board []int, row, col int) bool {
	for i := 0; i < row; i++ {
		if board[i] == col ||
			board[i]-i == col-row ||
			board[i]+i == col+row {
			return false
		}
	}
	return true
}

func formatBoard(board []int, size int) string {
	var sb strings.Builder
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if board[i] == j {
				sb.WriteString("Q ")
			} else {
				sb.WriteString(". ")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
