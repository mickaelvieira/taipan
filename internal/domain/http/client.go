package http

import (
	"bytes"
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/checksum"
	"github/mickaelvieira/taipan/internal/domain/url"
	"io/ioutil"
	nethttp "net/http"
	"time"
)

func checkRedirection(URL *url.URL, resp *nethttp.Response) (o *url.URL, f *url.URL) {
	o = URL
	f = URL
	if resp.Request != nil {
		if o.String() != resp.Request.URL.String() {
			f = &url.URL{
				URL: resp.Request.URL,
			}
		}
	}
	return
}

func makeResult(URL *url.URL, req *nethttp.Request, resp *nethttp.Response, reader *bytes.Reader, checksum []byte) *Result {
	originalURL, finalURL := checkRedirection(URL, resp)
	return &Result{
		Checksum:         checksum,
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

func makeClient() *nethttp.Client {
	return &nethttp.Client{}
}

func makeRequest(method string, URL *url.URL) (req *nethttp.Request, err error) {
	req, err = nethttp.NewRequest(method, URL.String(), nil)
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

// Client bot
type Client struct{}

// Head fetches the document and returns the result of the request
// if there is no result, that means:
// - we could not build the request
// - a network error occured
// - we could not read the body
func (f *Client) Head(URL *url.URL) (result *Result, err error) {
	fmt.Printf("Preforming HEAD request %s\n", URL)
	var req *nethttp.Request
	var resp *nethttp.Response

	client := makeClient()
	req, err = makeRequest("HEAD", URL)
	if err != nil {
		return
	}

	resp, err = client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var checksum []byte
	result = makeResult(URL, req, resp, nil, checksum)

	return
}

// Get fetches the document and returns the result of the request
// if there is no result, that means:
// - we could not build the request
// - a network error occured
// - we could not read the body
func (f *Client) Get(URL *url.URL) (result *Result, err error) {
	fmt.Printf("Preforming GET request %s\n", URL)
	var req *nethttp.Request
	var resp *nethttp.Response

	client := makeClient()
	req, err = makeRequest("GET", URL)
	if err != nil {
		return
	}

	resp, err = client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var content []byte
	content, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	reader := bytes.NewReader(content)
	checksum := checksum.FromBytes(content)

	result = makeResult(URL, req, resp, reader, checksum)

	return
}
