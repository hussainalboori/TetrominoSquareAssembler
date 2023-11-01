package main

import (
	"log"
	"os"
	"tetris-optimizer/Validation"
	// "time"
)

func main() {
	// startTime := time.Now()
	Inpout := os.Args[1:]
	if !validation.CheckIfValid(Inpout){
		log.Panicln("Error")
		return
		le
	}
}