package fileutil

import (
	"bytes"
	"image"
	"io"

	auyersteganography "github.com/auyer/steganography"
)

type AuyerPNGSteganographer struct{}

// Encode reads data from the input io.Reader and encodes it in a new image inside the provided buffer
func (s *AuyerPNGSteganographer) Encode(img image.Image, dataToEncode []byte, output *bytes.Buffer) error {
	return auyersteganography.Encode(output, img, dataToEncode)
}

// Decode writes encoded data from the image to the writer
func (s *AuyerPNGSteganographer) Decode(img image.Image, w io.Writer, numBytes uint64) error {
	rawdata := auyersteganography.Decode(uint32(numBytes), img)
	_, err := w.Write(rawdata)
	return err
}
