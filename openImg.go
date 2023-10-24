package main

import (
	"image"
	"os"

	"github.com/nfnt/resize"
	"github.com/nsf/termbox-go"
)

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
