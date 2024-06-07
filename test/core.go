package translateimage_test

import (
	"image"
	_ "image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"testing"

	langcodes "github.com/spywiree/langcodes"
	"github.com/spywiree/translateimage"
)

func translateFile(t *testing.T, path string) {
	path, err := filepath.Abs(path)
	if err != nil {
		t.Fatal(err)
	}

	img, err := translateimage.TranslateFile(
		path, langcodes.DETECT_LANGUAGE, langcodes.ENGLISH,
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

func translateImage(t *testing.T, path string) {
	r, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()

	src, _, err := image.Decode(r)
	if err != nil {
		t.Fatal(err)
	}

	img, err := translateimage.TranslateImage(
		src, langcodes.DETECT_LANGUAGE, langcodes.ENGLISH,
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

func translateReader(t *testing.T, path string) {
	r, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()

	img, err := translateimage.TranslateReader(
		r, langcodes.DETECT_LANGUAGE, langcodes.ENGLISH,
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
