package main

import (
	"log"
	"os"
	"tetris-optimizer/Validation" // Import the validation package
	"tetris-optimizer/src" // Import the src package
)

func main() {
	// Get command-line arguments (excluding the program name)
	Inpout := os.Args[1:]

	// Check if the input arguments are valid using the CheckIfValid function
	if !validation.CheckIfValid(Inpout){
		// If the input is not valid, log an error and exit the program
		log.Panicln("Error")
		return
	}
	src.Run(Inpout[0])
}