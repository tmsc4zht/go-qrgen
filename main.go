package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/skratchdot/open-golang/open"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(-1)
	}

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "err: %v", err)
		os.Exit(-1)
	}

}

func run() error {
	qrCode, err := qr.Encode(os.Args[1], qr.M, qr.Auto)
	if err != nil {
		return err
	}

	qrCode, err = barcode.Scale(qrCode, 200, 200)
	if err != nil {
		return err
	}

	file, err := os.CreateTemp("", "qrcode*.png")
	if err != nil {
		return err
	}

	if err := png.Encode(file, qrCode); err != nil {
		return err
	}

	if err := file.Close(); err != nil {
		return err
	}

	return open.Start(file.Name())

}
