package providers

import "time"

// TODO: define the full interface for Provider and allow streaming of data

// Provider defines an interface for providers to feed data up to aggy
type Provider interface {
	GetMetas(context *Context) ([]*Meta, error)
	Name() string
}

// Meta describes a top-level post on a provided site
type Meta struct {
	// Id describes the Id of the item within its provider
	Id string `json:"id"`
	// ProviderId describes the Id of the provider in aggy, to allow
	// resolution to its source
	ProviderId string `json:"provider_id"`
	// Title describes the title of the item
	Title string `json:"title"`
	// Author describes the author (or poster) of the item
	Author string `json:"author"`
	// Timestamp describes the timestamp the item was submitted
	Timestamp time.Time `json:"timestamp"`
	// Url describes where to locate the item
	Url string `json:"url"`
}
