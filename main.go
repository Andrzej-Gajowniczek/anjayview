package main

import (
	"fmt"
	"os"
)

func main() {
	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Current Directory:", currentDir)
}
