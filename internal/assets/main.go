package assets

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Assets represents the list of available assets
type Assets struct {
	App    string `json:"app"`
	Vendor string `json:"vendor"`
}

// MakeCDNBaseURL returns the CDN base URL
func MakeCDNBaseURL() string {
	return "https://" + os.Getenv("AWS_BUCKET")
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
