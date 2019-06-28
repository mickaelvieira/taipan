package resolvers

import "os"

// AppResolver app's root resolver
type AppResolver struct{}

// AppInfoResolver resolves information about the application
type AppInfoResolver struct{}

// Name return the application's name
func (r *AppInfoResolver) Name() string {
	return os.Getenv("APP_NAME")
}

// Version return the application's version
func (r *AppInfoResolver) Version() string {
	return os.Getenv("APP_VERSION")
}

// Info returns information about the application
func (r *AppResolver) Info() *AppInfoResolver {
	return &AppInfoResolver{}
}
