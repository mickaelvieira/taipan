package client

import (
	"bytes"
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/url"
	"net/http"
	"time"
)

func checkRedirection(req *http.Request, resp *http.Response) (r bool, u string) {
	u = req.RequestURI
	if resp.Request != nil {
		r = req.RequestURI != resp.Request.URL.String()
		if r {
			u = resp.Request.URL.String()
		}
	}
	return
}

func makeResult(req *http.Request, resp *http.Response, reader *bytes.Reader, checksum []byte) *Result {
	redirected, finalURI := checkRedirection(req, resp)

	return &Result{
		Checksum:         checksum,
		WasRedirected:    redirected,
		ContentType:      resp.Header.Get("Content-Type"),
		ReqURI:           req.RequestURI,
		FinalURI:         finalURI,
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
	return &http.Client{
		// CheckRedirect: func(req *http.Request, via []*http.Request) error {
		// 	return http.ErrUseLastResponse // @TODO I need to double check this. It does not seem to work
		// },
	}
}

func makeRequest(URL *url.URL) (req *http.Request, err error) {
	req, err = http.NewRequest("GET", URL.String(), nil)
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
