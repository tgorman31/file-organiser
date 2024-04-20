package main

import (
	"bufio"
	"fmt"
	"os"
)

func getDirItems(dir string) string {
	message := fmt.Sprintf("Dir is %v", dir)

	return message
}

// Step1: Get Dir Items

func main() {
	var dir string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a directory path: ")

	dir, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("An error occured while reading input: ", err)
	}

	// dir := "/user/code"

	dir = dir[:len(dir)-1]

	fmt.Println(getDirItems(dir))

}
