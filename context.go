package translateimage

import (
	"bytes"
	"errors"
	"image"
	"image/jpeg"
	"io"
	"path/filepath"

	"github.com/playwright-community/playwright-go"
	langcodes "github.com/spywiree/langcodes"
)

type Context struct {
	pw      *playwright.Playwright
	browser playwright.Browser
}

func (ctx *Context) TranslateFile(path string, source, target langcodes.LanguageCode) (*ImageData, error) {
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
	page, err := ctx.browser.NewPage(pageOptions)
	if err != nil {
		return nil, err
	}
	defer page.Close()

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

	return blob, err
}

func (ctx *Context) TranslateImage(img image.Image, source, target langcodes.LanguageCode) (*ImageData, error) {
	buf := new(bytes.Buffer)
	err := jpeg.Encode(buf, img, &jpeg.Options{Quality: 100})
	if err != nil {
		return nil, err
	}

	name, clean, err := tempFileFromReader(buf, "jpeg")
	if err != nil {
		return nil, err
	}
	defer clean() //nolint:errcheck

	return ctx.TranslateFile(name, source, target)
}

func (ctx *Context) TranslateReader(r io.Reader, source, target langcodes.LanguageCode) (*ImageData, error) {
	name, clean, err := tempFileFromReader(r, "png")
	if err != nil {
		return nil, err
	}
	defer clean() //nolint:errcheck

	return ctx.TranslateFile(name, source, target)
}

func (ctx *Context) Close() error {
	err := ctx.browser.Close()
	if err != nil {
		return err
	}

	return ctx.pw.Stop()
}

func NewContext() (*Context, error) {
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

	browser, err := pw.Firefox.Launch()
	if err != nil {
		return nil, err
	}

	return &Context{pw: pw, browser: browser}, nil
}
