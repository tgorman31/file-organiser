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
	File []File
}

type File struct {
	Name string
	Size int
}

var D []Dir
var f []File

// Takes a string input of a Directory path and an Int of the level the data wishes to be returned at
func Get_Dir_Items_Size(dir, fullPath string, dirLevel, currLevel int) int {
	t := tbl.CreateTable()
	var size int

	c, err := os.ReadDir(fullPath)
	check(err)

	for _, entry := range c {
		fullPath := filepath.Join(fullPath, entry.Name())
		if entry.IsDir() {
			currLevel++
			size += Get_Dir_Items_Size(dir, fullPath, dirLevel, currLevel)
			currLevel--
		} else {
			size += getFileSize(fullPath)
			// f = append(f, File{fullPath, size})
		}
	}
	if dirLevel >= currLevel {
		s := Readable_Size(size)
		fld := strings.Replace(fullPath, dir+"\\", "", -1)
		D = append(D, Dir{fld, size, f})
		if currLevel != 1 {
			t.Row(fld, s)
			sort_Directories(D)
		}
	}
	return size
}

// Takes a string input of a Directory path and an Int of the level the data wishes to be returned at
func Get_Dir_Items(dir, fullPath string, dirLevel, currLevel int) (int, []Dir) {
	t := tbl.CreateTable()
	var size int
	var sz int
	var dr []Dir
	var fl []File

	c, err := os.ReadDir(fullPath)
	check(err)

	for _, entry := range c {
		fullPath := filepath.Join(fullPath, entry.Name())
		if entry.IsDir() {
			currLevel++
			sz, dr = Get_Dir_Items(dir, fullPath, dirLevel, currLevel)
			size += sz
			fullPath := strings.Replace(fullPath, dir+"\\", "", -1)
			fl = append(fl, File{fullPath, size})
			currLevel--
		} else {
			size += getFileSize(fullPath)
			fullPath := strings.Replace(fullPath, dir+"\\", "", -1)
			fl = append(fl, File{fullPath, size})
		}
	}
	dr = append(dr, Dir{fullPath, size, fl})
	if dirLevel >= currLevel {
		s := Readable_Size(size)
		fld := strings.Replace(fullPath, dir+"\\", "", -1)
		// D = append(D, Dir{fld, size, f})
		if currLevel != 1 {
			t.Row(fld, s)
			sort_Directories(dr)
			Write_to_file(dr, "tmp\\"+fmt.Sprint(fld)+".txt")
			D = append(D, Dir{fld, size, fl})
		}
	}

	return size, dr
}

// Get Dir Items Size returns an int that is the size of the dir
// Get File Size returns the size of the specified file

// I need to know the files that exist inside each Dir
// Dir Items if a Dir is found will call itself
// Dir 1
// -----> Dir 2
// ------------> Dir 3
// -------------------> File 1
// -------------------> File 2
// -------------------> File 3
// When we reach Dir 3 there is no more Dir's only files
// Get Dir Items Size uses Get File Size to loop through these items, total them up and pass the size up the chain
// This repeats until the the size of each Dir is calculated
// Try have Get_Dir_Items to return an int size and a Dir[]

// Takes a string input of a Directory path and an Int of the level the data wishes to be returned at
// and sorts it by size
func Get_Sorted_Dir(dir, fullPath string, dirLevel, currLevel int) []Dir {
	var size int

	c, err := os.ReadDir(fullPath)
	check(err)

	for _, entry := range c {
		fullPath := filepath.Join(fullPath, entry.Name())
		if entry.IsDir() {
			currLevel++
			size += Get_Dir_Items_Size(dir, fullPath, dirLevel, currLevel)
			currLevel--
		} else {
			size += getFileSize(fullPath)
			f = append(f, File{fullPath, size})
		}
	}
	if dirLevel >= currLevel {
		fld := strings.Replace(fullPath, dir+"\\", "", -1)
		if currLevel != 1 {
			D = append(D, Dir{fld, size, f})
			sort_Directories(D)
		} else {
			D = append(D, Dir{"Total", size, f})
		}
	}
	return D
}

func Write_to_file(dir []Dir, fileName string) {
	fl, err := os.Create(fileName)
	check(err)
	str := fmt.Sprintln(dir)
	fl.WriteString(str + "\n")
	fl.Close()
}

func sort_Directories(dirs []Dir) {
	slices.SortFunc(dirs,
		func(a, b Dir) int {
			return cmp.Compare(a.Size, b.Size)
		})
}

func sort_Files(files []File) {
	slices.SortFunc(files,
		func(a, b File) int {
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
