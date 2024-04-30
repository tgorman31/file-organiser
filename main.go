package main

import (
	"bufio"
	cmd "file-organiser/cmd"
	tbl "file-organiser/table"
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
	test := false

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a directory path: ")

	dir, err := reader.ReadString('\n')

	check(err)
	dir = strings.TrimSpace(dir) //Handles windows \r\n for newlines

	if test {
		dir = "C:/Users/thoma/Downloads"
		dir = "subdir"
		dir = "C:/Users/thoma/Code"
		dir = "C:/Users/thoma/D&D/.obsidian"
		dir = "C:/Users/thoma/D&D"
	}

	dir = strings.Replace(dir, "/", "\\", -1)

	dirt, _, _ := cmd.Gather_Directories(dir, dir, 1)

	dirt = cmd.Filter_Dir(dirt, 1)

	dirt = cmd.Top_N_Files(dirt, 5)

	cmd.Write_to_file(dirt, "final.txt")

	tbl.Table(dirt)

}
