package main

import (
	"fmt"
	"image"
	"os"

	"image/color"
	_ "image/jpeg"
	_ "image/png"

	"github.com/nsf/termbox-go"
)

// rermboxRes is struct that hold all variable and other data type needed for this app to be stored between functions calls ////////////////
type termboxRes struct {
	//original size image to scroll if it's biger than console size
	active        bool
	movable       bool               //if cursors can scroll a picture biger than console resolution
	xMax          int                //img X console resolution
	yMax          int                //img y console resolution
	startPointX   int                //x index where to start to render img on console
	startPointY   int                //y index where to start to render img on console
	maxPointX     int                // if xMax > currentImgX =0 else maxPointX = currentImgX - xMax
	maxPointY     int                // if xMax > currentImgX =0 else maxPointX = currentImgX - xMax
	mode          termbox.OutputMode //rgb, 256, 8, 216, grayscale - these kind of values
	model         color.Model
	imgOrigPtr    *image.Image          //pointer to origina size picture
	information   string                // png, jpg,tiff , bmp or maybe other format detected by te decode function.
	currentImgX   int                   // img x resolution
	currentImgY   int                   // img y resolution
	currentImgMap [][]termbox.Attribute //mapping to tembox readable pixel format.
	/////// image stretched to the console resolution below
	imgScrPtr     *image.Image          //pointe to a picture that's fit to the console resolution
	desiredImgX   int                   // X - horizontal resolution of a resized picture
	desiredImgY   int                   // Y - verical resolution of a resized picture
	desiredImgMap [][]termbox.Attribute //mapping to resized picture - termbox data format, so far termbox.Cell but it can be modified by future Mem optimalizations
	// picture index holds current index of showed picture
	pictureIndex int      // number of picture in filelist to be displayed
	pictureCount int      // total count of pictures in current dir
	filelist     []string // slice of pictures location/filename.type
	dirPath      string   // dir passed as a parameter to the binary once executed of current dir if no parameter has been passed to executable
}

// main /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func main() {

	var konsola termboxRes
	konsola.movable = true

	// iniT() function is a method on konsola struct that fill in resolution max width and hight as well as initialize termbox screen wit RGB mode
	konsola.initT()
	konsola.xMax, konsola.yMax = termbox.Size()

	konsola.dirInit()

	konsola.doubleConsoleresolution()
	//konsola.renderBigScrollableImg()

	konsola.pictureCount = len(konsola.filelist)
	konsola.pictureIndex = 2
	//checkker("index: %d; filename:%s\ntotal pic count:%d\nfile list: %v\n", konsola.pictureIndex, konsola.filelist[konsola.pictureIndex], konsola.pictureCount, konsola.filelist)
	filePtr := konsola.openImg(konsola.filelist[konsola.pictureIndex])
	if filePtr == nil {
		fmt.Printf("panic: %v\n", "konsola.openImg zwraca nil")
		termbox.Close()
		os.Exit(123)
	}
	/*
		if konsola.movable == true {
			konsola.renderBigScrollableImg()
			termbox.Flush()
		}*/
	konsola.keyboardControl()

}

//End of main()/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
