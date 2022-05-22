package backend

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func Start() {
	// TODO: load configuration file

	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/{providerId}/{id}", item)

	AddMiddleware(r)

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
	panic("not implemented!")
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
