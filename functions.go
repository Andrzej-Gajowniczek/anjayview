package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/nsf/termbox-go"
)

func checkker(f string, i ...interface{}) {
	termbox.Close()
	fmt.Printf(f, i...)
	os.Exit(111)
}

// Function to check if a file has a graphic file extension
func isGraphicsFile(filename string) bool {
	extensions := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".svg"}
	ext := strings.ToLower(filepath.Ext(filename))
	for _, validExt := range extensions {
		if ext == validExt {
			return true
		}
	}
	return false
}

// Function to get the creation time of a file
func getFileCreationTime(filename string) (time.Time, error) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return time.Time{}, err
	}
	return fileInfo.ModTime(), nil
}

// double Console resolution() divide every screen block horizontally to upper and lower half a cursor which doubles Y axis resolution. Why ? because cursor is twice tall than it's width.
func (ter *termboxRes) doubleConsoleresolution() {
	var x, y int

	for y = 0; y < ter.yMax; y++ {
		for x = 0; x < ter.xMax; x++ {
			termbox.SetChar(x, y, '\u2580') //upper half a cursor color is foreground and lower half of cursor color is backgroud color
		}

	}

}

func (ter *termboxRes) PrintAt(x, y int, formatter string, args ...interface{}) {
	// Move the cursor to the specified coordinates
	//termbox.SetCursor(x, y)

	// Clear the current line at the specified coordinates
	for i := x; i < ter.xMax&len(args); i++ {
		termbox.SetChar(i, y, ' ')
	}

	// Format the text and print it at the specified location
	formattedText := fmt.Sprintf(formatter, args...)
	for i, r := range formattedText {
		termbox.SetCell(x+i, y, r, termbox.ColorDefault, termbox.ColorDefault)
	}

	// Flush the changes to the screen
	termbox.Flush()
}

// check
func checkError(e error, s string) {

	if e != nil {
		termbox.Close()
		log.Print("Error:", e, "during:", s)
		os.Exit(123)
	}
}
func (t *termboxRes) initT() {
	err := termbox.Init()
	checkError(err, "termbox initialization")

	t.mode = termbox.SetOutputMode(termbox.OutputRGB)
	termbox.HideCursor()
}

func (ter *termboxRes) refresh() {
	ter.doubleConsoleresolution()
	termbox.Clear(0, 0)
	ter.showPicture()
	termbox.Flush()
}
