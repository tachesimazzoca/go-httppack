package detectors

import (
	"testing"
)

func TestSuffixMimeDetector(t *testing.T) {
	dt := &SuffixMimeDetector{"text/html", []string{".html"}}
	if _, err := dt.Detect("index.html"); err != nil {
		t.Error(err)
	}
}

func TestDetectMimeType(t *testing.T) {
	type pattern struct {
		document string
		expected string
	}
	patterns := []pattern{
		{"", "application/octet-stream"},
		{"index.html", "text/html"},
		{"index.htm", "text/html"},
		{"robot.txt", "text/plain"},
		{"vue.min.js", "text/javascript"},
		{"style.css", "text/css"},
		{"app.json", "application/json"},
		{"config.xml", "application/xml"},
		{"config.yml", "application/x-yaml"},
		{"photo.jpg", "image/jpeg"},
		{"logo.png", "image/png"},
		{"spacer.gif", "image/gif"},
	}
	for _, ptn := range patterns {
		actual := DetectMimeType(ptn.document, BuiltInMimeDetectors)
		if actual != ptn.expected {
			t.Errorf("%s expected: %s, actual: %s", ptn.document, ptn.expected, actual)
		}
	}
}
