package client

import (
	"bytes"
	"github/mickaelvieira/taipan/internal/domain/checksum"
	"github/mickaelvieira/taipan/internal/domain/url"
	"io/ioutil"
	"net/http"
)

// Client bot
type Client struct{}

// Fetch fetches the document and returns the result of the request
// if there is no result, that means:
// - we could not build the request
// - a network error occured
// - we could not read the body
func (f *Client) Fetch(URL *url.URL) (result *Result, err error) {
	var req *http.Request
	var resp *http.Response

	client := makeClient()
	req, err = makeRequest(URL)
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

	result = makeResult(req, resp, reader, checksum)

	// Modify URL with the final URL .ie after all redirects
	// @TODO remove this shyte
	*URL = url.URL{URL: &*resp.Request.URL}

	return
}
