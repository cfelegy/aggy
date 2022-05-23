package providers

import "net/http"

// Context defines the context for all requests using a provider
type Context struct {
	// Client is an HTTP client with which to make requests
	Client *http.Client
}
