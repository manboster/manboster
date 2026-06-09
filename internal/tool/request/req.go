package request

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Response struct {
	StatusCode int    `json:"status_code"`
	Body       string `json:"body"`
}

// MakeRequest receive RunArgs and execute
func MakeRequest(args RunArgs) (string, error) {
	// get payload
	// if it's empty, it won't report any error
	payloadReader := strings.NewReader(args.Payload)

	// create HTTP request
	req, err := http.NewRequest(args.Verb, args.URL, payloadReader)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// 3. process headers
	if args.Headers != "" {
		var headersMap map[string]string
		err := json.Unmarshal([]byte(args.Headers), &headersMap)
		if err != nil {
			return "", fmt.Errorf("failed to parse header json: %w", err)
		}

		// append headers
		for key, value := range headersMap {
			req.Header.Set(key, value)
		}
	}

	// 4. set timeout
	client := &http.Client{
		Timeout: time.Duration(args.Timeout) * time.Second,
	}

	// 5. make request!
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			color.Yellow("[Manboster Request] failed to close response body")
		}
	}(resp.Body)

	// 6. read and print the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read the response: %w", err)
	}

	res, err := json.Marshal(Response{
		StatusCode: resp.StatusCode,
		Body:       string(body),
	})

	if err != nil {
		return "", fmt.Errorf("failed to marshal response: %w", err)
	}

	return string(res), nil
}
