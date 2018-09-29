package servers

import (
	"net/http"

	"github.com/tachesimazzoca/go-httppack/config"
	"github.com/tachesimazzoca/go-httppack/handlers"
)

func NewStaticWebServer(cfg config.Config) *http.Server {
	mux := http.NewServeMux()
	mux.Handle("/", handlers.NewAssetsHandler(cfg))

	var handler http.Handler
	handler = mux
	for _, x := range cfg.Interceptors() {
		handler = x.Intercept(handler)
	}
	return &http.Server{
		Addr:    cfg.Addr(),
		Handler: handler,
	}
}
