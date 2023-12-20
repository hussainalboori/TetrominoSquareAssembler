package src

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strings"
	"tetris-optimizer/FileHandler"
	"tetris-optimizer/algo"
)

func Run(FileName string) {
	data, _ := FileHandler.ReadFile(FileName)
	reegex := regexp.MustCompile(`\n\n+`)
	split := reegex.Split(data, -1)
	len := len(split)
	for i := 0; i < len; i++ {
		split[i] = CleanData(split[i])
	}
	tetro := []*algo.Tetromino{}
	var char uint8 = 'A'
	for i := 0; i < len; i++ {
		if char > 90 {
			char = 97
		}
		if char > 122 {
			log.Println("Alphabe are Over limet")
			return
		}
		tetro = append(tetro, NewTetro(split[i], char))
		char++
	}
	SmallestSquareFound := false
	for Dim := int(math.Ceil(math.Sqrt(float64(len*4)))); Dim < 10; Dim++{
		fmt.Printf("Square Dimension: %d\n", Dim)
		cells := GenerateBoard(Dim)
		board := &algo.Board{
			Dim:   Dim,
			Cells: cells,
		}
		if algo.Solve(board, tetro, 0) {
			SmallestSquareFound = true
			fmt.Println("Smallest square found!")
			PrintBoard(board)
			break
		}

		fmt.Println("-------------------------------------------------")
	}
	if !SmallestSquareFound {
		fmt.Println("No Solution Found. ")
	}

}

func NewTetro(s string, char uint8) *algo.Tetromino {
	line := strings.Split(s, "\n")
	line = ShiftPiece(line)
	return &algo.Tetromino{
		Letter: char,
		Shape:  line,
		Width:  len(line[0]),
		Height: len(line),
	}

}

func ShiftPiece(grid []string) []string {
	MainRow, MaxRow, MainCol, MaxCol := len(grid), 0, len(grid[0]), 0

	for row, RowStr := range grid {
		for col, char := range RowStr {
			if char == '#' {
				if row < MainRow {
					MainRow = row
				}
				if row > MaxRow {
					MaxRow = row
				}
				if col < MainCol {
					MainCol = col
				}
				if col > MaxCol {
					MaxCol = col
				}
			}
		}
	}
	NumRow := MaxRow - MainRow + 1
	NumClo := MaxCol - MainCol + 1

	ShiftGrid := make([]string, NumRow)
	for i := range ShiftGrid {
		ShiftGrid[i] = strings.Repeat(".", NumClo)
	}

	for Row := MainRow; Row <= MaxRow; Row++ {
		for Col := MainCol; Col <= MaxCol; Col++ {
			if grid[Row][Col] == '#' {
				ShiftGrid[Row-MainRow] = ShiftGrid[Row-MainRow][:Col-MainCol] + "#" + ShiftGrid[Row-MainRow][Col-MainCol+1:]
			}
		}
	}
	return ShiftGrid
}

func GenerateBoard(Dim int) [][]uint8{
	BoardForm := make([][]uint8, Dim)
	for i := 0; i < Dim; i++{
		BoardForm[i] = make([]uint8, Dim)
		for j := 0 ; j < Dim; j++{
			BoardForm[i][j] = '.'
		}
	}
	return BoardForm
}

func PrintBoard(Bord *algo.Board){
	for _, Row := range Bord.Cells{
		fmt.Println(string(Row))
	}
	fmt.Println()
}