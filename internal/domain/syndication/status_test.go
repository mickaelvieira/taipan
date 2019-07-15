package syndication

import (
	"fmt"
	"testing"
)

func TestIsTemporaryError(t *testing.T) {
	var testcase = []int{
		400,
		405,
		406,
		407,
		408,
		409,
		411,
		412,
		413,
		416,
		417,
		418,
		425,
		431,
		429,
		500,
		501,
		502,
		503,
		504,
		505,
		506,
		507,
		508,
		510,
		510,
		511,
	}

	for _, s := range testcase {
		name := fmt.Sprintf("Test IsTemporaryError [%d]", s)
		t.Run(name, func(t *testing.T) {
			r := IsHTTPErrorTemporary(s)
			if r != true {
				t.Errorf("Got [%t], wanted [true]", r)
			}
		})
	}
}

func TestIsPermamentError(t *testing.T) {
	var testcase = []int{
		401,
		402,
		403,
		404,
		410,
		414,
		415,
		421,
		422,
		423,
		424,
		426,
		428,
		451,
	}

	for _, s := range testcase {
		name := fmt.Sprintf("Test IsPermanentError [%d]", s)
		t.Run(name, func(t *testing.T) {
			r := IsHTTPErrorPermanent(s)
			if r != true {
				t.Errorf("Got [%t], wanted [true]", r)
			}
		})
	}
}
