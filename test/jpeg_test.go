package translateimage_test

import (
	"testing"

	"github.com/spywiree/translateimage"
)

func TestTranslateJpegFile(t *testing.T) {
	if PARALLEL {
		t.Parallel()
	}
	if !GLOBALCONTEXT {
		var err error
		ctx, err = translateimage.NewContext()
		if err != nil {
			t.Fatal(err)
		}
		defer ctx.Close()
	}

	translateFile(t, ctx, "image.jpeg",
		translateimage.Options{
			DebugMode: true,
			VideoPath: "dist/jpeg/file",
		},
	)
}

func TestTranslateJpegImage(t *testing.T) {
	if PARALLEL {
		t.Parallel()
	}
	if !GLOBALCONTEXT {
		var err error
		ctx, err = translateimage.NewContext()
		if err != nil {
			t.Fatal(err)
		}
		defer ctx.Close()
	}

	translateImage(t, ctx, "image.jpeg",
		translateimage.Options{
			DebugMode: true,
			VideoPath: "dist/jpeg/image",
		},
	)
}

func TestTranslateJpegReader(t *testing.T) {
	if PARALLEL {
		t.Parallel()
	}
	if !GLOBALCONTEXT {
		var err error
		ctx, err = translateimage.NewContext()
		if err != nil {
			t.Fatal(err)
		}
		defer ctx.Close()
	}

	translateReader(t, ctx, "image.jpeg",
		translateimage.Options{
			DebugMode: true,
			VideoPath: "dist/jpeg/reader",
		},
	)
}
