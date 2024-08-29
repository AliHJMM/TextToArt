package functions

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func TextToArtInTerminal(args []string) {
	txt := args[2]

	format := args[3]
	str := ""
	outputPtr := flag.String("output", "", "Output file name")
	flag.Parse()

	if *outputPtr == "" {
		fmt.Println("Usage: go run . [STRING] [BANNER]\nEX: go run . something standard")
		os.Exit(1)
	}

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
					str += slice[firstLine]
				}
				str += "\n"
			}
		} else if j != len(textSlice)-1 {
			str += "\n"
		}
	}
	os.WriteFile(*outputPtr, []byte(str), 0o644)
}