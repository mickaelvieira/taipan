package types

import (
	"database/sql/driver"
	"encoding/hex"
	"errors"
	"net/url"
)

// URI represents a URI within the application
type URI struct {
	*url.URL
}

// Value converts the value going into the DB
func (uri *URI) Value() (driver.Value, error) {
	value, err := url.QueryUnescape(uri.String())
	return value, err
}

// Scan converts the value coming from the DB
func (uri *URI) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	if v, ok := value.([]byte); ok {
		var u *url.URL
		u, err := url.ParseRequestURI(string(v))
		if err != nil {
			return errors.New("failed to parse URL during scanning")
		}
		*uri = URI{u}
		return nil
	}
	return errors.New("failed to scan URL")
}

// Checksum type
type Checksum []byte

func (c Checksum) String() string {
	return hex.EncodeToString(c)
}

// Value converts the value going into the DB
func (c Checksum) Value() (driver.Value, error) {
	if len(c) > 0 {
		return hex.EncodeToString(c), nil
	}
	return nil, nil
}

// Scan converts the value coming from the DB
func (c *Checksum) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	if v, ok := value.([]byte); ok {
		d, err := hex.DecodeString(string(v))
		if err != nil {
			return errors.New("failed to decode Checksum")
		}
		*c = d
		return nil
	}
	return errors.New("failed to scan Checksum")
}
