package assets

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Assets represents the list of available assets
type Assets struct {
	Styles string `json:"styles"`
	App    string `json:"app"`
	Vendor string `json:"vendor"`
}

// LoadAssetsDefinition loads the list of available assets
func LoadAssetsDefinition() Assets {
	content, err := ioutil.ReadFile("../../web/static/hashes.json")

	if err != nil {
		log.Fatal(err)
	}

	var assets Assets
	json.Unmarshal(content, &assets)

	return assets
}
