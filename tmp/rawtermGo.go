package main

import (
	"fmt"
)

func printColoredText(x, y int, foregroundColor, backgroundColor int, text string) {
	// ANSI escape codes for color formatting
	fgColorCode := 30 + foregroundColor
	bgColorCode := 40 + backgroundColor
	resetCode := 0 // Reset to default colors

	// Position the cursor at the specified coordinates
	positionCode := fmt.Sprintf("\033[%d;%dH", y, x)

	// Set foreground and background colors
	colorCode := fmt.Sprintf("\033[%d;%dm", fgColorCode, bgColorCode)

	// Reset colors after the text
	reset := fmt.Sprintf("\033[%dm", resetCode)

	// Print the formatted text
	fmt.Print(positionCode, colorCode, text, reset)
}

func main() {
	fmt.Print("\033[2J") // Clear the terminal
	fmt.Print("\033[H")  // Move the cursor to the top-left corner

	// Print colored text at coordinates (10, 5)
	printColoredText(20, 10, 3, 5, "Hello, colored world!")

	// Move the cursor to a new line
	fmt.Println()
}
