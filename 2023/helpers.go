package helper

import (
	"log"
	"os"
)

func OpenFile(filePath string) *os.File {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return file
}