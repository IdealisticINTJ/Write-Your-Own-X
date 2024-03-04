
package main

import (
	"./challenge1"
	"flag"
	"fmt"
	"os"
)

func main() {
	var countBytes, countLines, countWords, countChars bool

	flag.BoolVar(&countBytes, "c", false, "Count bytes")
	flag.BoolVar(&countLines, "l", false, "Count lines")
	flag.BoolVar(&countWords, "w", false, "Count words")
	flag.BoolVar(&countChars, "m", false, "Count characters")

	flag.Parse()

	if !countBytes && !countLines && !countWords && !countChars {
		countBytes = true
		countLines = true
		countWords = true
	}

	filenames := flag.Args()

	if len(filenames) == 0 {
		challenge1.CountFromStdin(countBytes, countLines, countWords, countChars)
	} else {
		challenge1.CountFromFiles(filenames, countBytes, countLines, countWords, countChars)
	}
}
