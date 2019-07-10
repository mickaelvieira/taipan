package image

import (
	"encoding/base64"
	"io"
	"strings"
)

const prefix = "data:"
const token = ";base64"
const separator = ","

// GetBase64Reader returns an IO reader for the base64 data
func GetBase64Reader(d string) io.Reader {
	return base64.NewDecoder(base64.StdEncoding, strings.NewReader(d))
}

// GetBase64ContentType retrieves the content type from the base64 URI
func GetBase64ContentType(u string) string {
	idx := strings.LastIndex(u, separator)
	if idx < 0 {
		return ""
	}
	s := u[0:idx]
	s = strings.TrimPrefix(s, prefix)
	s = strings.TrimSuffix(s, token)
	return s
}

// GetBase64Data retrieves the data from the base64 URI
func GetBase64Data(u string) string {
	idx := strings.LastIndex(u, separator)
	if idx < 0 {
		return ""
	}
	return u[idx+len(separator):]
}
