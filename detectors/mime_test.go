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
		pattern{"", "application/octet-stream"},
		pattern{"index.html", "text/html"},
		pattern{"index.htm", "text/html"},
		pattern{"robot.txt", "text/plain"},
		pattern{"vue.min.js", "text/javascript"},
		pattern{"style.css", "text/css"},
		pattern{"app.json", "application/json"},
		pattern{"config.xml", "application/xml"},
		pattern{"config.yml", "application/x-yaml"},
		pattern{"photo.jpg", "image/jpeg"},
		pattern{"logo.png", "image/png"},
		pattern{"spacer.gif", "image/gif"},
	}
	for _, ptn := range patterns {
		actual := DetectMimeType(ptn.document, BuiltInMimeDetectors)
		if actual != ptn.expected {
			t.Errorf("%s expected: %s, actual: %s", ptn.document, ptn.expected, actual)
		}
	}
}
