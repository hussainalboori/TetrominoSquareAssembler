package src

import (
	"log"
	"regexp"
	"tetris-optimizer/FileHandler"
	"tetris-optimizer/algo"
)

func Run(FileName string) {
	data, _ := FileHandler.ReadFile(FileName)
	reegex := regexp.MustCompile(`\n\n+`)
	split := reegex.Split(data, -1)
	len := len(split)
	for i := 0; i < len; i++{
		split[i] = CleanData(split[i])
	}
	tetro := []*algo.Tetromino{}
	var char uint8 = 'A'
	for i := 0; i < len; i++{
		if char > 90 {
			char = 97
		}
		if char > 122 {
			log.Println("Alphabe are Over limet")
			return
		}
		tetro = append(tetro, ,)
	}
	
}