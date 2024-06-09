package translateimage_test

import (
	"image"
	_ "image/jpeg"
	"log"
	"os"
	"path/filepath"
	"testing"

	langcodes "github.com/spywiree/langcodes"
	"github.com/spywiree/translateimage"
)

func translateFile(t *testing.T, ctx *translateimage.Context, path string, options translateimage.Options) {
	path, err := filepath.Abs(path)
	if err != nil {
		t.Fatal(err)
	}

	img, err := ctx.TranslateFile(
		path, langcodes.DETECT_LANGUAGE, langcodes.ENGLISH, options,
	)
	if err != nil {
		t.Fatal(err)
	}

	w, err := os.Create(options.VideoPath + "/output.png")
	if err != nil {
		t.Fatal(err)
	}
	defer w.Close()

	data, err := img.ConvertTo("image/png")
	if err != nil {
		t.Fatal(err)
	}

	_, err = w.Write(data)
	if err != nil {
		t.Fatal(err)
	}
}

func translateImage(t *testing.T, ctx *translateimage.Context, path string, options translateimage.Options) {
	r, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()

	src, _, err := image.Decode(r)
	if err != nil {
		t.Fatal(err)
	}

	img, err := ctx.TranslateImage(
		src, langcodes.DETECT_LANGUAGE, langcodes.ENGLISH, options,
	)
	if err != nil {
		t.Fatal(err)
	}

	w, err := os.Create(options.VideoPath + "/output.png")
	if err != nil {
		t.Fatal(err)
	}
	defer w.Close()

	data, err := img.ConvertTo("image/png")
	if err != nil {
		t.Fatal(err)
	}

	_, err = w.Write(data)
	if err != nil {
		t.Fatal(err)
	}
}

func translateReader(t *testing.T, ctx *translateimage.Context, path string, options translateimage.Options) {
	r, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()

	img, err := ctx.TranslateReader(
		r, langcodes.DETECT_LANGUAGE, langcodes.ENGLISH, options,
	)
	if err != nil {
		t.Fatal(err)
	}

	w, err := os.Create(options.VideoPath + "/output.png")
	if err != nil {
		t.Fatal(err)
	}
	defer w.Close()

	data, err := img.ConvertTo("image/png")
	if err != nil {
		t.Fatal(err)
	}

	_, err = w.Write(data)
	if err != nil {
		t.Fatal(err)
	}
}

const PARALLEL bool = true
const GLOBALCONTEXT bool = true

var ctx *translateimage.Context

func TestMain(t *testing.M) {
	if GLOBALCONTEXT {
		var err error
		ctx, err = translateimage.NewContext()
		if err != nil {
			log.Fatalln(err)
		}
		defer ctx.Close()
	}

	os.Exit(t.Run())
}
