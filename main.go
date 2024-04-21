package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	lg "github.com/charmbracelet/lipgloss"
	tbl "github.com/charmbracelet/lipgloss/table"
)

// var style = lg.NewStyle().
// 	Bold(true).
// 	Foreground(lg.Color("#FAFAFA")).
// 	Background(lg.Color("#7D56F4")).
// 	PaddingTop(2).
// 	PaddingLeft(4).
// 	Width(22)

var HeaderStyle = lg.NewStyle().
	Bold(true).
	Align(lg.Center).
	Foreground(lg.Color("#7D56F4")).
	// Background(lg.Color("#7D56F4")).
	// PaddingTop(2).
	// PaddingLeft(4).
	Width(22)

var EvenRowStyle = lg.NewStyle().
	Foreground(lg.Color("#FAFAFA")).
	Align(lg.Center).
	// Background(lg.Color("#7D56F4")).
	Width(22)

var OddRowStyle = lg.NewStyle().
	Foreground(lg.Color("#FAFAFA")).
	Align(lg.Center).
	// Background(lg.Color("#7D56F4")).
	Width(22)

func getDirItems(dir string, dirLevel, currLevel int) int64 {
	// Takes a strig input of a Directory path and an Int of the level the data wishes to be returned at
	var size int64

	c, err := os.ReadDir(dir)
	check(err)

	// fmt.Println("Directory:", dir)

	for _, entry := range c {
		fullPath := filepath.Join(dir, entry.Name())
		if entry.IsDir() {
			currLevel++
			size += getDirItems(fullPath, dirLevel, currLevel)
			currLevel--
		} else {
			size += getFileSize(fullPath)
		}
	}
	if dirLevel >= currLevel {
		fmt.Println(dir, "size =", convertToMB(size), "bytes")
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

func convertToMB(bytes int64) float64 {
	mbSize := float64(bytes) / 1000000
	return mbSize
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
		dir = "C:/Users/thoma/Code"
		dir = "C:/Users/thoma/D&D"
		dir = "C:/Users/thoma/Downloads"
	}

	t := tbl.New().
		Border(lg.NormalBorder()).
		BorderStyle(lg.NewStyle().Foreground(lg.Color("99"))).
		StyleFunc(func(row, col int) lg.Style {
			switch {
			case row == 0:
				return HeaderStyle
			case row%2 == 0:
				return EvenRowStyle
			default:
				return OddRowStyle
			}
		}).
		Headers("LANGUAGE", "FORMAL", "INFORMAL")

	t.Row("English", "You look absolutely fabulous.", "How's it going?")

	size := getDirItems(dir, 2, 1)
	fmt.Println("Dir:", dir, "\nSize:", convertToMB(size), "MB")

	fmt.Println(t)

}
