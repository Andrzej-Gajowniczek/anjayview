package main

import (
	"fmt"
	"golang.org/x/term"
	"os"
)

func main() {
	fd := int(os.Stdout.Fd())
	x,y, err := term.GetSize(fd)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Terminal size: %d columns x %d rows\n", x, y)
}

