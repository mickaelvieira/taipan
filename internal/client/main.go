package client

import (
	"bytes"
	"github/mickaelvieira/taipan/internal/domain/checksum"
	"github/mickaelvieira/taipan/internal/domain/url"
	"io"
	"io/ioutil"
	"net/http"
)

// Client bot
type Client struct{}

// Fetch fetches the document
func (f *Client) Fetch(URL *url.URL) (*Result, io.Reader, error) {
	client := makeClient()
	req, err := makeRequest(URL)
	if err != nil {
		return nil, nil, err
	}

	var resp *http.Response
	resp, err = client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	var content []byte
	content, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	reader := bytes.NewReader(content)
	cs := checksum.FromBytes(content)
	result := makeResult(URL, req, resp, cs)

	// Modify URL with the final URL
	*URL = url.URL{URL: &*resp.Request.URL}

	return result, reader, nil
}
