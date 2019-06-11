package client

import (
	"bytes"
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/url"
	"net/http"
	"time"
)

func checkRedirection(URL *url.URL, resp *http.Response) (o *url.URL, f *url.URL, r bool) {
	o = URL
	f = URL
	if resp.Request != nil {
		r = o.String() != resp.Request.URL.String()
		if r {
			f = &url.URL{URL: resp.Request.URL}
		}
	}
	return
}

func makeResult(URL *url.URL, req *http.Request, resp *http.Response, reader *bytes.Reader, checksum []byte) *Result {
	originalURL, finalURL, redirected := checkRedirection(URL, resp)
	return &Result{
		Checksum:         checksum,
		WasRedirected:    redirected,
		ContentType:      resp.Header.Get("Content-Type"),
		ReqURI:           originalURL,
		FinalURI:         finalURL,
		ReqMethod:        req.Method,
		ReqHeaders:       fmt.Sprintf("%s", req.Header),
		RespStatusCode:   resp.StatusCode,
		RespReasonPhrase: resp.Status,
		RespHeaders:      fmt.Sprintf("%s", resp.Header),
		CreatedAt:        time.Now(),
		Content:          reader,
	}
}

func makeClient() *http.Client {
	return &http.Client{}
}

func makeRequest(method string, URL *url.URL) (req *http.Request, err error) {
	req, err = http.NewRequest(method, URL.String(), nil)
	if err != nil {
		return
	}

	req.Header.Add("Accept", "text/html,application/xhtml+xml")
	req.Header.Add("Accept-Language", "en-GB,en;q=0.9,en-US;q=0.8")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Pragma", "no-cache")
	// req.Header.Add("User-Agent", os.Getenv("BOT_USER_AGENT"))

	return
}
