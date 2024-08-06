package ioexamples

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadFileExample() {
	// open file
	f, err := os.Open("ioexamples/read_file_line_by_line.go")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// do something with a line
		fmt.Printf("%s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
