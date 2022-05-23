package providers

import "errors"

type HnProvider struct {
}

func (p HnProvider) Name() string {
	return "hn"
}

func (p HnProvider) GetMetas(context *Context) ([]*Meta, error) {
	// todo: hn api keeps kicking back a 503
	return nil, errors.New("not implemented")
}
