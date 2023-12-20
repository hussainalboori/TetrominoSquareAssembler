package algo

import "fmt"

// Board represents the tetro board.
type Board struct {
	Dim   int 		// Dimension of the square board
	Cells [][]uint8	// Matrix representing the board cells
}

// Tetromino represents a Tetris piece.
type Tetromino struct {
	Letter        uint8  
	Shape         []string
	Width, Height int
}

// CheckFit checks if a Tetromino can be placed at the specified row and column on the board.
func CheckFit(board *Board, tetromino *Tetromino, row, col int) bool {
	for i := 0; i < tetromino.Height; i++ {
		for j := 0; j < tetromino.Width; j++ {
			if row+i >= board.Dim || col+j >= board.Dim {
				return false
			}
			if tetromino.Shape[i][j] != '.' && board.Cells[row+i][col+j] != '.' {
				return false
			}
		}
	}
	return true
}

func PlaceTetromino(board *Board, tetromino *Tetromino, row, col int) {
	for i := 0; i < tetromino.Height; i++ {
		for j := 0; j < tetromino.Width; j++ {
			if tetromino.Shape[i][j] != '.' {
				board.Cells[row+i][col+j] = tetromino.Letter
			}
		}
	}
}

func RemoveTetromino(board *Board, tetromino *Tetromino, row, col int) {
	for i := 0; i < tetromino.Height; i++ {
		for j := 0; j < tetromino.Width; j++ {
			if tetromino.Shape[i][j] != '.' {
				board.Cells[row+i][col+j] = '.'
			}
		}
	}
}

func Solve(board *Board, tetrominos []*Tetromino, n int) bool {
	if len(tetrominos) == n {
		return true // Base case
	}
	for row := 0; row < board.Dim; row++ {
		for col := 0; col < board.Dim; col++ {
			if CheckFit(board, tetrominos[n], row, col) {
				PlaceTetromino(board, tetrominos[n], row, col)
				if Solve(board, tetrominos, n+1) {
					return true
				}
				RemoveTetromino(board, tetrominos[n], row, col)
			}
		}
	}
	return false
}

func PrintProcess(board [][]uint8) {
	for _, sl := range board {
		for _, b := range sl {
			fmt.Printf("%v ", string(b))
		}
		fmt.Println()
	}
}
