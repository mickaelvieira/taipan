package assets

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// AssetsBasePath assets directory
const AssetsBasePath = "/static"

// Assets represents the list of available assets
type Assets struct {
	App    string `json:"app"`
	Vendor string `json:"vendor"`
}

// AppendFileServerBasePath removes the base path from assets filepaths
func (a *Assets) AppendFileServerBasePath() {
	a.App = AssetsBasePath + a.App
	a.Vendor = AssetsBasePath + a.Vendor
}

// MakeCDNBaseURL returns the CDN base URL
func MakeCDNBaseURL() string {
	return "https://" + os.Getenv("AWS_BUCKET")
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
