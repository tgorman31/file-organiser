package main

import (
	"fmt"
)

func getDirItems(dir string) string {
	message := fmt.Sprintf("Dir is %v", dir)

	return message
}

// Step0: Hello Worls

func main() {
	dir := "/user/code"
	fmt.Println(getDirItems(dir))
}
