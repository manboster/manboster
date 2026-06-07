package downloader

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/fatih/color"
)

func DownloadNotifyRunner(done chan struct{}, ticker *time.Ticker, currentDownloaded *int64, totalSize int64) {
	previous := int64(0)
	down := int64(0)
	for {
		select {
		case <-ticker.C:
			previous = atomic.LoadInt64(&down)
			down = atomic.LoadInt64(currentDownloaded)
			if totalSize > 0 {
				percent := float64(down) / float64(totalSize) * 100
				speed := (float64(down) - float64(previous)) / 5 / 1024 / 1024
				color.Blue(fmt.Sprintf("\r[Manboster Downloader] Now Downloading: %.2f%%  -  %.2f MB / %.2f MB, %.2f MB/s",
					percent, float64(down)/1024/1024, float64(totalSize)/1024/1024, speed))
			} else {
				color.Blue(fmt.Sprintf("\r[Manboster Downloader] Now Downloading: %.2f MB", float64(down)/1024/1024))
			}
		case <-done:
			return
		}
	}
}
