package downloader

import (
	"os"
	"github.com/go-resty/resty/v2"
)

func FileInfo(url string) (string, error) {
	client := resty.New()
	resp, err := client.R().SetHeader("Accept", "*/*").Head(url)

	if err != nil {
		os.Exit(1)
	}

	contentLength := resp.Header().Get("Content-Length")
	return contentLength, nil
}
