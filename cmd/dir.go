package cmd

import (
	"cmp"
	// tbl "file-organiser/style"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

type Dir struct {
	Name  string
	Size  int
	Depth int
	File  []File
}

type File struct {
	Name string
	Size int
}

// var d []Dir
// var f []File

// Function to gather directories
func Gather_Directories(dir, fullPath string, dirLevel, currLevel int) []Dir {
	var dr []Dir

	c, err := os.ReadDir(fullPath)
	check(err)

	for _, entry := range c {
		entryPath := filepath.Join(fullPath, entry.Name())

		if entry.IsDir() {

			subDirs := Gather_Directories(dir, entryPath, dirLevel, currLevel+1)

			dr = append(dr, subDirs...)

			fld := strings.Replace(entryPath, dir+"\\", "", -1)
			dirEntry := Dir{fld, 0, currLevel + 1, nil}

			if !dirExists(dr, fld) {
				dr = append(dr, dirEntry)
			}
		}
	}
	return dr
}

// Helper function to check if a directory already exists in the list
func dirExists(dirs []Dir, dirName string) bool {
	for _, dir := range dirs {
		if dir.Name == dirName {
			return true
		}
	}
	return false
}

// Function to find a directory by name and apply a given function to update it
func Update_Dir(dirs []Dir, dirName string, updateFn func(*Dir)) []Dir {
	for i, dir := range dirs {
		if dir.Name == dirName {
			updateFn(&dirs[i]) // Apply the update function to the directory
			break
		}
	}
	return dirs
}

// Function to add a specified size to a directory
func addSize(size int) func(*Dir) {
	return func(dir *Dir) {
		dir.Size += size
	}
}

// Function to add new files to a directory
func addFiles(files []File) func(*Dir) {
	return func(dir *Dir) {
		dir.File = append(dir.File, files...)
	}
}

// Function to gather files for Dir
func Gather_Files(path string, prefix_path string) []File {
	var fl []File

	fullPath := filepath.Join(prefix_path, path)

	c, err := os.ReadDir(fullPath)
	check(err)

	for _, entry := range c {
		if !entry.IsDir() {
			fullPath = filepath.Join(prefix_path+"\\"+path, entry.Name())
			size := getFileSize(fullPath)
			fl = append(fl, File{entry.Name(), size})
		}
	}
	return fl
}

// Update Dir based on gathered files
func Update_Dirs(dirs []Dir, prefixPath string) []Dir {
	var size int
	for _, dir := range dirs {
		size = 0
		files := Gather_Files(dir.Name, prefixPath) // Gather files from the directory

		Update_Dir(dirs, dir.Name, addFiles(files)) //
		size = get_Dir_Size(dirs, dir.Name)
		Update_Dir(dirs, dir.Name, addSize(size))
	}
	return dirs
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

func get_Dir_Size(dirs []Dir, dirName string) int {
	var size int
	for _, dir := range dirs {
		if dir.Name == dirName {
			for i, _ := range dir.File {
				size += dir.File[i].Size
			}
		}
	}
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
