package backend

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var poller *Poller

func Start() {
	// TODO: load configuration file

	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/{providerId}/{id}", item)

	AddMiddleware(r)

	poller = NewPoller()

	srv := http.Server{
		Addr:    "0.0.0.0:4000", // TODO: from configuration
		Handler: r,
	}
	err := srv.ListenAndServe() // TODO: use TLS
	if err != nil {
		log.WithError(err).Panicln("web server failed")
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	metas, errors := poller.PollAllProviders()
	if len(errors) > 0 {
		for _, e := range errors {
			log.WithError(e).Warnln("PollAllProviders encountered an error")
		}
	}
	b, err := json.Marshal(metas)
	if err != nil {
		panic(fmt.Errorf("failed to marshal metas: %w", err))
	}
	_, err = w.Write(b)
	if err != nil {
		panic(fmt.Errorf("failed to write bytes: %w", err))
	}
	w.WriteHeader(http.StatusOK)
}

func item(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var ok bool
	var id, categoryId string
	if categoryId, ok = vars["providerId"]; !ok {
		// TODO: error
		return
	}
	if id, ok = vars["id"]; !ok {
		// TODO: error
		return
	}
	_ = id
	_ = categoryId
	panic("not implemented!")
}
