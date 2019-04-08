package detectors

import (
	"fmt"
	"strings"
)

type SuffixMimeDetector struct {
	MimeType string
	Suffix   []string
}

func (md *SuffixMimeDetector) Detect(s string) (string, error) {
	ok := false
	for _, x := range md.Suffix {
		if ok = strings.HasSuffix(s, x); ok {
			break
		}
	}
	if ok {
		return md.MimeType, nil
	} else {
		return "", fmt.Errorf("Not match %s", md)
	}
}

var BuiltInMimeDetectors = []Detector{
	&SuffixMimeDetector{"text/html", []string{".html", ".htm"}},
	&SuffixMimeDetector{"text/css", []string{".css"}},
	&SuffixMimeDetector{"text/javascript", []string{".js"}},
	&SuffixMimeDetector{"image/jpeg", []string{".jpg", ".jpeg"}},
	&SuffixMimeDetector{"image/gif", []string{".gif"}},
	&SuffixMimeDetector{"image/png", []string{".png"}},
	&SuffixMimeDetector{"image/svg+xml", []string{".svg"}},
	&SuffixMimeDetector{"application/xml", []string{".xml"}},
	&SuffixMimeDetector{"application/json", []string{".json"}},
	&SuffixMimeDetector{"application/x-yaml", []string{".yml", ".yaml"}},
	&SuffixMimeDetector{"text/plain", []string{".txt"}},
}

func DetectMimeType(doc string, mds []Detector) string {
	for _, md := range mds {
		if mt, err := md.Detect(doc); err == nil {
			return mt
		}
	}
	return "application/octet-stream"
}
