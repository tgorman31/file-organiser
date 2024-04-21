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

	err = os.Chdir(dir)

	fmt.Println(dir)
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
		if !entry.IsDir() {
			fileInfo, err := os.Stat(entry.Name())
			check(err)
			fmt.Println(" ", fileInfo.Size())
		}
	}
	// return
}

// func getDirSize(file string) float64 {
// 	// Takes a file path and calculates its size
// 	size := 1.0

// 	return size
// }

func check(e error) {
	if e != nil {
		panic(e)
	}
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
		dir = "C:/Users/thoma/D&D"
	}

	getDirItems(dir)

}
