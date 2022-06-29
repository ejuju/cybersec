package fileutil

import (
	"bytes"
	"image"
	"io"
)

type PNGSteganographer interface {
	Encode(img image.Image, dataToEncode []byte, output *bytes.Buffer) error
	Decode(img image.Image, w io.Writer, numBytes uint64) error
}
