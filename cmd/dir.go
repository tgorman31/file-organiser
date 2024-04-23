package cmd

import (
	"cmp"
	tbl "file-organiser/style"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

type Dir struct {
	Name string
	Size int
}

var d []Dir

// Takes a string input of a Directory path and an Int of the level the data wishes to be returned at
func Get_Dir_Items(dir, fullPath string, dirLevel, currLevel int) int {
	t := tbl.CreateTable()
	var size int

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
		d = append(d, Dir{fld, size})
		if currLevel != 1 {
			t.Row(fld, s)
			sort_Directories(d)
		}
	}
	return size
}

// Takes a string input of a Directory path and an Int of the level the data wishes to be returned at
// and sorts it by size
func Get_Sorted_Dir(dir, fullPath string, dirLevel, currLevel int) []Dir {
	var size int

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
		fld := strings.Replace(fullPath, dir+"\\", "", -1)
		if currLevel != 1 {
			d = append(d, Dir{fld, size})
			sort_Directories(d)
		} else {
			d = append(d, Dir{"Total", size})
		}
	}
	return d
}

func sort_Directories(dirs []Dir) {
	slices.SortFunc(dirs,
		func(a, b Dir) int {
			return cmp.Compare(a.Size, b.Size)
		})
}

func getFileSize(file string) int {
	// Takes a file name and calculates its size
	fileInfo, err := os.Stat(file)
	check(err)
	var size int = int(fileInfo.Size())
	return size
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Readable_Size(bytes int) string {
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
