package backend

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func AddMiddleware(r *mux.Router) {
	r.Use(middlewareLog, middlewareRecover)
}

func middlewareRecover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				if errorErr, ok := err.(error); ok {
					log.WithError(errorErr).Errorln("recovered error from request")
				} else {
					log.WithField("recovered", err).Errorln("recovered from request")
				}

				w.WriteHeader(http.StatusInternalServerError)
				// TODO: write a JSON body containing the reason
			}
		}()
		next.ServeHTTP(w, r)
	})
}

type logResponseRecorder struct {
	status int
	http.ResponseWriter
}

func (w *logResponseRecorder) WriteHeader(statusCode int) {
	w.status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func middlewareLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrappedWriter := &logResponseRecorder{
			ResponseWriter: w,
			status:         0,
		}
		defer func() {
			end := time.Now()
			code := wrappedWriter.status
			if code == 0 {
				// default to Status OK
				w.WriteHeader(http.StatusOK)
				code = http.StatusOK
			}
			log.Debugf("%s %s: %d took %dns",
				r.Method,
				r.URL,
				code,
				end.Sub(start).Nanoseconds(),
			)
		}()
		next.ServeHTTP(wrappedWriter, r)
	})
}
