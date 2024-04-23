package cmd

import (
	tbl "file-organiser/style"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Takes a string input of a Directory path and an Int of the level the data wishes to be returned at
func Get_Dir_Items(dir, fullPath string, dirLevel, currLevel int) int64 {
	t := tbl.CreateTable()
	var size int64

	c, err := os.ReadDir(fullPath)
	check(err)

	// fmt.Println("Directory:", dir)

	for _, entry := range c {
		fullPath := filepath.Join(fullPath, entry.Name())
		if entry.IsDir() {
			currLevel++
			size += Get_Dir_Items(dir, fullPath, dirLevel, currLevel)
			currLevel--
		} else {
			size += getFileSize(fullPath)
		}
	}
	if dirLevel >= currLevel {
		s := Readable_Size(size)
		fld := strings.Replace(fullPath, dir+"\\", "", -1)
		if currLevel != 1 {
			t.Row(fld, s)
		}
	}
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

func Readable_Size(bytes int64) string {
	var size float64
	var sizeType string
	switch {
	case bytes >= 1000000000:
		size = float64(bytes) / 1000000000
		sizeType = "GB"
	case bytes >= 1000000:
		size = float64(bytes) / 1000000
		sizeType = "MB"
	case bytes >= 1000:
		size = float64(bytes) / 1000
		sizeType = "KB"
	default:
		size = float64(bytes)
		sizeType = " B"
	}
	mbSize := fmt.Sprintf("%.2f "+sizeType, size)

	return mbSize
}
