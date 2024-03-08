package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	delimiter := flag.String("d", "\t", "delimiter character")
	fields := flag.String("f", "", "comma or whitespace-separated list of fields")
	flag.Parse()

	var indices []int
	if *fields != "" {
		for _, field := range strings.FieldsFunc(*fields, func(r rune) bool { return r == ',' }) {
			index := 0
			if _, err := fmt.Sscanf(field, "%d", &index); err == nil {
				indices = append(indices, index-1)
			}
		}
	}

	file := os.Stdin
	if flag.NArg() > 0 {
		var err error
		file, err = os.Open(flag.Arg(0))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
			os.Exit(1)
		}

		fields := strings.Split(line, *delimiter)
		var selectedFields []string
		for _, index := range indices {
			if index < len(fields) {
				selectedFields = append(selectedFields, fields[index])
			}
		}
		fmt.Println(strings.Join(selectedFields, *delimiter))
	}
}
