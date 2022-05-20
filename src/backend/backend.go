package backend

import "net/http"

func Start() {
	// TODO: load configuration file

	srv := http.Server{
		Addr: "0.0.0.0:4000", // TODO: from configuration
	}
	err := srv.ListenAndServe() // TODO: use TLS
	if err != nil {
		// TODO: report error
		panic(err)
	}
}
