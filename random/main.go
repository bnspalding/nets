package main

import (
	"flag"
	"image/color"
	"math/rand"

	"github.com/bnspalding/img-gen/nets"
)

var backgroundColor = color.NRGBA{0, 0, 0, 255}

const lineAlpha = 20
const startCornerPercent = .5
const oversize = 1.0
const fromEdge = false
const white = 240
const black = 0

func main() {
	var size int
	flag.IntVar(&size, "size", 320, "pixel dimensions of image (square)")

	var count int
	flag.IntVar(&count, "n", 3, "number of files to generate")

	var la int
	flag.IntVar(&la, "line alpha", lineAlpha, "alpha of lines drawn")

	var seedInput int64
	flag.Int64Var(&seedInput, "seed", 1, "seed")

	flag.Parse()
	r := rand.New(rand.NewSource(seedInput))

	for i := 0; i < count; i++ {
		config := nets.Config{
			Size:               size,
			Name:               "rand",
			StartCornerPercent: startCornerPercent,
			Seed:               seedInput,
			Oversize:           oversize,

			Background: backgroundColor,
			Color1:     color.NRGBA{uint8(r.Intn(255)), uint8(r.Intn(255)), uint8(r.Intn(255)), uint8(la)},
			Color2:     color.NRGBA{uint8(r.Intn(255)), uint8(r.Intn(255)), uint8(r.Intn(255)), uint8(la)},
			Color3:     color.NRGBA{uint8(r.Intn(255)), uint8(r.Intn(255)), uint8(r.Intn(255)), uint8(la)},

			Debug:    false,
			LongName: true,
			FromEdge: false,
		}

		nets.Draw(&config)
	}

}
