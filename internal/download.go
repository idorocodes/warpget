package internal

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-resty/resty/v2"
	"github.com/vbauerster/mpb/v8"
	"github.com/vbauerster/mpb/v8/decor"
)

func DownloadFile(url string, chunksStr string, outputPath string) error {
	numChunks, _ := strconv.Atoi(chunksStr)
	meta, _,err := FileInfo(url) 
	if err != nil {
		return err
	}

	
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Truncate(meta)

	// 2. Initialize Progress Bar
	p := mpb.New()
	totalBar := p.AddBar(meta,
		mpb.PrependDecorators(decor.Name("Downloading: ")),
		mpb.AppendDecorators(decor.Percentage(decor.WCSyncSpace)),
	)

	// 3. Orchestration
	client := resty.New()
	chunkSize := meta / int64(numChunks)

	for i := 0; i < numChunks; i++ {
		start := int64(i) * chunkSize
		end := start + chunkSize - 1
		if i == numChunks-1 { end = meta - 1 }

		// Spawn a goroutine for each chunk
		go func(start, end int64) {
			rangeHeader := fmt.Sprintf("bytes=%d-%d", start, end)
			
			// Request chunk
			resp, _ := client.R().SetHeader("Range", rangeHeader).Get(url)
			data := resp.Body()

			// Write to file at specific offset
			file.WriteAt(data, start)
			
			// Update progress bar
			totalBar.IncrBy(len(data))
		}(start, end)
	}

	p.Wait()
	fmt.Println("\nDownload Complete!")
	return nil
}