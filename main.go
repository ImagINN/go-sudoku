package main

import "fmt"

const size = 9

func main() {
	list := []string{".96.4...1", "1...6...4", "5.481.39.", "..795..43", ".3..8....", "4.5.23.18", ".1.63..59", ".59.7.83.", "..359...7"}
	// list := []string{"1.58.2...", ".9..764.5", "2..4..819", ".19..73.6", "762.83.9.", "....61.5.", "..76...3.", "43..2.5.1", "6..3.89.."}
	// list := []string{"..5.3..81", "9.285..6.", "6....4.5.", "..74.283.", "34976...5", "..83..49.", "15..87..2", ".9....6..", ".26.495.3"}
	// list := []string{"34.91..2.", ".96.8..41", "..8.2..7.", ".6..57.39", "1.2.6.7..", "97..3..64", "45.2.8..6", ".8..9..5.", "6.3..189."}
	// list := []string{"..73..4.5", "....2.9..", "253.6487.", ".9.74.36.", "....3..8.", "8362.9.47", "1..8.26.3", "6......18", ".8261...4"}
	// list := []string{"935..7..8", "...3.8.7.", "6..5..49.", ".73..4...", "4..175.8.", ".618..247", ".187.....", "..6.8.75.", "75.4.3862"}
	// list := []string{"..5.2...1", ".8735..46", "4...6.5..", ".5.9.....", ".7..3541.", "69314.857", "7415..6.8", "...284..5", "5.....3.4"}
	// list := []string{"..75...3.", "8..23...9", ".3479.86.", "..3..4198", ".4815...3", "..6.23..7", "351.6.78.", "4..31...6", ".7...5..2"}
	// list := []string{"53..7....", "6..195...", ".98....6.", "8...6...3", "4..8.3..1", "7...2...6", ".6....28.", "...419..5", "....8..79"}
	// list := []string{".932..8.", "27.3.85..", ".8.73.254", "9758...31", "....74.6.", "6.45.38.7", "7....2.48", "32.4...7.", "..8.579.."}
	// list := []string{"..213.748", "8.4.....2", ".178.26..", ".68.9.27.", ".932....4", "5..46.3..", "..9.24.23", "..63..19.", "385..1.2."}
	// list := []string{"not", "a", "sudoku"}

	if isInputValid(list) {
		board := parseBoard(list)
		fmt.Println("Başlangıç Tablosu:")
		printBoard(board)

		if isBoardValid(board) {
			if solveSudoku(board) {
				fmt.Println("\nÇözülen Sudoku:")
				printBoard(board)
			} else {
				fmt.Println("Çözüm bulunamadı.")
			}
		} else {
			fmt.Println("Geçersiz başlangıç tahtası")
		}
	} else {
		fmt.Println("Geçersiz giriş: Tahta 9x9 olmalıdır.")
	}
}

func isInputValid(list []string) bool {
	if len(list) != size {
		return false
	}
	for _, row := range list {
		if len(row) != size {
			return false
		}
	}
	return true
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

func isBoardValid(board [][]string) bool {
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			if board[row][col] != "." {
				num := board[row][col]
				board[row][col] = "." // Geçici olarak boş bırakıyoruz
				if !isValidMove(board, row, col, num) {
					return false
				}
				board[row][col] = num // Rakamı geri koyuyoruz
			}
		}
	}
	return true
}

func solveSudoku(board [][]string) bool {
	row, col, found := findEmptyCell(board)
	if !found {
		return true
	}

	for num := '1'; num <= '9'; num++ {
		numStr := string(num)
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

	startRow, startCol := (row/3)*3, (col/3)*3
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
