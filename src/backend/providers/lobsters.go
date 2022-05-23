package providers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type LobstersProvider struct {
}

const lobstersName = "lobsters"

func (p LobstersProvider) Name() string {
	return lobstersName
}

const lobstersUrlBase = "https://lobste.rs/"
const lobstersUrlHottest = lobstersUrlBase + "hottest.json"

//const lobstersUrlNewest = lobstersUrlBase + "newest.json"

type lobstersApiMeta struct {
	ShortId   string          `json:"short_id"`
	Title     string          `json:"title"`
	CreatedAt time.Time       `json:"created_at"`
	Url       string          `json:"url"`
	User      lobstersApiUser `json:"submitter_user"`
}
type lobstersApiUser struct {
	Username string `json:"username"`
}

func (a lobstersApiMeta) toMeta() Meta {
	return Meta{
		Id:         a.ShortId,
		ProviderId: lobstersName,
		Title:      a.Title,
		Timestamp:  a.CreatedAt,
		Url:        a.Url,
		Author:     a.User.Username,
	}
}

func (p LobstersProvider) GetMetas(context *Context) ([]*Meta, error) {
	// todo: hottest/newest switch (put on context?)

	r, err := context.Client.Get(lobstersUrlHottest)
	if err != nil {
		return nil, fmt.Errorf("lobsters get failed at req: %w", err)
	}
	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("lobsters get failed at req: status %s", r.Status)
	}

	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("lobsters get failed at read: %w", err)
	}
	var apiMetas []lobstersApiMeta
	err = json.Unmarshal(b, &apiMetas)
	if err != nil {
		return nil, fmt.Errorf("lobsters get failed at unmarshal: %w", err)
	}

	metas := make([]*Meta, len(apiMetas))
	for i, a := range apiMetas {
		m := a.toMeta()
		metas[i] = &m
	}

	return metas, nil
}
