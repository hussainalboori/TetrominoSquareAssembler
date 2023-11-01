package src

import (
	"log"
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
		ShiftGrid[i] = strings.Replace("0", NumClo)
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
