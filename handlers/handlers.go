package handlers

import (
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/tachesimazzoca/go-httppack/config"
	"github.com/tachesimazzoca/go-httppack/detectors"
)

type assetsHandler struct {
	cfg config.Config
}

func NewAssetsHandler(cfg config.Config) http.Handler {
	return &assetsHandler{cfg}
}

func (h *assetsHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	doc := path.Join(h.cfg.DocumentRoot(), directoryIndex(req.URL.Path))
	// check if the file exists
	_, err := os.Stat(doc)
	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	// open the file
	f, err := os.Open(doc)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()

	rw.Header().Set("Content-Type",
		detectors.DetectMimeType(doc, h.cfg.MimeDetectors()))

	rw.WriteHeader(http.StatusOK)

	buf := make([]byte, h.cfg.BufferSize())
	for {
		n, err := f.Read(buf)
		if err != nil {
			break
		}
		if n > 0 {
			_, err := rw.Write(buf[:n])
			if err != nil {
				break
			}
			continue
		}
		break
	}
}

func directoryIndex(s string) string {
	if s == "" {
		return "/index.html"
	}
	if strings.HasSuffix(s, "/") {
		return s + "index.html"
	} else {
		return s
	}
}
