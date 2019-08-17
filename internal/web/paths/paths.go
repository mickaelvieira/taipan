package paths

import (
	"os"
)

// GetBasePath returns base path
func GetBasePath(withFileServer bool) string {
	if withFileServer {
		return "/static"
	}
	return ""
}

// GetTemplatesDir returns the path to the directory containing the template files
func GetTemplatesDir() string {
	return os.Getenv("APP_WEB_DIR") + "/templates"
}

// GetStaticDir returns the path to the directory containing the static files
func GetStaticDir() string {
	return os.Getenv("APP_WEB_DIR") + "/static"
}

// GetScriptsDir returns the path to the directory containing the Javascript files
func GetScriptsDir() string {
	return os.Getenv("APP_WEB_DIR") + "/static/js"
}

// GetCDNBaseURL returns the CDN base URL
func GetCDNBaseURL() string {
	return "https://" + os.Getenv("AWS_BUCKET")
}

// GetGraphQLSchema returns the path to the graphQL schema
func GetGraphQLSchema() string {
	return os.Getenv("APP_WEB_DIR") + "/graphql/schema.graphql"
}
