package readfile

import (
	"bufio"
	"log"
	"os"
)

// GetFileContents is a helper function for reading txt files
func GetFileContents(filePath string) []string {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var listOfLines []string

	for scanner.Scan() {
		listOfLines = append(listOfLines, scanner.Text())
	}

	file.Close()

	return listOfLines
}