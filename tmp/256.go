package main

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

type kolorki struct {
	r uint8
	g uint8
	b uint8
}

func main() {

	mapa := make(map[uint8]kolorki, 6)

	termbox.Init()
	termbox.SetOutputMode(2)
	defer termbox.Close()
	x, y := termbox.Size()
	termbox.Clear(0, 0)
	var i, c uint8

	for xx := 0; xx < x; xx++ {
		for yy := 0; yy < y; yy++ {

			var g uint8 = i * 6  //zielony
			var r uint8 = i * 36 //czerwony
			var b uint8 = i      //niebieski
			mapa[i] = kolorki{r, g, b}
			r = r

			i++
			if i > 5 {
				i = 0

			}

			g = g
			b = b
			c = b + 17
			termbox.SetCell(xx, yy, '\u2584', termbox.Attribute(17+r+g), termbox.Attribute(c))

		}

	}

	termbox.Flush()
	character := termbox.PollEvent().Key
	character = character
	termbox.Close()
	var t uint8
	for t = 0; t <= 5; t++ {
		fmt.Println(t, mapa[t])
	}
}
