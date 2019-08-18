package assets

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// AssetsBasePath assets directory
const AssetsBasePath = "/static"

// Assets contains the list of JS scripts
type Assets []string

// MakeCDNBaseURL returns the CDN base URL
func MakeCDNBaseURL() string {
	return "https://" + os.Getenv("AWS_BUCKET")
}

// LoadAssetsDefinition loads the list of available assets
func LoadAssetsDefinition(path string, useFileServer bool) Assets {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	var assets []string
	for _, file := range files {
		if strings.HasPrefix(file.Name(), "vendor") ||
			strings.HasPrefix(file.Name(), "app") ||
			strings.HasPrefix(file.Name(), "materialui") ||
			strings.HasPrefix(file.Name(), "react") && strings.HasSuffix(file.Name(), ".js") {
			if useFileServer {
				assets = append(assets, AssetsBasePath+"/js/"+file.Name())
			} else {
				assets = append(assets, "/js/"+file.Name())
			}
		}
	}
	return assets
}