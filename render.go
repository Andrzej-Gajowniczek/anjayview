package main

import (
	"github.com/nsf/termbox-go"
)

// showPicture func is to move data from currentImgMap into buffered termbox screen
func (ter *termboxRes) showPicture() {
	termbox.Clear(1, 1)
	ter.doubleConsoleresolution()

	var x, y int
	flip := true
	if ter.movable && ((ter.currentImgX > ter.xMax) || (ter.currentImgY > (ter.yMax * 2))) {

		for y = 0; (y < ter.yMax*2) && (y < (ter.currentImgY-ter.startPointY)*2); y++ { //*2 means sigle cursor is splited horizontally for upper and lower block-pixel what doubles Y resolution
			if y%2 == 0 {
				flip = true
			} else {
				flip = false
			}
			for x = 0; (x < ter.xMax) && (x < ter.currentImgX-ter.startPointX); x++ {
				attr := ter.currentImgMap[y+ter.startPointY][x+ter.startPointX]
				//checkker("y:%d, ter.startPointY:%d, x:%d,  ter.startPointX:%d\n ter.xMax:%d, ter.currentImgX:%d, ter.yMax:%d\n",
				//	y, ter.startPointY, x, ter.startPointX, ter.xMax, ter.currentImgX, ter.yMax)
				//log.Println(x, y)
				if flip {
					termbox.SetFg(x, (y / 2), attr)
				} else {
					termbox.SetBg(x, (y / 2), attr)
				}

			}
		}

	} else {

		startPointX := (ter.xMax - ter.currentImgX) / 2
		startPointY := (ter.yMax*2 - ter.currentImgY) / 2

		for y = 0; y < ter.currentImgY; y++ {

			if ter.yMax%2 == 1 {
				ter.active = 1
			} else {
				ter.active = 0
			}

			if y%2 == 0 {
				flip = true
			} else {
				flip = false
			}
			for x = 0; x < ter.currentImgX; x++ {
				attr := ter.currentImgMap[y][x]
				if flip {
					termbox.SetFg(x+startPointX, ((y + startPointY + ter.active) / 2), attr)
				} else {
					termbox.SetBg(x+startPointX, ((y + startPointY + ter.active) / 2), attr)
				}

			}
		}
		//ter.PrintAt(20, 20, "sX:%d,sY:%d", startPointX, startPointY)
		if ter.xMax < ter.currentImgX || ter.yMax < ter.currentImgY/2 {

			for y = 0; y < ter.desiredImgY; y++ {
				if y%2 == 0 {
					flip = true
				} else {
					flip = false
				}
				for x = 0; x < ter.desiredImgX; x++ {
					attr := ter.desiredImgMap[y][x]
					if flip {
						termbox.SetFg(x, ((y) / 2), attr)
					} else {
						termbox.SetBg(x, ((y) / 2), attr)
					}
				}
			}

		}
	}

	termbox.Flush()

}
