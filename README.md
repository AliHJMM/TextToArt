# TextToArt


# 💻 Tech Stack:
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

## DESCRIPTION
TextToArt is a command-line tool written in Go that converts text into ASCII art using different banners. The tool reads from a predefined text file and outputs the converted text either to the terminal or a specified output file, making it a versatile utility for generating stylized text art.


## AUTHOR
- [Ali Hasan](https://github.com/AliHJMM)

## Usage

### How to Run
1. Clone the Repo
2. Navigate to the project directory
3.1 Run `go run main.go "Text"`
3.2 Run `go run main.go "Text" "Banner"`
3.2 Run `go run main.go -output=output.txt "Your Text" "banner"`

### Implementation Details

#### Algorithm

TextToArt processes the input string by converting each character into its ASCII art equivalent based on the selected banner format. The ASCII art is then printed to the terminal or written to a file. The tool ensures that only valid characters are processed and handles different file outputs gracefully.

### Features
- Text to ASCII Art Conversion: Converts any string into ASCII art using a specified banner format.
- Terminal Display: Displays the ASCII art directly in the terminal.
- File Output: Saves the ASCII art into a file for further use or sharing.
