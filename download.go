package translateimage

import (
	_ "embed"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/playwright-community/playwright-go"
)

func mustStringify(s string) string {
	data, err := json.Marshal(s)
	if err != nil {
		panic(
			fmt.Sprintf(`mustStringify(%s): %s`,
				strconv.Quote(s), err.Error(),
			),
		)
	}
	return string(data)
}

func execute(page playwright.Page, code, mainCall string) (any, error) {
	return page.Evaluate(code + "\n" + mainCall)
}

//go:embed js/download.js
var downloadJs string

func download(page playwright.Page, url string) (*ImageData, error) {
	v, err := execute(page, downloadJs,
		fmt.Sprintf(`download(%s)`, mustStringify(url)),
	)
	if err != nil {
		return nil, err
	}

	m := v.(map[string]any)

	data, err := base64.StdEncoding.DecodeString(m["b64data"].(string))
	if err != nil {
		return nil, err
	}

	return &ImageData{
		mimeType: m["contentType"].(string),
		data:     data,
	}, err
}
