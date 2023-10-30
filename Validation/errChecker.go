package validation

import (
	
	"tetris-optimizer/FileHandler"
	"path/filepath"
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

}

func CheckFileExt(s string) bool {
	Ext := filepath.Ext(s)
	if Ext != ".txt" {
		return false
	}
	return true
}
