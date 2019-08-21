package resolvers

import (
	"os"
)

// AppRootResolver app's root resolver
type AppRootResolver struct{}

// AppInfo resolves information about the application
type AppInfo struct{}

// Name return the application's name
func (r *AppInfo) Name() string {
	return os.Getenv("APP_NAME")
}

// Version return the application's version
func (r *AppInfo) Version() string {
	return os.Getenv("APP_VERSION")
}

// Info returns information about the application
func (r *AppRootResolver) Info() *AppInfo {
	return &AppInfo{}
}
