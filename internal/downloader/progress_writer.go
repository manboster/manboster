package downloader

import (
	"io"
	"sync/atomic"
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
