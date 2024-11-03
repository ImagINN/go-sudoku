package main

import "fmt"

const size = 9

func main() {
	list := []string{".96.4...1", "1...6...4", "5.481.39.", "..795..43", ".3..8....", "4.5.23.18", ".1.63..59", ".59.7.83.", "..359...7"}

	board := parseBoard(list)
	fmt.Println("Başlangıç Tablosu:")
	printBoard(board)

	if solveSudoku(board) {
		fmt.Println("\nÇözülen Sudoku:")
		printBoard(board)
	} else {
		fmt.Println("Çözüm bulunamadı.")
	}
}

func parseBoard(list []string) [][]string {
	board := make([][]string, size)
	for i := 0; i < size; i++ {
		board[i] = make([]string, size)
		for j := 0; j < size; j++ {
			if list[i][j] == '.' {
				board[i][j] = "."
			} else {
				board[i][j] = string(list[i][j])
			}
		}
	}
	return board
}

func solveSudoku(board [][]string) bool {
	row, col, found := findEmptyCell(board)
	if !found {
		return true
	}

	for num := '1'; num <= '9'; num++ {
		numStr := string(num) //rune değerini stringe çevirdim
		if isValidMove(board, row, col, numStr) {
			board[row][col] = numStr
			if solveSudoku(board) {
				return true
			}
			board[row][col] = "."
		}
	}
	return false
}

func findEmptyCell(board [][]string) (int, int, bool) {
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			if board[row][col] == "." {
				return row, col, true
			}
		}
	}
	return 0, 0, false
}

func isValidMove(board [][]string, row, col int, num string) bool {
	for i := 0; i < size; i++ {
		if board[row][i] == num || board[i][col] == num {
			return false
		}
	}

	startRow, startCol := (row/3)*3, (col/3)*3 //3x3'lük bloğun başlangıç koordinatları
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][startCol+j] == num {
				return false
			}
		}
	}
	return true
}

func printBoard(board [][]string) {
	for _, row := range board {
		for _, val := range row {
			fmt.Printf("%s ", val)
		}
		fmt.Println()
	}
}
