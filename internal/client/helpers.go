package client

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

func makeResult(URL *url.URL, req *http.Request, resp *http.Response, checksum []byte) *Result {
	return &Result{
		Checksum:         checksum,
		ContentType:      resp.Header.Get("Content-Type"),
		ReqURI:           URL.String(),
		ReqMethod:        req.Method,
		ReqHeaders:       fmt.Sprintf("%s", req.Header),
		RespStatusCode:   resp.StatusCode,
		RespReasonPhrase: resp.Status,
		RespHeaders:      fmt.Sprintf("%s", resp.Header),
		CreatedAt:        time.Now(),
	}
}

func makeClient() *http.Client {
	return &http.Client{
		// CheckRedirect: func(req *http.Request, via []*http.Request) error {
		// 	return http.ErrUseLastResponse // @TODO I need to double check this. It does not seem to work
		// },
	}
}

func makeRequest(URL *url.URL) (*http.Request, error) {
	req, err := http.NewRequest("GET", URL.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "text/html,application/xhtml+xml")
	req.Header.Add("Accept-Language", "en-GB,en;q=0.9,en-US;q=0.8")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("User-Agent", os.Getenv("BOT_USER_AGENT"))

	return req, nil
}
