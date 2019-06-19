package nets

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/fogleman/gg"
)

type Config struct {
	Size               int
	Name               string
	StartCornerPercent float64
	Seed               int64
	Oversize           float64

	Background color.NRGBA
	Color1     color.NRGBA
	Color2     color.NRGBA
	Color3     color.NRGBA
	LineAlpha  uint8

	Debug    bool
	LongName bool
	FromEdge bool
}

func Draw(c *Config) {

	var size = c.Size
	sizef := float64(size)
	var targetColor = c.Color2
	var white = c.Color1
	var black = c.Color3

	seed := rand.NewSource(c.Seed)
	r := rand.New(seed)
	dc := gg.NewContext(size, size)
	dc.ScaleAbout(c.Oversize, c.Oversize, sizef/2, sizef/2)

	dc.SetColor(c.Background)
	dc.DrawRectangle(0, 0, sizef, sizef)
	dc.Fill()

	xMin := 0.0
	xMax := sizef + (sizef / 8)
	yMin := 0 - (sizef / 8)
	yMax := sizef

	for i := 0; i < size*400; i++ {
		var x1, y1 float64
		if c.FromEdge {
			var offset = r.Float64() * sizef * c.StartCornerPercent
			if r.Float64() > .5 {
				x1 = offset
			} else {
				y1 = offset
			}
		} else {
			x1 = r.Float64() * sizef * c.StartCornerPercent
			y1 = r.Float64() * sizef * c.StartCornerPercent
		}

		dc.MoveTo(x1, (sizef - y1))

		var xScale = r.Float64()
		var yScale = r.Float64()

		var x2 = xMax - (xScale*(xMax-xMin) + xMin)
		var y2 = yScale*(yMax-yMin) + yMin

		var lineColor color.NRGBA
		if xScale > yScale {
			// set lightness between white and the target color
			lineColor = lerp(targetColor, white, xScale)
		} else {
			// set lightness between the target color and black
			lineColor = lerp(targetColor, black, yScale)
		}
		dc.SetColor(lineColor)
		dc.LineTo(x2, y2)
		dc.Stroke()

		if c.Debug {
			dc.SetColor(color.NRGBA{255, 0, 0, 255})
			dc.DrawRectangle(x1, (sizef - y1), 1, 1)
			dc.Fill()
			dc.SetColor(color.NRGBA{0, 255, 0, 255})
			dc.DrawRectangle(x2, y2, 1, 1)
			dc.Fill()
		}

	}
	var fileName string
	if c.LongName {
		fileName = longName(c.Name, size, c.Seed,
			c.Color1, c.Color2, c.Color3)
	} else {
		fileName = fmt.Sprintf("%s.png", c.Name)
	}
	err := dc.SavePNG(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println("Completed", fileName)
}

func lerp(c1, c2 color.NRGBA, f float64) color.NRGBA {
	if f >= 1.0 {
		return c2
	} else if f <= 0.0 {
		return c1
	}

	c1R, c1G, c1B, c1A := asFloats(c1)
	c2R, c2G, c2B, c2A := asFloats(c2)

	return color.NRGBA{
		R: uint8(c1R + (c2R-c1R)*f),
		G: uint8(c1G + (c2G-c1G)*f),
		B: uint8(c1B + (c2B-c1B)*f),
		A: uint8(c1A + (c2A-c1A)*f),
	}
}

func asFloats(c color.NRGBA) (r, g, b, a float64) {
	r = float64(c.R)
	g = float64(c.G)
	b = float64(c.B)
	a = float64(c.A)
	return
}

func asString(c color.NRGBA) string {
	return fmt.Sprintf("r%dg%db%d", c.R, c.G, c.B)
}

func longName(name string, s int, sd int64, c1, c2, c3 color.NRGBA) string {
	return fmt.Sprintf("%s-s%d-c1%s-c2%s-c3%s-sd%d.png",
		name,
		s,
		asString(c1),
		asString(c2),
		asString(c3),
		sd,
	)
}
