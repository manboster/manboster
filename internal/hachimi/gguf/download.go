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
	var startByte int64 = 0

	info, err := os.Stat(savePath)
	if err == nil && !info.IsDir() {
		startByte = info.Size()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create %w", err)
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

	var totalSize int64 = 0
	if resp.ContentLength > 0 {
		totalSize = startByte + resp.ContentLength
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusPartialContent {
		return fmt.Errorf("server returned code: %d %s", resp.StatusCode, resp.Status)
	}

	flags := os.O_CREATE | os.O_WRONLY
	if startByte > 0 && resp.StatusCode == http.StatusPartialContent {
		flags |= os.O_APPEND // support so append
	} else {
		flags |= os.O_TRUNC // not support partial then clean
	}

	file, err := os.OpenFile(savePath, flags, 0666)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			color.Yellow(fmt.Sprintf("[Manboster Hachimi Provider] Failed to close file: %s", err))
		}
	}(file)

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	var currentDownloaded int64 = 0

	done := make(chan struct{})
	defer close(done)

	// open a backend goroutine
	go DownloadNotifyRunner(done, ticker, &currentDownloaded, totalSize)

	// wrap struct
	pw := &progressWriter{
		target:     file,
		downloaded: &currentDownloaded,
	}

	_, err = io.Copy(pw, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write: %w", err)
	}

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
