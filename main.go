package main

import (
	cmd "file-organiser/cmd"
	tbl "file-organiser/table"
)

func main() {
	var dir string

	path, level, num := cmd.Path_Suplied()

	if path == "dir" {
		dir = cmd.User_Input()
	}

	dirt, _, _ := cmd.Gather_Directories(dir, dir, 1)

	dirt = cmd.Filter_Dir(dirt, *level)

	dirt = cmd.Top_N_Files(dirt, *num)

	cmd.Write_to_file(dirt, "final.txt")

	tbl.Table(dirt)

}
