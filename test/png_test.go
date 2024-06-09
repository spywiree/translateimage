package translateimage_test

import (
	"testing"

	"github.com/spywiree/translateimage"
)

func TestTranslatePngFile(t *testing.T) {
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

	translateFile(t, ctx, "image.png",
		translateimage.Options{
			DebugMode: true,
			VideoPath: "dist/png/file",
		},
	)
}

func TestTranslatePngImage(t *testing.T) {
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

	translateImage(t, ctx, "image.png",
		translateimage.Options{
			DebugMode: true,
			VideoPath: "dist/png/image",
		},
	)
}

func TestTranslatePngReader(t *testing.T) {
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

	translateReader(t, ctx, "image.png",
		translateimage.Options{
			DebugMode: true,
			VideoPath: "dist/png/reader",
		},
	)
}
