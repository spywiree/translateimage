package translateimage

import (
	"bytes"
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"

	"github.com/playwright-community/playwright-go"
	languagecodes "github.com/spywiree/langcodes"
)

func tempFileFromReader(r io.Reader) (string, func() error, error) {
	f, err := os.CreateTemp("", "*.jpg")
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
	// Logger.Println("Driver and Firefox browser installed")

	pw, err := playwright.Run()
	if err != nil {
		return nil, err
	}
	defer pw.Stop() //nolint:errcheck
	// Logger.Println("Playwright instance started")

	browser, err := pw.Firefox.Launch()
	if err != nil {
		return nil, err
	}
	defer browser.Close()
	// Logger.Println("Firefox instance started")

	var pageOptions playwright.BrowserNewPageOptions
	if Debug.Enabled {
		videoPath, err := filepath.Abs(Debug.VideoPath)
		if err != nil {
			return nil, err
		}

		pageOptions = playwright.BrowserNewPageOptions{
			RecordVideo: &playwright.RecordVideo{
				Dir: videoPath,
			},
		}
	}
	page, err := browser.NewPage(pageOptions)
	if err != nil {
		return nil, err
	}
	// Logger.Println("Created a new page")

	url := "https://translate.google.pl/?op=images"
	url += "&hl=en"
	url += "&sl=" + string(source)
	url += "&tl=" + string(target)

	_, err = page.Goto(url)
	if err != nil {
		return nil, err
	}
	// Logger.Println("Set the site URL to:", url)

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
	// Logger.Println("The file has been uploaded")

	imgElement := page.Locator(
		`div.CMhTbb:nth-child(2) > img:nth-child(1)`,
	)
	blobUrl, err := imgElement.GetAttribute("src")
	if err != nil {
		return nil, err
	}
	// Logger.Println("Image element has been found")

	blob, err := download(page, blobUrl)
	if err != nil {
		return nil, err
	}
	// Logger.Println("Image downloaded")

	var out image.Image
	switch blob.ContentType {
	case "image/png":
		out, err = png.Decode(bytes.NewBuffer(blob.Data))
	case "image/jpeg":
		out, err = jpeg.Decode(bytes.NewBuffer(blob.Data))
	}
	if err != nil {
		return nil, err
	}
	// Logger.Println("Image decoded")

	return out, err
}

func TranslateImage(img image.Image, source, target languagecodes.LanguageCode) (image.Image, error) {
	buf := new(bytes.Buffer)
	err := jpeg.Encode(buf, img, &jpeg.Options{Quality: 100})
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
