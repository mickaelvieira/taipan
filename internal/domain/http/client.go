package http

import (
	"bytes"
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/checksum"
	"github/mickaelvieira/taipan/internal/domain/url"
	"github/mickaelvieira/taipan/internal/logger"
	"io/ioutil"
	"log"
	nethttp "net/http"
	"time"
)

func checkRedirection(URL *url.URL, res *nethttp.Response) (o *url.URL, f *url.URL) {
	o = URL
	f = URL
	if res != nil && res.Request != nil {
		if o.String() != res.Request.URL.String() {
			f = &url.URL{
				URL: res.Request.URL,
			}
		}
	}
	return
}

func makeSuccessfulResult(URL *url.URL, req *nethttp.Request, res *nethttp.Response, b []byte) *Result {
	o, f := checkRedirection(URL, res)
	cs := checksum.FromBytes(b)
	// https://www.openmymind.net/Go-Slices-And-The-Case-Of-The-Missing-Memory/
	// b := bytes.NewBuffer(make([]byte, 0, res.ContentLength))
	// b.ReadFrom(res.Body)
	// body := buffer.Bytes()

	return &Result{
		Checksum:         cs,
		ContentType:      res.Header.Get("Content-Type"),
		ReqURI:           o,
		FinalURI:         f,
		ReqMethod:        req.Method,
		ReqHeaders:       fmt.Sprintf("%s", req.Header),
		RespStatusCode:   res.StatusCode,
		RespReasonPhrase: res.Status,
		RespHeaders:      fmt.Sprintf("%s", res.Header),
		CreatedAt:        time.Now(),
		Content:          bytes.NewReader(b),
	}
}

func makeFailedResult(URL *url.URL, req *nethttp.Request, err error) *Result {
	return &Result{
		ReqURI:        URL,
		FinalURI:      URL,
		ReqMethod:     req.Method,
		ReqHeaders:    fmt.Sprintf("%s", req.Header),
		CreatedAt:     time.Now(),
		Failed:        true,
		FailureReason: err.Error(),
	}
}

func mustCreateClient() *nethttp.Client {
	return &nethttp.Client{
		Timeout: 10 * time.Second,
	}
}

func mustCreateRequest(method string, URL *url.URL) *nethttp.Request {
	r, err := nethttp.NewRequest(method, URL.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	r.Header.Add("Accept", "text/html,application/xhtml+xml")
	r.Header.Add("Accept-Language", "en-GB,en;q=0.9,en-US;q=0.8")
	r.Header.Add("Cache-Control", "no-cache")
	r.Header.Add("Pragma", "no-cache")
	// req.Header.Add("User-Agent", os.Getenv("BOT_USER_AGENT"))

	return r
}

// Client bot
type Client struct{}

// Get fetches the document and returns the result of the request
// if there is no result, that means:
// - we could not build the request
// - a network error occured
// - we could not read the body
func (c *Client) Get(URL *url.URL) *Result {
	logger.Info(fmt.Sprintf("Fetching %s", URL.String()))

	client := mustCreateClient()
	req := mustCreateRequest("GET", URL)
	res, err := client.Do(req)
	if err != nil {
		return makeFailedResult(URL, req, err)
	}
	defer res.Body.Close()

	// Let's sure we can read the full response before continuing
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return makeFailedResult(URL, req, err)
	}

	return makeSuccessfulResult(URL, req, res, b)
}
