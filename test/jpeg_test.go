package translateimage_test

import (
	"testing"

	"github.com/spywiree/translateimage"
)

func TestTranslateJpegFile(t *testing.T) {
	translateimage.Debug.VideoPath = "dist/jpeg/file"
	translateFile(t, "image.jpeg")
}

func TestTranslateJpegImage(t *testing.T) {
	translateimage.Debug.VideoPath = "dist/jpeg/image"
	translateImage(t, "image.jpeg")
}

func TestTranslateJpegReader(t *testing.T) {
	translateimage.Debug.VideoPath = "dist/jpeg/reader"
	translateReader(t, "image.jpeg")
}
