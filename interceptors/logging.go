package interceptors

import (
	"log"
	"net/http"
)

type LoggingInterceptor struct {
}

func (f *LoggingInterceptor) Intercept(handler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			log.Printf("%s %s", req.Method, req.RequestURI)
			handler.ServeHTTP(rw, req)
		})
}
