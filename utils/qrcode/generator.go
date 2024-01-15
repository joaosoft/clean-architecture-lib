package qrcode

import (
	code "github.com/skip2/go-qrcode"
	"image"
	"io"
)

type IQrCode interface {
	Generate() ([]byte, error)
	Write(io.Writer) error
	Image() image.Image
}

type QrCode struct {
	*code.QRCode
	size int
}

func NewQRCode(content string, size int, removeBorder bool) IQrCode {
	// Initialize QRCode
	qrCode, _ := code.New(content, code.Medium)
	qrCode.DisableBorder = removeBorder

	return &QrCode{
		QRCode: qrCode,
		size:   size,
	}
}

func (q *QrCode) Generate() ([]byte, error) {
	return q.QRCode.PNG(q.size)
}

func (q *QrCode) Write(to io.Writer) error {
	return q.QRCode.Write(q.size, to)
}

func (q *QrCode) Image() image.Image {
	return q.QRCode.Image(q.size)
}
