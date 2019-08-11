package assets

import (
	"encoding/json"
	"github/mickaelvieira/taipan/internal/domain/url"
	"io/ioutil"
	"log"
	"os"
)

// AssetsBasePath assets directory
const AssetsBasePath = "/static"

// Assets represents the list of available assets
type Assets struct {
	App        string `json:"app"`
	Vendor     string `json:"vendor"`
	React      string `json:"react"`
	Materialui string `json:"materialui"`
}

// AppendFileServerBasePath removes the base path from assets filepaths
func (a *Assets) AppendFileServerBasePath() {
	a.App = AssetsBasePath + a.App
	a.Vendor = AssetsBasePath + a.Vendor
	a.React = AssetsBasePath + a.React
	a.Materialui = AssetsBasePath + a.Materialui
}

// MakeCDNBaseURL returns the CDN base URL
func MakeCDNBaseURL() string {
	return "https://" + os.Getenv("AWS_BUCKET")
}

// MakeImageURL returns an image's URL based on its name
func MakeImageURL(name string) *url.URL {
	u, err := url.FromRawURL("https://" + os.Getenv("AWS_BUCKET") + "/" + name)
	if err != nil {
		u = &url.URL{}
	}
	return u
}

// GetBasePath returns base path
func GetBasePath(useFileServer bool) string {
	if useFileServer {
		return AssetsBasePath
	}
	return ""
}

// LoadAssetsDefinition loads the list of available assets
func LoadAssetsDefinition(path string) *Assets {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	var assets Assets
	json.Unmarshal(content, &assets)

	return &assets
}
