package cmd

import (
	"cmp"
	"sort"
	"strings"

	// tbl "file-organiser/style"
	"fmt"
	"os"
	"path/filepath"
	"slices"
)

type Dir struct {
	Name  string
	Size  int
	Depth int
	File  []File
}

type File struct {
	Name  string
	Size  int
	isDir bool
}

// Function to gather directories
func Gather_Directories(dir, fullPath string, dirLevel, currLevel int) ([]Dir, []File, int) {
	var dr []Dir
	total_size := 0
	files := []File{}

	c, err := os.ReadDir(fullPath)
	check(err)

	for _, entry := range c {
		entryPath := filepath.Join(fullPath, entry.Name())

		if entry.IsDir() {

			sub_dir, dir_files, dir_size := Gather_Directories(dir, entryPath, dirLevel, currLevel+1)

			total_size += dir_size

			dr = append(dr, sub_dir...)

			fld := strings.Replace(entryPath, dir+"\\", "", -1)

			sort_Files(dir_files)

			dirEntry := Dir{fld, dir_size, currLevel, dir_files}

			files = append(files, File{entry.Name(), dir_size, true})

			if !dirExists(dr, fld) {
				dr = append(dr, dirEntry)
			}
		} else {
			file_size := getFileSize(entryPath)
			total_size += file_size
			files = append(files, File{entry.Name(), file_size, false})
		}
	}
	return dr, files, total_size
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

// // Function to find a directory by name and apply a given function to update it
// func Update_Dir(dirs []Dir, dirName string, updateFn func(*Dir)) []Dir {
// 	for i, dir := range dirs {
// 		if dir.Name == dirName {
// 			updateFn(&dirs[i]) // Apply the update function to the directory
// 			break
// 		}
// 	}
// 	return dirs
// }

// // Function to add a specified size to a directory
// func addSize(size int) func(*Dir) {
// 	return func(dir *Dir) {
// 		dir.Size += size
// 	}
// }

// // Function to add new files to a directory
// func addFiles(files []File) func(*Dir) {
// 	return func(dir *Dir) {
// 		dir.File = append(dir.File, files...)
// 	}
// }

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

	sort.Slice(files, func(i, j int) bool {
		return files[i].Size > files[j].Size
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

// // Adapted from the SortKeys example in the "sort" docs
// // By is the type of a "less" function that defines the ordering of its Planet arguments.
// type By func(f1, f2 *File) bool

// // Sort is a method on the function type, By, that sorts the argument slice according to the function.
// func (by By) Sort(files []File) {
// 	fls := &fileSorter{
// 		files: files,
// 		by:    by, //The Sort method's receiver is the fuction (closure) that defines the sort order
// 	}
// 	sort.Sort(fls)
// }

// // fileSorter joins a By function and a slice of Planets to be sorted.
// type fileSorter struct {
// 	files []File
// 	by    func(f1, f2 *File) bool // Closure use in the Less method
// }

// // Len is part of sort.Interface.
// func (s *fileSorter) Len() int {
// 	return len(s.files)
// }

// // Swap is part of sort.Interface
// func (s *fileSorter) Swap(i, j int) {
// 	s.files[i], s.files[j] = s.files[j], s.files[i]
// }

// // Less is part of sort.Interface. It is implemented by calling th "by" closure in the sorter
// func (s *fileSorter) Less(i, j int) bool {
// 	return s.by(&s.files[i], &s.files[j])
// }
//
// A function using the above
// func sort_Files(files []File) {
//
// 	size := func(f1, f2 *File) bool {
// 		return f1.Size > f2.Size
// 	}
// 	By(size).Sort(files)
//
// 	name := func(f1, f2 *File) bool {
// 	return f1.name < f2.name
// }
// }

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
