package main

import (
	"bufio"
	cmd "file-organiser/cmd"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var dir string
	test := true

	// t := tbl.CreateTable()

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

	// size, sortedDirs := cmd.Get_Dir_Items(dir, dir, 2, 1)
	dirt := cmd.Gather_Directories(dir, dir, 2, 1)

	cmd.Update_Dirs(dirt, dir)
	cmd.Write_to_file(dirt, "final.txt")

	// fmt.Println(size)
	// Output the sorted directories
	// fmt.Println("Sorted directories:")
	// for _, dir := range sortedDirs {
	// 	t.Row(dir.Name, cmd.Readable_Size(dir.Size))
	// 	// fmt.Printf("%s: %d bytes\n", dir.Name, dir.Size)
	// }
	// fmt.Println(t)
}
