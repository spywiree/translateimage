package translateimage_test

import (
	"testing"

	"github.com/spywiree/translateimage"
)

func TestTranslatePngFile(t *testing.T) {
	translateimage.Debug.VideoPath = "dist/png/file"
	translateFile(t, "image.png")
}

func TestTranslatePngImage(t *testing.T) {
	translateimage.Debug.VideoPath = "dist/png/image"
	translateImage(t, "image.png")
}

func TestTranslatePngReader(t *testing.T) {
	translateimage.Debug.VideoPath = "dist/png/reader"
	translateReader(t, "image.png")
}
