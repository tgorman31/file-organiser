package main

import (
	"bufio"
	cmd "file-organiser/cmd"
	tbl "file-organiser/style"
	"fmt"
	"os"
	"strings"
)

// Step1: Get Dir Items
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var dir string
	test := true

	t := tbl.CreateTable()

	if !test {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter a directory path: ")

		dir, err := reader.ReadString('\n')

		check(err)
		dir = strings.TrimSpace(dir) //Handles windows \r\n for newlines

	} else {
		dir = "subdir"
		dir = "C:/Users/thoma/Code"
		dir = "C:/Users/thoma/Downloads"
		dir = "C:/Users/thoma/D&D"
	}

	dir = strings.Replace(dir, "/", "\\", -1)

	// size := cmd.Get_Dir_Items(dir, dir, 2, 1)
	// t.Row("Total", cmd.Readable_Size(size))
	// fmt.Println("Dir:", dir, "\nSize:", convertToMB(size), "MB")

	// fmt.Println(t)

	sortedDirs := cmd.Get_Sorted_Dir(dir, dir, 2, 1)

	// Output the sorted directories
	fmt.Println("Sorted directories:")
	for _, dir := range sortedDirs {
		t.Row(dir.Name, cmd.Readable_Size(dir.Size))
		// fmt.Printf("%s: %d bytes\n", dir.Name, dir.Size)
	}
	fmt.Println(t)
}
