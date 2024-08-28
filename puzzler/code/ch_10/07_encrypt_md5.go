package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func md5HashString(data string) string {
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

func md5HashFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hasher := md5.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func main() {
	data := "hello, world"
	hash := md5HashString(data)
	fmt.Printf("MD5 hash of '%s': %s\n", data, hash)

	filename := "testfile.txt"
	fileContent := []byte("This is a test file.go.")
	if err := os.WriteFile(filename, fileContent, 0644); err != nil {
		fmt.Println("Error creating file.go:", err)
		return
	}

	fileHash, err := md5HashFile(filename)
	if err != nil {
		fmt.Println("Error hashing file.go:", err)
		return
	}
	fmt.Printf("MD5 hash of file.go '%s': %s\n", filename, fileHash)

	if err := os.Remove(filename); err != nil {
		fmt.Println("Error removing file.go:", err)
	}
}
