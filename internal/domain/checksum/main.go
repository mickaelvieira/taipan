package checksum

import (
	"crypto/sha256"
	"database/sql/driver"
	"encoding/hex"
	"errors"
)

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

// FromBytes creates a sha256 checksum from bytes
func FromBytes(b []byte) Checksum {
	buf := sha256.New()
	buf.Write(b)
	return buf.Sum(nil)
}
