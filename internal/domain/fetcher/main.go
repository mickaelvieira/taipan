package fetcher

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

// RequestLog represents an entry in the history logs
type RequestLog struct {
	ID               string
	Checksum         []byte
	ContentType      string
	ReqURI           string
	ReqMethod        string
	ReqHeaders       string
	RespStatusCode   int
	RespReasonPhrase string
	RespHeaders      string
	CreatedAt        time.Time
}

// ChecksumToString returns a human readable version of the checksum
func (l *RequestLog) ChecksumToString() string {
	return hex.EncodeToString(l.Checksum)
}

// SetChecksumFromString set the checksum from an hex string
func (l *RequestLog) SetChecksumFromString(checksum string) {
	b, err := hex.DecodeString(checksum)
	if err == nil {
		l.Checksum = b
	}
}

// Fetcher bot
type Fetcher struct {
	Log    *RequestLog
	Reader *bytes.Reader
	client *http.Client
}

// Fetch fetches the document
func (f *Fetcher) Fetch(URL *url.URL) error {
	f.reset()

	client := f.getClient()
	req := f.getRequest(URL)
	resp, err := client.Do(req)

	if err == nil {
		reqLog := &RequestLog{
			ContentType:      resp.Header.Get("Content-Type"),
			ReqURI:           URL.String(),
			ReqMethod:        req.Method,
			ReqHeaders:       fmt.Sprintf("%s", req.Header),
			RespStatusCode:   resp.StatusCode,
			RespReasonPhrase: resp.Status,
			RespHeaders:      fmt.Sprintf("%s", resp.Header),
			CreatedAt:        time.Now(),
		}

		if resp.StatusCode == 200 {
			var b []byte
			b, err = ioutil.ReadAll(resp.Body)
			if err == nil {
				f.Reader = bytes.NewReader(b)
				buf := sha256.New()
				buf.Write(b)
				reqLog.Checksum = buf.Sum(nil)
			}
		}

		f.Log = reqLog
	}
	resp.Body.Close()

	return err
}

func (f *Fetcher) reset() {
	f.Log = nil
	f.Reader = nil
}

func (f *Fetcher) getClient() *http.Client {
	if f.client == nil {
		f.client = &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
	}
	return f.client
}

func (f *Fetcher) getRequest(URL *url.URL) *http.Request {
	req, err := http.NewRequest("GET", URL.String(), nil)

	if err != nil {
		log.Fatal("unable to create the request")
	}

	req.Header.Add("Accept", "text/html,application/xhtml+xml")
	req.Header.Add("Accept-Language", "en-GB,en;q=0.9,en-US;q=0.8")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("User-Agent", os.Getenv("BOT_USER_AGENT"))

	return req
}
