package functions

import (
	"fmt"
	"os"
	"strings"
)

func TextToArtInFile(args []string) {
	txt := args[1]
	format := args[2]
	textSlice := strings.Split(txt, "\\n")

	if !charValidation(txt) {
		fmt.Println("Error : invalid char")
		os.Exit(1)
	}
	file, err := os.ReadFile(format + ".txt")
	if err != nil {
		fmt.Println("Error : reading file")
		os.Exit(1)
	}
	slice := strings.Split(string(file), "\n")
	for j, txt := range textSlice {
		if txt != "" {
			for i := 0; i < 8; i++ {
				for _, v := range txt {
					firstLine := int(v-32)*9 + 1 + i
					fmt.Print(slice[firstLine])
				}
				fmt.Println()
			}
		} else if j != len(textSlice)-1 {
			fmt.Println("")
		}
	}
}