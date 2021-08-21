package config

import (
	"testing"

	"github.com/tachesimazzoca/go-httppack/detectors"
)

func TestHost(t *testing.T) {
	cfg := New()
	if actual := cfg.Host(); actual != "" {
		t.Error("Host expected:127.0.0.1 actual: ", actual)
	}
}

func TestPort(t *testing.T) {
	cfg := New()
	if actual := cfg.Port(); actual != 4000 {
		t.Error("Port expected:4000 actual: ", actual)
	}
}

func TestDocumentRoot(t *testing.T) {
	type pattern struct {
		docRoot  string
		expected string
	}
	patterns := []pattern{
		{"", "."},
		{"..", ".."},
		{"public/", "public"},
		{"public///", "public"},
		{"/var/www/", "/var/www"},
	}
	for _, ptn := range patterns {
		cfg := New(DocumentRoot(ptn.docRoot))
		if actual := cfg.DocumentRoot(); actual != ptn.expected {
			t.Errorf("%s expected:%s actual:%s", ptn.docRoot, ptn.expected, actual)
		}
	}
}
func TestEnsureBufferSize(t *testing.T) {
	type pattern struct {
		n        int
		expected int
	}
	patterns := []pattern{
		{-1, defaultBufferSize},
		{0, defaultBufferSize},
		{1, 1},
		{2, 2},
	}
	for _, ptn := range patterns {
		cfg := New(BufferSize(ptn.n))
		if actual := cfg.BufferSize(); actual != ptn.expected {
			t.Errorf("%d expected: %d, actual: %d", ptn.n, ptn.expected, actual)
		}
	}
}

func TestEnsureMimeDetectors(t *testing.T) {
	type pattern struct {
		mds      []detectors.Detector
		expected int
	}
	patterns := []pattern{
		{nil, 0},
		{[]detectors.Detector{}, 0},
		{[]detectors.Detector{nil}, 1},
	}
	for _, ptn := range patterns {
		cfg := New(MimeDetectors(ptn.mds))
		if actual := len(cfg.MimeDetectors()); actual != ptn.expected {
			t.Errorf("%s expected: %d, actual: %d", ptn.mds, ptn.expected, actual)
		}
	}
}
