package main

import (
	"os"

	"github.com/nsf/termbox-go"
)

func (ter *termboxRes) keyboardControl() {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowUp:
				//fmt.Println("Up arrow key pressed")

				if ter.movable {
					ter.startPointY--
					if ter.startPointY < 0 {
						ter.startPointY = 0
					}
					ter.showPicture()
				} else {

				}
				termbox.Flush()
			case termbox.KeyArrowDown:

				if ter.movable {
					//fmt.Println("Down arrow key pressed")
					ter.startPointY++
					if ter.startPointY > (ter.currentImgY - ter.desiredImgY) { //*2 means count with double density because block pixel is a half as tall as whole cursor hight
						ter.startPointY = ter.startPointY - 1
					}
					ter.showPicture()
					//ter.PrintAt(0, 0, "currentY:%d, statrtY:%d, ter.yMax:%d", ter.currentImgY, ter.startPointY, ter.desiredImgY)
				} else {

				}

				//ter.PrintAt(0, 30, "sx:%d, sy:%d  ter.currentImgY - ter.yMax:%d, imgY:%d", ter.startPointX, ter.startPointY, (ter.currentImgY - ter.yMax), ter.currentImgY)
				termbox.Flush()
			case termbox.KeyArrowLeft:
				if ter.movable {
					//fmt.Println("Left arrow key pressed")
					ter.startPointX--
					if ter.startPointX < 0 {
						ter.startPointX = 0
					}
					ter.showPicture()
				} else {

				}

				termbox.Flush()
			case termbox.KeyArrowRight:
				if ter.movable {
					//fmt.Println("Right arrow key pressed")

					ter.startPointX++
					if ter.startPointX > (ter.currentImgX - ter.xMax) {
						ter.startPointX = ter.currentImgX - (ter.xMax)
					}
					ter.showPicture()
				} else {

				}

				termbox.Flush()
			case termbox.KeyCtrlC | termbox.KeyEsc:

				termbox.Close()
				os.Exit(0)
				//	break //eventLoop // Exit the event loop when Ctrl+C is pressed

			//default:
			//	ter.showPicture()
			//	termbox.Flush()
			case termbox.KeySpace:
				ter.pictureIndex = ter.pictureIndex + 1
				//ter.showPicture()
				if ter.pictureIndex >= ter.pictureCount {
					ter.pictureIndex = 0
				}
				ter.openImg(ter.filelist[ter.pictureIndex])
				ter.startPointX = 0
				ter.startPointY = 0
				ter.showPicture()
				//default:
			}
			switch ev.Ch {
			case 'q':
				termbox.Close()
				os.Exit(0)

			case 'f':

				ter.movable = !ter.movable
				ter.openImg(ter.filelist[ter.pictureIndex])
				ter.startPointX = 0
				ter.startPointY = 0
				ter.showPicture()
				/*
					if ter.movable {
						ter.movable = false
					} else {
						ter.movable = true
					}*/
				//time.Sleep(200 * time.Millisecond)

			case 'g':

			}
		}
	}

}
