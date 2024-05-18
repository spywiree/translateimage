package translateimage_test

import (
	"image"
	_ "image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"testing"

	languagecodes "github.com/spywiree/langcodes"
	"github.com/spywiree/translateimage"
)

const InputImage = "image.png"

func TestTranslateFile(t *testing.T) {
	translateimage.Debug.VideoPath = "dist/file"

	path, err := filepath.Abs(InputImage)
	if err != nil {
		t.Fatal(err)
	}

	img, err := translateimage.TranslateFile(
		path, languagecodes.DETECT_LANGUAGE, languagecodes.ENGLISH,
	)
	if err != nil {
		t.Fatal(err)
	}

	w, err := os.Create(translateimage.Debug.VideoPath + "/output.png")
	if err != nil {
		t.Fatal(err)
	}
	defer w.Close()

	err = png.Encode(w, img)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTranslateImage(t *testing.T) {
	translateimage.Debug.VideoPath = "dist/image"

	r, err := os.Open(InputImage)
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()

	src, _, err := image.Decode(r)
	if err != nil {
		t.Fatal(err)
	}

	img, err := translateimage.TranslateImage(
		src, languagecodes.DETECT_LANGUAGE, languagecodes.ENGLISH,
	)
	if err != nil {
		t.Fatal(err)
	}

	w, err := os.Create(translateimage.Debug.VideoPath + "/output.png")
	if err != nil {
		t.Fatal(err)
	}
	defer w.Close()

	err = png.Encode(w, img)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTranslateReader(t *testing.T) {
	translateimage.Debug.VideoPath = "dist/reader"

	r, err := os.Open(InputImage)
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()

	img, err := translateimage.TranslateReader(
		r, languagecodes.DETECT_LANGUAGE, languagecodes.ENGLISH,
	)
	if err != nil {
		t.Fatal(err)
	}

	w, err := os.Create(translateimage.Debug.VideoPath + "/output.png")
	if err != nil {
		t.Fatal(err)
	}
	defer w.Close()

	err = png.Encode(w, img)
	if err != nil {
		t.Fatal(err)
	}
}

func init() {
	translateimage.Debug.Enabled = true
}
