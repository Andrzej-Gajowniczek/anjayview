package main

import (
	"github.com/nsf/termbox-go"
)

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
			//log.Println(x, y)
			if flip {
				termbox.SetFg(x, (y / 2), attr)
			} else {
				termbox.SetBg(x, (y / 2), attr)
			}

		}
	}

}
