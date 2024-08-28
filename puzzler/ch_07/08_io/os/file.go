package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	// create & open file

	newFile, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer func(newFile *os.File) {
		err := newFile.Close()
		if err != nil {

		}
	}(newFile)
	fmt.Println("New file created:", newFile.Name())
	fd := newFile.Fd()
	fmt.Println("File descriptor:", fd)

	tempFile, err := os.CreateTemp("", "tempfile*.txt")
	if err != nil {
		fmt.Println("Error creating temporary file:", err)
		return
	}
	defer func(tempFile *os.File) {
		err := tempFile.Close()
		if err != nil {

		}
	}(tempFile)
	fmt.Println("Temporary file created:", tempFile.Name())

	fd = tempFile.Fd()
	file := os.NewFile(fd, tempFile.Name())
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	fmt.Println("File created with file descriptor:", file.Name())

	openedFile, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func(openedFile *os.File) {
		err := openedFile.Close()
		if err != nil {

		}
	}(openedFile)
	fmt.Println("File opened for reading:", openedFile.Name())

	openedFile, err = os.OpenFile("example.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file for read/write:", err)
		return
	}
	defer func(openedFile *os.File) {
		err := openedFile.Close()
		if err != nil {

		}
	}(openedFile)
	fmt.Println("File opened for read/write:", openedFile.Name())

	// Ch*

	err = file.Chdir()
	if err != nil {
		return
	}

	if err := file.Chmod(0666); err != nil {
		fmt.Println("Error changing file mode:", err)
	}

	if err := file.Chown(os.Getuid(), os.Getgid()); err != nil {
		fmt.Println("Error changing file owner and group:", err)
	}

	// Set*Deadline

	deadline := time.Now().Add(5 * time.Second)
	if err := file.SetDeadline(deadline); err != nil {
		fmt.Println("Error setting write deadline:", err)
	}

	readDeadline := time.Now().Add(5 * time.Second)
	if err := file.SetReadDeadline(readDeadline); err != nil {
		fmt.Println("Error setting read deadline:", err)
	}

	writeDeadline := time.Now().Add(5 * time.Second)
	if err := file.SetWriteDeadline(writeDeadline); err != nil {
		fmt.Println("Error setting write deadline:", err)
	}

	//

	file, err = os.OpenFile("example.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}
	fmt.Println("File info:", fileInfo.Name())

	newSize := int64(100)
	if err := file.Truncate(newSize); err != nil {
		fmt.Println("Error truncating file:", err)
		return
	}
	fmt.Println("File truncated successfully.")

	if err := file.Sync(); err != nil {
		fmt.Println("Error syncing file to disk:", err)
		return
	}
	fmt.Println("File synced to disk.")

	rawConn, err := file.SyscallConn()
	if err != nil {
		fmt.Println("Error getting raw file connection:", err)
		return
	}
	fmt.Println("Raw file connection:", rawConn)

	//

	file, err = os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	bytesToWrite := []byte("Hello, World!")
	n, err := file.Write(bytesToWrite)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Printf("Wrote %d bytes to the file.\n", n)

	off := int64(10)
	bytesToWriteAt := []byte("Golang")
	n, err = file.WriteAt(bytesToWriteAt, off)
	if err != nil {
		fmt.Println("Error writing to file at offset:", err)
		return
	}
	fmt.Printf("Wrote %d bytes at offset %d.\n", n, off)

	strToWrite := "This is a string."
	n, err = file.WriteString(strToWrite)
	if err != nil {
		fmt.Println("Error writing string to file:", err)
		return
	}
	fmt.Printf("Wrote %d bytes as string to the file.\n", n)

	n2, err := file.WriteTo(os.Stdout)
	if err != nil {
		fmt.Println("Error writing to io.Writer:", err)
		return
	}
	fmt.Printf("Wrote %d bytes to io.Writer.\n", n2)

	//

	file, err = os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// 读取文件内容到字节数组
	data := make([]byte, 100)
	n, err = file.Read(data)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Printf("Read %d bytes from the file: %s\n", n, string(data))

	off = int64(10)
	_, err = file.Seek(off, 0)
	if err != nil {
		fmt.Println("Error seeking file:", err)
		return
	}
	n, err = file.Read(data)
	if err != nil {
		fmt.Println("Error reading file at offset:", err)
		return
	}
	fmt.Printf("Read %d bytes at offset %d: %s\n", n, off, string(data))

	dirEntries, err := file.ReadDir(0)
	if err != nil {
		fmt.Println("Error reading directory entries:", err)
		return
	}
	fmt.Println("Directory entries:")
	for _, entry := range dirEntries {
		fmt.Println(entry.Name())
	}

	n2, err = file.ReadFrom(os.Stdin)
	if err != nil {
		fmt.Println("Error reading from io.Reader:", err)
		return
	}
	fmt.Printf("Read %d bytes from io.Reader.\n", n2)

	offset, err := file.Seek(0, 0)
	if err != nil {
		fmt.Println("Error seeking file:", err)
		return
	}
	fmt.Printf("File pointer moved to offset: %d\n", offset)

	fileInfos, err := file.Readdir(0)
	if err != nil {
		fmt.Println("Error reading directory info:", err)
		return
	}
	fmt.Println("Directory file info:")
	for _, info := range fileInfos {
		fmt.Println(info.Name())
	}

	names, err := file.Readdirnames(0)
	if err != nil {
		fmt.Println("Error reading directory names:", err)
		return
	}
	fmt.Println("Directory names:")
	for _, name := range names {
		fmt.Println(name)
	}

}
