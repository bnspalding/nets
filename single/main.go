package main

import (
	"flag"
	"image/color"

	"github.com/bnspalding/img-gen/nets"
)

var backgroundColor = color.NRGBA{128, 128, 128, 0}

const lineAlpha = 20
const startCornerPercent = .5
const oversize = 1.0
const fromEdge = false
const white = 240
const black = 0

func main() {
	var size int
	flag.IntVar(&size, "size", 320, "pixel dimensions of image (square)")

	var out string
	flag.StringVar(&out, "out", "out", "name of image file to save to")

	var debug, longName bool
	flag.BoolVar(&debug, "debug", false, "show debug info")
	flag.BoolVar(&longName, "longName", true, "add additional info to file name")

	var c1r, c1g, c1b int
	flag.IntVar(&c1r, "c1r", white, "color1 R value")
	flag.IntVar(&c1g, "c1g", white, "color1 G value")
	flag.IntVar(&c1b, "c1b", white, "color1 B value")

	var c2r, c2g, c2b int
	flag.IntVar(&c2r, "c2r", 128, "color2 R value")
	flag.IntVar(&c2g, "c2g", 128, "color2 G value")
	flag.IntVar(&c2b, "c2b", 128, "color2 B value")

	var c3r, c3g, c3b int
	flag.IntVar(&c3r, "c3r", black, "color3 R value")
	flag.IntVar(&c3g, "c3g", black, "color3 G value")
	flag.IntVar(&c3b, "c3b", black, "color3 B value")

	var la int
	flag.IntVar(&la, "line alpha", lineAlpha, "alpha of lines drawn")

	var seedInput int64
	flag.Int64Var(&seedInput, "seed", 1, "seed")

	flag.Parse()

	config := nets.Config{
		Size:               size,
		Name:               out,
		StartCornerPercent: startCornerPercent,
		Seed:               seedInput,
		Oversize:           oversize,

		Background: backgroundColor,
		Color1:     color.NRGBA{uint8(c1r), uint8(c1g), uint8(c1b), uint8(la)},
		Color2:     color.NRGBA{uint8(c2r), uint8(c2g), uint8(c2b), uint8(la)},
		Color3:     color.NRGBA{uint8(c3r), uint8(c3g), uint8(c3b), uint8(la)},

		Debug:    debug,
		LongName: longName,
		FromEdge: fromEdge,
	}

	nets.Draw(&config)
}
