package functions

import (
	"log"
	"os"
)

// loadImage loads an image file and returns the byte slice
func LoadImage(filename string) []byte {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		log.Fatalf("Error getting file info: %v", err)
	}

	data := make([]byte, stat.Size())
	_, err = file.Read(data)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	return data
}
