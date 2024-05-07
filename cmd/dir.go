package cmd

import (
	"bufio"
	"flag"
	"sort"
	"strings"

	// tbl "file-organiser/style"
	"fmt"
	"os"
	"path/filepath"
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
func Gather_Directories(dir, fullPath string, currLevel int) ([]Dir, []File, int) {
	var dr []Dir
	total_size := 0
	files := []File{}

	c, err := os.ReadDir(fullPath)
	check(err)

	for _, entry := range c {
		entryPath := filepath.Join(fullPath, entry.Name())

		if entry.IsDir() {

			sub_dir, dir_files, dir_size := Gather_Directories(dir, entryPath, currLevel+1)

			total_size += dir_size

			dr = append(dr, sub_dir...)

			fld := strings.Replace(entryPath, dir+"\\", "", -1)

			sort_Files(dir_files)

			dirEntry := Dir{fld, dir_size, currLevel, dir_files}

			files = append(files, File{entry.Name(), dir_size, true})

			if !dirExists(dr, fld) {
				dr = append(dr, dirEntry)
				sort_Directories(dr)
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

func Is_Dir(fullpath string) bool {
	_, err := os.ReadDir(fullpath)

	if err != nil {
		return false
	}
	return true
}

func User_Input() (dir string, level *int, num *int) {
	var path string

	flag.StringVar(&path, "d", "dir", "a file path")
	level = flag.Int("l", 1, "The depth level to return")
	num = flag.Int("n", 5, "The top n files to return")

	flag.Parse()
	if path == "dir" {
		fmt.Print("Enter a directory path: ")

		reader := bufio.NewReader(os.Stdin)
		var err error
		dir, err = reader.ReadString('\n')
		check(err)
		dir = strings.TrimSpace(dir) //Handles windows \r\n for newlines

	} else {
		dir = path
	}

	dir = strings.Replace(dir, "/", "\\", -1)

	return dir, level, num

}

func Filter_Dir(dirs []Dir, depth int) []Dir {
	var d []Dir
	for _, dir := range dirs {
		if dir.Depth == depth {
			d = append(d, dir)
		}
	}
	return d
}

func Top_N_Files(dirs []Dir, n int) []Dir {
	var d []Dir
	var f []File
	for _, dir := range dirs {
		for i, file := range dir.File {
			if i == n {
				break
			}
			f = append(f, File{file.Name, file.Size, file.isDir})
		}
		d = append(d, Dir{dir.Name, dir.Size, dir.Depth, f})
		f = nil
	}
	return d
}

func Get_Files_From_Dir(dirs []Dir, dir_name string) []File {
	var f []File
	for _, dir := range dirs {
		if dir.Name == dir_name {
			f = append(f, dir.File...)
		}
	}
	return f
}

func Write_to_file(dir []Dir, fileName string) {
	fl, err := os.Create(fileName)
	check(err)
	str := fmt.Sprintln(dir)
	fl.WriteString(str + "\n")
	fl.Close()
}

func sort_Directories(dirs []Dir) {
	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].Size > dirs[j].Size
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
