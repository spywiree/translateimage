package translateimage

import (
	"io"
	"os"
)

func tempFileFromReader(r io.Reader, ext string) (string, func() error, error) {
	f, err := os.CreateTemp("", "*."+ext)
	if err != nil {
		return "", nil, err
	}
	defer f.Close()

	_, err = io.Copy(f, r)
	if err != nil {
		return "", nil, err
	}

	clean := func() error {
		return os.Remove(f.Name())
	}

	return f.Name(), clean, nil
}
