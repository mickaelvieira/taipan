package assets

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// AssetsBasePath assets directory
const AssetsBasePath = "/static"

const jsDir = "js"
const cssDir = "css"

// Assets contains the list of JS and CSS files
type Assets struct {
	Scripts []string
	Styles  []string
}

// MakeCDNBaseURL returns the CDN base URL
func MakeCDNBaseURL() string {
	return "https://" + os.Getenv("AWS_BUCKET")
}

// LoadAssetsDefinition loads the list of available assets
func LoadAssetsDefinition(path string, useFileServer bool) *Assets {
	jsFiles, err := ioutil.ReadDir(path + "/" + jsDir)
	if err != nil {
		log.Fatal(err)
	}

	var scripts []string
	for _, file := range jsFiles {
		if strings.HasPrefix(file.Name(), "vendor") ||
			strings.HasPrefix(file.Name(), "app") ||
			strings.HasPrefix(file.Name(), "materialui") ||
			strings.HasPrefix(file.Name(), "react") && strings.HasSuffix(file.Name(), ".js") {
			if useFileServer {
				scripts = append(scripts, AssetsBasePath+"/"+jsDir+"/"+file.Name())
			} else {
				scripts = append(scripts, "/"+jsDir+"/"+file.Name())
			}
		}
	}

	cssFiles, err := ioutil.ReadDir(path + "/" + cssDir)
	if err != nil {
		log.Fatal(err)
	}

	var styles []string
	for _, file := range cssFiles {
		if strings.HasPrefix(file.Name(), "vendor") && strings.HasSuffix(file.Name(), ".css") {
			if useFileServer {
				styles = append(styles, AssetsBasePath+"/"+cssDir+"/"+file.Name())
			} else {
				styles = append(styles, "/"+cssDir+"/"+file.Name())
			}
		}
	}

	return &Assets{
		Scripts: scripts,
		Styles:  styles,
	}
}
