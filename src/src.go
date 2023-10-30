package src

import (
	"regexp"
	"tetris-optimizer/FileHandler"
)

func Run(FileName string) {
	data, _ := FileHandler.ReadFile(FileName)
	reegex := regexp.MustCompile(`\n\n+`)
	split := reegex.Split(data, -1)
	
}