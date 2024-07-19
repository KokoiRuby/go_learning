package main

import (
	"compress/bzip2"
	"compress/flate"
	"compress/gzip"
	"compress/lzw"
	"compress/zlib"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	filePath := "compressed_file.ext"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// Determine the compression format based on the file extension
	var reader io.Reader
	switch {
	case isBzip2(filePath):
		reader = bzip2.NewReader(file)
	case isZlib(filePath):
		reader, err = zlib.NewReader(file)
	case isGzip(filePath):
		reader, err = gzip.NewReader(file)
	case isLzw(filePath):
		reader = lzw.NewReader(file, lzw.MSB, 8)
	case isFlate(filePath):
		reader = flate.NewReader(file)
	default:
		fmt.Println("Unsupported compression format")
		return
	}

	// Read the compressed file content
	buffer := make([]byte, 1024)
	n, err := reader.Read(buffer)
	if err != nil && err != io.EOF {
		fmt.Println("Error reading compressed file:", err)
		return
	}

	// Process the compressed file content
	content := string(buffer[:n])
	fmt.Println("Compressed file content:")
	fmt.Println(content)
}

// Helper functions to determine the compression format based on file extension

func isBzip2(filePath string) bool {
	return filepath.Ext(filePath) == ".bz2"
}

func isZlib(filePath string) bool {
	return filepath.Ext(filePath) == ".zlib"
}

func isGzip(filePath string) bool {
	return filepath.Ext(filePath) == ".gz"
}

func isLzw(filePath string) bool {
	return filepath.Ext(filePath) == ".lzw"
}

func isFlate(filePath string) bool {
	return filepath.Ext(filePath) == ".flate"
}
