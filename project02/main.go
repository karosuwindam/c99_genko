package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"

	"github.com/otaviokr/go-epaper-lib"
)

var (
	M2in7bw = epaper.Model{Width: 176, Height: 264, StartTransmission: 0x13}
)

func main() {
	epd, err := Setup()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	epd.Init()
	epd.ClearScreen()
	ImagePut(epd, 0, 0, "sample_out.jpg")

}

func Setup() (*epaper.EPaper, error) {

	epd, err := epaper.New(M2in7bw)
	if err != nil {
		return nil, err
	}
	return epd, nil
}

func ImagePut(epd *epaper.EPaper, x, y int, name string) {
	var img image.Image
	var err error
	image_t, _ := os.Open(name)
	img, err = jpeg.Decode(image_t)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	epd.AddLayer(img, x, y, true)
	epd.PrintDisplay()

}
