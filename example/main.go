package main

import (
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"

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

func Translate_sync() {
	ctx, err := translateimage.NewContext()
	if err != nil {
		log.Fatalln(err)
	}
	defer ctx.Close()

	var wg sync.WaitGroup
	// imagePaths := []string{} //Your image paths
	imagePaths := []string{
		"/downloadDir/09343436/image1.png",
		"/downloadDir/09343436/image2.png",
		"/downloadDir/09343436/image3.png",
	}

	log.Printf("imagePaths", imagePaths)
	for i, path := range imagePaths {
		wg.Add(1)

		log.Printf("path", path)
		dir, err := filepath.Abs(filepath.Dir("."))
		if err != nil {
			log.Printf("Error:", err)
			return
		}
		abs_path := dir + path
		log.Printf("abs_path", abs_path)

		go func(i int) {
			img, err := ctx.TranslateFile(abs_path,
				languagecodes.DETECT_LANGUAGE,
				languagecodes.ENGLISH)
			if err != nil {
				log.Println(err)
				wg.Done()
				return
			}
			data, err := img.ConvertTo("image/png")
			// log.Printf("data:", data)
			if err != nil {
				log.Println(err)
				wg.Done()
				return
			}
			err = os.WriteFile("outputDir/"+strconv.Itoa(i+1)+".png", data, 0666)
			log.Printf("outputDir/" + strconv.Itoa(i+1) + ".png")
			if err != nil {
				log.Println(err)
			}

			wg.Done()
		}(i)
	}

	wg.Wait()
}

func main() {
	Translate_sync()
	TranslateFile()
	TranslateImage()
	TranslateReader()
}
