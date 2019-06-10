package client

import (
	"bytes"
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/url"
	"log"
	"net/http"
	"time"
)

func checkRedirection(req *http.Request, resp *http.Response) (o *url.URL, f *url.URL, r bool) {
	var err error
	o, err = url.FromRawURL(req.RequestURI)
	if err != nil {
		log.Fatal(err)
	}

	// Fianl is the original URL
	// before an redirect happened
	f = o
	if resp.Request != nil {
		r = o.String() != resp.Request.URL.String()
		if r {
			f = &url.URL{URL: resp.Request.URL}
		}
	}
	return
}

func makeResult(req *http.Request, resp *http.Response, reader *bytes.Reader, checksum []byte) *Result {
	originalURL, finalURL, redirected := checkRedirection(req, resp)

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
	return &http.Client{
		// CheckRedirect: func(req *http.Request, via []*http.Request) error {
		// 	return http.ErrUseLastResponse // @TODO I need to double check this. It does not seem to work
		// },
	}
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
