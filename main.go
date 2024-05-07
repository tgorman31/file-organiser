package main

import (
	cmd "file-organiser/cmd"
	tbl "file-organiser/table"
)

// func check(e error) {
// 	if e != nil {
// 		panic(e)
// 	}
// }

func main() {

	dir, level, num := cmd.User_Input()

	if !cmd.Is_Dir(dir) {
		panic("Supplied String is not a recognised Directory")
	}

	dirt, _, _ := cmd.Gather_Directories(dir, dir, 1)

	dirt = cmd.Filter_Dir(dirt, *level)

	dirt = cmd.Top_N_Files(dirt, *num)

	cmd.Write_to_file(dirt, "final.txt")

	tbl.Table(dirt)

}
