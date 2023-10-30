package validation

import (
	"path/filepath"
	"strings"
	"tetris-optimizer/FileHandler"
)

func CheckIfValid(Inpout []string) bool {
	if len(Inpout) != 1 {
		return false
	}

	if !CheckFileExt(Inpout[0]) {
		return false
	}

	if !FileHandler.FileExists(Inpout[0]) {
		return false
	}

	Data, err := FileHandler.ReadFile(Inpout[0])
	if err != nil {
		return false
	}
	
	if !CheckFileData(Data) {
		return false
	}

	blk := strings.Split(Data, "\n\n")
	for _, blks := range blk {
		if !ValidatTertromino(blks){
			return false
		}
	}
	return true
}

func CheckFileExt(s string) bool {
	Ext := filepath.Ext(s)
	if Ext != ".txt" {
		return false
	}
	return true
}

func CheckFileData(data string) bool{
	split := strings.Split(data, "/n/n")
	for _, block := range split {
		lines := strings.Split(strings.Trim(data, "\n\t"), "\n")
		NumRows := len(lines)
		NumClom := len(lines[0])
		if NumRows != 4 || NumClom != 4 || strings.Count(block, "#") != 4 || strings.Count(block, ".") != 12 {
			return false
		}
	}
	return true
}

func ValidatTertromino(input string) bool{
	lines := strings.Split(strings.Trim(input, "\n\t"), "\n")
	NumRows := len(lines)
	NumClom := len(lines[0])
	
	for i := 0; i < NumRows; i++ {
		lines[i] = "." + lines[i]+ "."
	}
	lines = append([]string{strings.Repeat(".", NumClom+2)}, lines...)
	lines = append(lines, strings.Repeat(".", NumClom+2))
	AdjBlocks := 0
	for i := 1; i <= NumRows; i++{
		for j :=1; j <= NumClom; j++ {
			if lines[i][j] == '#' {
				if lines[i-1][j] == '#' { 
					AdjBlocks++
				}
				if lines[i+1][j] == '#' { 
					AdjBlocks++
				}
				if lines[i][j-1] == '#' { 
					AdjBlocks++
				}
				if lines[i][j+1] == '#' { 
					AdjBlocks++
				}
			}
		}
	}
	return AdjBlocks == 6 || AdjBlocks == 8	
}
