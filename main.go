package main

import (
	"bufio"
	cmd "file-organiser/cmd"
	tbl "file-organiser/table"
	"flag"
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
	var path string
	var dir string
	test := false

	flag.StringVar(&path, "d", "dir", "a file path")
	level := flag.Int("l", 1, "The depth level to return")
	num := flag.Int("n", 5, "The top n files to return")

	flag.Parse()
	if !test {
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
	} else {
		dir = "C:/Users/thoma/Downloads"
		dir = "subdir"
		dir = "C:/Users/thoma/Code"
		dir = "C:/Users/thoma/D&D/.obsidian"
		dir = "C:/Users/thoma/D&D"
	}

	dir = strings.Replace(dir, "/", "\\", -1)

	dirt, _, _ := cmd.Gather_Directories(dir, dir, 1)

	dirt = cmd.Filter_Dir(dirt, *level)

	dirt = cmd.Top_N_Files(dirt, *num)

	cmd.Write_to_file(dirt, "final.txt")

	tbl.Table(dirt)

}
