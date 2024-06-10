package main

import (
	"image/png"
	"log"
	"os"
	"path/filepath"

	languagecodes "github.com/spywiree/langcodes"
	"github.com/spywiree/translateimage"
)

func TranslateFile() {
	path, err := filepath.Abs("image.png")
	if err != nil {
		log.Fatalln(err)
	}

	img, err := translateimage.TranslateFile(
		path, languagecodes.DETECT_LANGUAGE, languagecodes.ENGLISH,
	)
	if err != nil {
		log.Fatalln(err)
	}

	f, err := os.Create("output1.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		log.Fatalln(err)
	}
}

func TranslateImage() {
	r, err := os.Open("image.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Close()

	src, err := png.Decode(r)
	if err != nil {
		log.Fatalln(err)
	}

	img, err := translateimage.TranslateImage(
		src, languagecodes.DETECT_LANGUAGE, languagecodes.ENGLISH,
	)
	if err != nil {
		log.Fatalln(err)
	}

	w, err := os.Create("output2.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer w.Close()

	err = png.Encode(w, img)
	if err != nil {
		log.Fatalln(err)
	}
}

func TranslateReader() {
	r, err := os.Open("image.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Close()

	img, err := translateimage.TranslateReader(
		r, languagecodes.DETECT_LANGUAGE, languagecodes.ENGLISH,
	)
	if err != nil {
		log.Fatalln(err)
	}

	w, err := os.Create("output3.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer w.Close()

	err = png.Encode(w, img)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	TranslateFile()
	TranslateImage()
	TranslateReader()
}
