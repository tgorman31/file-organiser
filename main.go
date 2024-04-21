package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getDirItems(dir string) {

	c, err := os.ReadDir(dir)
	check(err)

	fmt.Println(dir)
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}
	// return
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Step1: Get Dir Items

func main() {
	var dir string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a directory path: ")

	dir, err := reader.ReadString('\n')

	check(err)
	dir = strings.TrimSpace(dir) //Handles windows \r\n for newlines

	// os.MkdirAll("subdir/parent/child", 0755)

	// createEmptyFile := func(name string) {
	// 	d := []byte("")
	// 	check(os.WriteFile(name, d, 0644))
	// }

	// createEmptyFile("subdir/parent/file2")
	// createEmptyFile("subdir/parent/file3")
	// createEmptyFile("subdir/parent/child/file4")

	// fmt.Println(getDirItems(dir))

	getDirItems(dir)

}
