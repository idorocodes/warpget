package internal

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/go-resty/resty/v2"
)
 
func FileInfo(url string) (int64, bool, error) {
	client := resty.New()
	
	// Perform the HEAD request
	resp, err := client.R().SetHeader("Accept", "*/*").Head(url)
	if err != nil {
		return 0, false, err
	}

	// Get Content-Length
	contentLength := resp.Header().Get("Content-Length")
	lengthInt, err := strconv.ParseInt(contentLength, 10, 64)
	if err != nil {
		return 0, false, fmt.Errorf("invalid content length: %w", err)
	}

	// Check Accept-Ranges header
	acceptRanges := resp.Header().Get("Accept-Ranges") == "bytes"

	if lengthInt <= 0 {
		return 0, false, errors.New("content length unknown")
	}

	fmt.Printf("File accessed successfully. Size: %d bytes, Range Support: %v\n", lengthInt, acceptRanges)

	return lengthInt, acceptRanges, nil
}