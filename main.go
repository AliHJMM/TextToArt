package main

import (
	"fmt"
	"os"
	"TextToArt/functions"
)

func main() {
	args := os.Args 
	if len(args) == 2 {
		functions.TextToArt(args) 
	} else if len(args) == 3 {
		functions.TextToArtInTerminal(args) 
	} else if len(args) == 4 {
		functions.TextToArtInFile(args)
	} else {
		fmt.Println("Usage: go run . [STRING] [BANNER]\nEX: go run . something standard")
	}
}