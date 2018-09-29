package handlers

import (
	"net/http"
	"os"
	"path"
	"testing"

	"github.com/tachesimazzoca/go-httppack/config"
)

type mockResponseWriter struct {
	header     http.Header
	statusCode int
	body       [][]byte
}

func (rw *mockResponseWriter) Header() http.Header {
	return rw.header
}

func (rw *mockResponseWriter) WriteHeader(sc int) {
	rw.statusCode = sc
}

func (rw *mockResponseWriter) Write(buf []byte) (int, error) {
	rw.body = append(rw.body, buf)
	return len(buf), nil
}

func fromWorkingDir(s string) string {
	wd, err := os.Getwd()
	if err == nil {
		return path.Join(wd, s)
	} else {
		return s
	}
}

func TestAssetsHandler(t *testing.T) {
	rw := &mockResponseWriter{map[string][]string{}, 0, [][]byte{}}
	req := &http.Request{RequestURI: "/"}
	h := NewAssetsHandler(config.New(
		config.DocumentRoot(fromWorkingDir("/fixtures"))))
	h.ServeHTTP(rw, req)
	if actual := rw.Header().Get("Content-Type"); actual != "text/html" {
		t.Errorf("Content-Type expected: text/html, actual: %s", actual)
	}
	if actual := rw.statusCode; actual != 200 {
		t.Errorf("statusCode expected: 200, actual: %d", actual)
	}
	if actual := rw.body; string(actual[0]) != "test" {
		t.Errorf("body expected: test, actual: %s", actual)
	}
}

func TestDirectoryIndex(t *testing.T) {
	type pattern struct {
		uri      string
		expected string
	}
	patterns := []pattern{
		pattern{"", "/index.html"},
		pattern{"/", "/index.html"},
		pattern{"/index.html", "/index.html"},
		pattern{" ", "/index.html"},
		pattern{" /foo.html", "/foo.html"},
		pattern{"/bar.html ", "/bar.html"},
	}
	for _, ptn := range patterns {
		if actual := directoryIndex(ptn.uri); actual != ptn.expected {
			t.Errorf("%s expected: %s, actual: %s", ptn.uri, ptn.expected, actual)
		}
	}
}
