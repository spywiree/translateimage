package translateimage

import (
	"bytes"
	"errors"
	"image"
	"image/png"
	"io"
	"os"

	"github.com/playwright-community/playwright-go"
	languagecodes "github.com/spywiree/langcodes"
)

func tempFileFromReader(r io.Reader) (string, func() error, error) {
	f, err := os.CreateTemp("", "*.png")
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

// Path must be absolute.
// Supported file types: .jpg, .jpeg, .png.
func TranslateFile(path string, source, target languagecodes.LanguageCode) (image.Image, error) {
	err := playwright.Install(
		&playwright.RunOptions{
			Browsers: []string{"firefox"},
		},
	)
	if err != nil {
		return nil, err
	}

	pw, err := playwright.Run()
	if err != nil {
		return nil, err
	}
	defer pw.Stop() //nolint:errcheck

	browser, err := pw.Firefox.Launch()
	if err != nil {
		return nil, err
	}
	defer browser.Close()

	page, err := browser.NewPage()
	if err != nil {
		return nil, err
	}

	url := "https://translate.google.pl/?op=images"
	url += "&hl=en"
	url += "&sl=" + string(source)
	url += "&tl=" + string(target)

	_, err = page.Goto(url)
	if err != nil {
		return nil, err
	}

	err = page.GetByRole(
		*playwright.AriaRoleButton,
		playwright.PageGetByRoleOptions{
			Name: "Reject all",
		},
	).Click(
		playwright.LocatorClickOptions{
			Timeout: playwright.Float(5000), //5s
		},
	)
	if err != nil && !errors.Is(err, playwright.ErrTimeout) {
		return nil, err
	}

	err = page.GetByRole(
		*playwright.AriaRoleTextbox,
		playwright.PageGetByRoleOptions{
			Name: "Browse your files",
		},
	).SetInputFiles(path)
	if err != nil {
		return nil, err
	}

	imgElement := page.Locator(
		`div.CMhTbb:nth-child(2) > img:nth-child(1)`,
	)
	blobUrl, err := imgElement.GetAttribute("src")
	if err != nil {
		return nil, err
	}

	blob, err := download(page, blobUrl)
	if err != nil {
		return nil, err
	}

	out, _, err := image.Decode(bytes.NewBuffer(blob.Data))
	if err != nil {
		return nil, err
	}

	return out, err
}

func TranslateImage(img image.Image, source, target languagecodes.LanguageCode) (image.Image, error) {
	buf := new(bytes.Buffer)
	err := png.Encode(buf, img)
	if err != nil {
		return nil, err
	}

	name, clean, err := tempFileFromReader(buf)
	if err != nil {
		return nil, err
	}
	defer clean() //nolint:errcheck

	return TranslateFile(name, source, target)
}

func TranslateReader(r io.Reader, source, target languagecodes.LanguageCode) (image.Image, error) {
	name, clean, err := tempFileFromReader(r)
	if err != nil {
		return nil, err
	}
	defer clean() //nolint:errcheck

	return TranslateFile(name, source, target)
}
