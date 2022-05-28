package backend

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func AddMiddleware(r *mux.Router) {
	r.Use(middlewareCors, middlewareLog, middlewareRecover)
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

func middlewareRecover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				if errorErr, ok := err.(error); ok {
					log.WithError(errorErr).
						Errorln("recovered error from request")
				} else {
					log.WithField("recovered", err).
						Errorln("recovered from request")
				}

				w.WriteHeader(http.StatusInternalServerError)

				body := make(map[string]any)
				body["code"] = 500
				body["reason"] = fmt.Sprint(err)
				b, err := json.Marshal(body)
				if err != nil {
					log.WithError(err).
						Errorln("failed to marshal error for writing")
					return
				}
				_, err = w.Write(b)
				if err != nil {
					log.WithError(err).
						Errorln("failed to write error to response")
				}
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

			var durStr string
			dur := end.Sub(start)
			musPart := dur.Microseconds() % 1000
			msPart := dur.Milliseconds()
			durStr = fmt.Sprintf("%d.%dms", msPart, musPart)

			log.Debugf("%s %s: %d took %s",
				r.Method,
				r.URL,
				code,
				durStr,
			)
		}()
		next.ServeHTTP(wrappedWriter, r)
	})
}
