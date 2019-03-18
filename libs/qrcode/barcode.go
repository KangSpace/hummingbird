package qrcode

import "image"

// Contains some meta information about a barcode
type Metadata struct {
	// the name of the barcode kind
	CodeKind string
	// contains 1 for 1D barcodes or 2 for 2D barcodes
	Dimensions byte
}

// a rendered and encoded barcode
type BarCode interface {
	image.Image
	// returns some meta information about the barcode
	Metadata() Metadata
	// the data that was encoded in this barcode
	Content() string
}

// Additional interface that some barcodes might implement to provide
// the value of its checksum.
type BarCodeIntCS interface {
	BarCode
	CheckSum() int
}
