package util

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
)

func ReaderToBase64URL(reader io.Reader) (string, error) {
	// 1. read all image into this
	bytes, err := io.ReadAll(reader)
	if err != nil {
		return "", fmt.Errorf("failed to read: %w", err)
	}

	return base64.StdEncoding.EncodeToString(bytes), nil
}

func BytesToBase64URL(bytes []byte) (string, error) {
	// check MIME Type dynamically
	mimeType := http.DetectContentType(bytes)

	// 3. from []byte to base64 string
	base64Str := base64.StdEncoding.EncodeToString(bytes)

	// 4. make it as data uri format
	base64URL := fmt.Sprintf("data:%s;base64,%s", mimeType, base64Str)

	return base64URL, nil
}
