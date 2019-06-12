package client

import (
	"bytes"
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/checksum"
	"github/mickaelvieira/taipan/internal/domain/url"
	"io/ioutil"
	"net/http"
)

// Client bot
type Client struct{}

// Head fetches the document and returns the result of the request
// if there is no result, that means:
// - we could not build the request
// - a network error occured
// - we could not read the body
func (f *Client) Head(URL *url.URL) (result *Result, err error) {
	fmt.Printf("Preforming HEAD request %s\n", URL)
	var req *http.Request
	var resp *http.Response

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
	var req *http.Request
	var resp *http.Response

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
