package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/nfnt/resize"
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

// renderBigScrollableImg func is to move data from currentImgMap into buffered termbox screen
func (ter *termboxRes) renderBigScrollableImg() {
	var x, y int
	flip := true

	for y = 0; (y < ter.yMax*2) && (y < ter.currentImgY-(ter.startPointY*2)); y++ { //*2 means sigle cursor is splited horizontally for upper and lower block-pixel what doubles Y resolution
		if y%2 == 0 {
			flip = true
		} else {
			flip = false
		}
		for x = 0; (x < ter.xMax) && (x < ter.currentImgX-ter.startPointX); x++ {
			attr := ter.currentImgMap[y+ter.startPointY][x+ter.startPointX]
			//checkker("y:%d, ter.startPointY:%d, x:%d,  ter.startPointX:%d\n ter.xMax:%d, ter.currentImgX:%d, ter.yMax:%d\n",
			//	y, ter.startPointY, x, ter.startPointX, ter.xMax, ter.currentImgX, ter.yMax)
			log.Println(x, y)
			if flip == true {
				termbox.SetFg(x, (y / 2), attr)
			} else {
				termbox.SetBg(x, (y / 2), attr)
			}

		}
	}

}

// openImg method reads data from file taken from s variable which holds passed filename to this function
func (ter *termboxRes) openImg(s string) *os.File {
	reader, err := os.Open(s)
	checkError(err, "opening file")
	defer reader.Close()
	m, struna, err := image.Decode(reader)
	ter.imgOrigPtr = &m
	checkError(err, "decoding img")
	ter.information = struna

	bounds := m.Bounds()
	ter.model = m.ColorModel()

	ter.currentImgX = bounds.Dx()
	ter.currentImgY = bounds.Dy()
	//checkker("wymiary:%dx%d\n", ter.currentImgX, ter.currentImgY)
	//checkker("yMin:%d, yMax:%d, xMin:%d, xMax:%d\n", bounds.Min.Y, bounds.Max.Y, bounds.Min.X, bounds.Max.Y)
	new := make([][]termbox.Attribute, ter.currentImgY)
	//checkker("new:%v\n", cap(new))
	for i, _ := range new {
		new[i] = make([]termbox.Attribute, ter.currentImgX)
	}
	ter.currentImgMap = new

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := m.At(x, y).RGBA()
			// A color's RGBA method returns values in the range [0, 65535].
			// Shifting by 12 reduces this to the range [0, 15].
			a = a >> 8
			r = r >> 8
			g = g >> 8
			b = b >> 8
			ter.currentImgMap[y][x] = termbox.RGBToAttribute(uint8(r), uint8(g), uint8(b))
		}
	}

	//konvert image to fit console
	newImg := resize.Resize(uint(ter.xMax), uint(ter.yMax<<1), m, resize.Bilinear)
	ter.imgScrPtr = &newImg
	ter.desiredImgX = newImg.Bounds().Size().X
	ter.desiredImgY = newImg.Bounds().Size().Y

	//checkker("newX:%d, newY:%d\n", ter.desiredImgX, ter.desiredImgY)

	newSmall := make([][]termbox.Attribute, ter.desiredImgY)
	for i, _ := range newSmall {
		newSmall[i] = make([]termbox.Attribute, ter.desiredImgX)

	}

	return reader
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
