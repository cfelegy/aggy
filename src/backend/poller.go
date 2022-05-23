package backend

import (
	"fmt"
	"net/http"

	"github.com/cfelegy/aggy/src/backend/providers"
)

// Providers contains a list of all providers registered with the application
var Providers map[string]providers.Provider

func init() {
	Providers = make(map[string]providers.Provider)

	l := providers.LobstersProvider{}
	Providers[l.Name()] = l
}

// Poller holds the state for the poller, which polls the providers for new
// metas and comments
type Poller struct {
	context *providers.Context
}

// NewPoller sets up the state and returns an instance of a Poller
func NewPoller() *Poller {
	return &Poller{
		context: &providers.Context{
			Client: http.DefaultClient,
		},
	}
}

// PollAllProviders calls PollProvider on each of the providers in the Providers
// map. It returns a slice of pointers to Meta objects, as well as a slice of
// errors. This is so that if one provider returns an error, the entire call
// is not invalidated.
func (p *Poller) PollAllProviders() ([]*providers.Meta, []error) {
	metas := make([]*providers.Meta, 0)
	errors := make([]error, 0)

	for _, provider := range Providers {
		m, err := p.PollProvider(provider)
		if err != nil {
			errors = append(errors, err)
		}
		metas = append(metas, m...)
	}

	return metas, errors
}

// PollProviders polls the top-level metas off of a provider, and returns
// a slice of pointers to Meta objects, and an error.
func (p *Poller) PollProvider(provider providers.Provider) ([]*providers.Meta, error) {
	m, err := provider.GetMetas(p.context)
	if err != nil {
		name := provider.Name()
		if err != nil {
			err = fmt.Errorf("failed to poll provider '%s': %w", name, err)
			return nil, err
		}
	}
	return m, nil
}
