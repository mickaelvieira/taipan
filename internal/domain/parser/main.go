package parser

import (
	"errors"
	"fmt"
	"net/http"
)

// FetchAndParse fetch the content from the provided URL
// and returns a document containing the relevant information we need
func FetchAndParse(URL string) (*Document, error) {
	if !IsURLValid(URL) {
		return nil, errors.New("Invalid URL")
	}

	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	document := MustParse(resp.Body)

	return document, nil
}
