package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/ejuju/cybersec/pkg/fileutil"
)

func main() {
	outputImagePath := "./output.jpg"
	steganographer := &fileutil.AuyerPNGSteganographer{}

	// get the original image
	width := 1024
	height := 1024
	baseimg := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			baseimg.Set(x, y, color.RGBA{255, 255, 255, 255})
		}
	}

	// encode message in new image (buffer)
	msg := []byte("i'm hidden")
	output := &bytes.Buffer{}
	err := steganographer.Encode(baseimg, msg, output)
	if err != nil {
		panic(err)
	}

	// save new image to file
	f, err := os.Create(outputImagePath)
	if err != nil {
		panic(err)
	}
	outputimg, err := png.Decode(output)
	if err != nil {
		panic(err)
	}
	err = png.Encode(f, outputimg)
	if err != nil {
		panic(err)
	}

	// open output image
	f, err = os.Open(outputImagePath)
	if err != nil {
		panic(err)
	}
	outputimg, err = png.Decode(f)
	if err != nil {
		panic(err)
	}

	// decode hidden data from image
	buf := &bytes.Buffer{}
	err = steganographer.Decode(outputimg, buf, uint64(len(msg)))
	if err != nil {
		panic(err)
	}

	fmt.Println("Found message:", buf.String())
}
