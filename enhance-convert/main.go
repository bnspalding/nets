package main

import (
	"flag"
	"image/color"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/bnspalding/img-gen/nets"
)

var backgroundColor = color.NRGBA{0, 0, 0, 255}

const lineAlpha = 20
const startCornerPercent = .5
const oversize = 1.0
const fromEdge = false
const white = 240
const black = 0

type colorString struct {
	r string
	g string
	b string
}

func main() {
	var size int
	flag.IntVar(&size, "size", 320, "pixel dimensions of image (square)")

	flag.Parse()

	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	files, err := ioutil.ReadDir(pwd)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		if !strings.HasSuffix(f.Name(), ".png") {
			continue
		}
		if !isBigger(size, f.Name()) {
			continue
		}
		colors, seedInput := parse(
			strings.TrimSuffix(f.Name(), ".png"))
		config := nets.Config{
			Size:               size,
			Name:               "nets",
			StartCornerPercent: .5,
			Seed:               seedInput,
			Oversize:           oversize,

			Background: backgroundColor,
			Color1:     colors[0],
			Color2:     colors[1],
			Color3:     colors[2],

			Debug:    false,
			LongName: true,
			FromEdge: fromEdge,
		}
		nets.Draw(&config)
		err = os.Remove(f.Name())
		if err != nil {
			panic(err)
		}
	}

}

func parse(f string) ([]color.NRGBA, int64) {
	colors := make([]color.NRGBA, 0)
	var sd int64
	segments := strings.Split(f, "-")
	for _, segment := range segments {
		if segment[0] == 'c' {
			c := parseColor(segment)
			colors = append(colors, c)
		}
		if segment[:2] == "sd" {
			sd = parseSeed(segment)
		}
	}
	return colors, sd
}

func parseColor(s string) color.NRGBA {
	rI := strings.Index(s, "r")
	gI := strings.Index(s, "g")
	bI := strings.Index(s, "b")
	r, err := strconv.ParseUint(s[rI+1:gI], 10, 8)
	if err != nil {
		panic(err)
	}
	g, err := strconv.ParseUint(s[gI+1:bI], 10, 8)
	if err != nil {
		panic(err)
	}
	b, err := strconv.ParseUint(s[bI+1:], 10, 8)
	if err != nil {
		panic(err)
	}
	return color.NRGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: 20,
	}
}

func parseSeed(s string) int64 {
	sd, err := strconv.ParseInt(s[2:], 10, 64)
	if err != nil {
		panic(err)
	}
	return sd
}

func isBigger(size int, fname string) bool {
	fsize := getSize(fname)
	return size > fsize
}

func getSize(fname string) int {
	segments := strings.Split(fname, "-")
	for _, seg := range segments {
		if seg[0] != 's' {
			continue
		}
		if !unicode.IsDigit(rune(seg[1])) {
			continue
		}
		size, err := strconv.Atoi(seg[1:])
		if err != nil {
			panic(err)
		}
		return size
	}

	panic("no size parsed from file name")
	return 0
}
