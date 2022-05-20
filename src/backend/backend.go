package backend

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	// TODO: load configuration file

	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/{providerId}/{id}", item)

	srv := http.Server{
		Addr:    "0.0.0.0:4000", // TODO: from configuration
		Handler: r,
	}
	err := srv.ListenAndServe() // TODO: use TLS
	if err != nil {
		// TODO: report error
		panic(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	println("/")
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
	println(categoryId, id)
}
