package image

import (
	"encoding/base64"
	"io"
	"strings"
)

const prefix = "data:"
const token = ";base64"
const separator = ","

// GetBase64Reader returns an IO reader for the base64 input
func GetBase64Reader(i string) io.Reader {
	return base64.NewDecoder(base64.StdEncoding, strings.NewReader(i))
}

// GetContentType retrieves the content type from the base64 input
func GetContentType(i string) string {
	idx := strings.LastIndex(i, separator)
	if idx == -1 {
		return ""
	}
	s := i[0:idx]
	s = strings.TrimPrefix(s, prefix)
	s = strings.TrimSuffix(s, token)
	return s
}

// GetBase64Data retrieves the data from the base64 input
func GetBase64Data(i string) string {
	idx := strings.LastIndex(i, separator)
	if idx == -1 {
		return ""
	}
	return i[idx+len(separator):]
}
