package client

import (
	"encoding/hex"
	"time"
)

// Result represents an entry in the history logs
type Result struct {
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
func (l *Result) ChecksumToString() string {
	return hex.EncodeToString(l.Checksum)
}

// SetChecksumFromString set the checksum from an hex string
func (l *Result) SetChecksumFromString(checksum string) {
	b, err := hex.DecodeString(checksum)
	if err == nil {
		l.Checksum = b
	}
}

// IsContentDifferent have we fetched a new document
func (l *Result) IsContentDifferent(prev *Result) bool {
	return prev != nil && prev.ChecksumToString() == l.ChecksumToString()
}
