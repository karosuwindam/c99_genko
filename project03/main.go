package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"os"
	"strings"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

func main() {
	tmp := writedata("1234０００ＴＴTT", 50.0)
	file, err := os.Create(`test.png`)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	file.Write(tmp)
}

func writedata(text string, size float64) []byte {

	// フォントファイルを読み込み
	ftBinary, err := ioutil.ReadFile("/usr/share/fonts/truetype/fonts-japanese-gothic.ttf")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	ft, err := truetype.Parse(ftBinary)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fontsize := size
	opt := truetype.Options{
		Size:              float64(fontsize),
		DPI:               0,
		Hinting:           0,
		GlyphCacheEntries: 0,
		SubPixelsX:        0,
		SubPixelsY:        0,
	}

	slice := strings.Split(text, "")
	imageWidth_t := 0.0
	for _, str := range slice {
		if len(str) == 1 {
			imageWidth_t += fontsize*0.5 + fontsize*0.1
		} else {
			imageWidth_t += fontsize + fontsize*0.05
		}
	}
	imageWidth := int(imageWidth_t)
	imageHeight := int(fontsize)
	textTopMargin := int(fontsize * 0.9)

	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	// draw.Draw(img, img.Bounds(), image.White, image.ZP, draw.Src)

	face := truetype.NewFace(ft, &opt)

	dr := &font.Drawer{
		Dst:  img,
		Src:  image.Black,
		Face: face,
		Dot:  fixed.Point26_6{},
	}

	dr.Dot.X = (fixed.I(imageWidth) - dr.MeasureString(text)) / 2
	dr.Dot.Y = fixed.I(textTopMargin)

	dr.DrawString(text)

	buf := &bytes.Buffer{}
	err = png.Encode(buf, img)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return buf.Bytes()
}
