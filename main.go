package main

import (
	"fmt"
	"os"
	"TextToArt/functions"
)

func main() {
	args := os.Args 
	if len(args) == 2 {
		TextToArt(args) 
	} else if len(args) == 3 {
		TextToArtInTerminal(args) 
	} else if len(args) == 4 {
		TextToArtInFile(args)
	} else {
		fmt.Println("Usage: go run . [STRING] [BANNER]\nEX: go run . something standard")
	}
}