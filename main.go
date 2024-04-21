package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func getDirItems(dir string) int64 {
	var size int64
	c, err := os.ReadDir(dir)
	check(err)

	// fmt.Println("Directory:", dir)

	for _, entry := range c {
		fullPath := filepath.Join(dir, entry.Name())
		if entry.IsDir() {
			size += getDirItems(fullPath)

		} else {
			size += getFileSize(fullPath)
		}
	}

	// fmt.Println(dir, "size =", size, "bytes")
	return size
}

func getFileSize(file string) int64 {
	// Takes a file name and calculates its size
	fileInfo, err := os.Stat(file)
	check(err)
	size := fileInfo.Size()
	return size
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Step1: Get Dir Items

func main() {
	var dir string
	test := true

	if !test {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter a directory path: ")

		dir, err := reader.ReadString('\n')

		check(err)
		dir = strings.TrimSpace(dir) //Handles windows \r\n for newlines

	} else {
		dir = "subdir"
		dir = "C:/Users/thoma/D&D"
	}

	size := getDirItems(dir)

	fmt.Println("Dir:", dir, "\nSize:", size)

}
