package main

import (
	"bingbong/item"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
	"strconv"

	"github.com/nfnt/resize"
)

const (
	dim = 128
)

func main() {
	err := os.MkdirAll("/tmp/bingbangs", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	nb := 1
	if len(os.Args) > 1 {
		var err error
		nb, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Printf("Building %d bingbong(s)", nb)

	for i := 0; i < nb; i++ {
		m := image.NewRGBA(image.Rect(0, 0, dim, dim))

		it := item.Random()
		for _, part := range []string{it.Body, it.Eyes, it.Mouth, it.Glasses, it.Hat} {
			drawPart(part, m)
		}

		f, err := os.Create(fmt.Sprintf("/tmp/bingbangs/%d.png", i))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		err = png.Encode(f, m)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("File %q saved", f.Name())
	}
}

func drawPart(part string, m *image.RGBA) {
	f, err := os.Open(fmt.Sprintf("assets/%s.png", part))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	img, err := png.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	img = resize.Resize(dim, dim, img, resize.Bicubic)
	draw.Draw(m, m.Bounds(), img, image.ZP, draw.Over)
}
