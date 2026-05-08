package gguf

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/fatih/color"
)

type progressWriter struct {
	target     io.Writer // real target
	downloaded *int64    // downloaded bytes
}

func (pw *progressWriter) Write(p []byte) (n int, err error) {
	n, err = pw.target.Write(p)
	// atomic add downloaded data
	atomic.AddInt64(pw.downloaded, int64(n))
	return n, err
}
func Download(ctx context.Context, url string, savePath string) error {
	// set temporary filepath
	tempPath := savePath + ".downloading"

	// any files there?
	if info, err := os.Stat(savePath); err == nil && !info.IsDir() {
		color.Yellow("[Manboster Downloader] it already exists, quit downloading...")
		return nil
	}

	// checkpoint
	var startByte int64 = 0
	info, err := os.Stat(tempPath)
	if err == nil && !info.IsDir() {
		startByte = info.Size()
	}

	// make request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	if startByte > 0 {
		req.Header.Set("Range", "bytes="+strconv.FormatInt(startByte, 10)+"-")
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to request: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			color.Yellow(fmt.Sprintf("[Manboster Hachimi Provider] Failed to close body req: %s", err))
		}
	}(resp.Body)

	// check the code
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusPartialContent {
		return fmt.Errorf("server returned code: %d %s", resp.StatusCode, resp.Status)
	}

	// calculate actual size
	var totalSize int64 = 0
	if resp.ContentLength > 0 {
		totalSize = startByte + resp.ContentLength
	}

	// set downloading
	flags := os.O_CREATE | os.O_WRONLY
	if startByte > 0 && resp.StatusCode == http.StatusPartialContent {
		flags |= os.O_APPEND
	} else {
		flags |= os.O_TRUNC
		startByte = 0
	}

	file, err := os.OpenFile(tempPath, flags, 0666)
	if err != nil {
		return fmt.Errorf("failed to open temp file: %w", err)
	}

	// download status management
	currentDownloaded := startByte
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	done := make(chan struct{})

	// start download notify runner
	go DownloadNotifyRunner(done, ticker, &currentDownloaded, totalSize)

	// wrapper
	pw := &progressWriter{
		target:     file,
		downloaded: &currentDownloaded,
	}

	_, err = io.Copy(pw, resp.Body)

	err = file.Close()
	if err != nil {
		return fmt.Errorf("failed to write data: %w", err)
	}
	close(done)

	err = os.Rename(tempPath, savePath)
	if err != nil {
		return fmt.Errorf("failed to rename temp file: %w", err)
	}

	color.Green(fmt.Sprintf("[Manboster Downloader]: Successfully download! Saving to %s", savePath))
	return nil
}

func DownloadNotifyRunner(done chan struct{}, ticker *time.Ticker, currentDownloaded *int64, totalSize int64) {
	for {
		select {
		case <-ticker.C:
			down := atomic.LoadInt64(currentDownloaded)
			if totalSize > 0 {
				percent := float64(down) / float64(totalSize) * 100
				color.Blue(fmt.Sprintf("\r[Manboster Downloader] Now Downloading: %.2f%%  -  %.2f MB / %.2f MB",
					percent, float64(down)/1024/1024, float64(totalSize)/1024/1024))
			} else {
				color.Blue(fmt.Sprintf("\r[Manboster Downloader] Now Downloading: %.2f MB", float64(down)/1024/1024))
			}
		case <-done:
			return
		}
	}
}
