package config

import (
	"fmt"
	"strings"

	"github.com/tachesimazzoca/go-httppack/detectors"
	"github.com/tachesimazzoca/go-httppack/interceptors"
)

type Config interface {
	Addr() string
	Host() string
	Port() int
	DocumentRoot() string
	MimeDetectors() []detectors.Detector
	Interceptors() []interceptors.Interceptor
	BufferSize() int
}

type configInstance struct {
	host          string
	port          int
	documentRoot  string
	mimeDetectors []detectors.Detector
	interceptors  []interceptors.Interceptor
	bufferSize    int
}

func (cfg *configInstance) Addr() string {
	return fmt.Sprintf("%s:%d", cfg.host, cfg.port)
}

func (cfg *configInstance) Host() string {
	return cfg.host
}

func (cfg *configInstance) Port() int {
	return cfg.port
}

func (cfg *configInstance) DocumentRoot() string {
	return cfg.documentRoot
}

func (cfg *configInstance) MimeDetectors() []detectors.Detector {
	return cfg.mimeDetectors
}

func (cfg *configInstance) Interceptors() []interceptors.Interceptor {
	return cfg.interceptors
}

func (cfg *configInstance) BufferSize() int {
	return cfg.bufferSize
}

type ConfigOption func(*configInstance) error

const defaultPort = 4000

func Port(p int) ConfigOption {
	if p < 0 {
		p = defaultPort
	}
	return func(cfg *configInstance) error {
		cfg.port = p
		return nil
	}
}

func DocumentRoot(docRoot string) ConfigOption {
	s := strings.TrimSpace(docRoot)
	s = strings.TrimRight(s, "/")
	if s == "" {
		s = "."
	}
	return func(cfg *configInstance) error {
		cfg.documentRoot = s
		return nil
	}
}

func MimeDetectors(mds []detectors.Detector) ConfigOption {
	return func(cfg *configInstance) error {
		cfg.mimeDetectors = mds
		return nil
	}
}

func Interceptors(ics []interceptors.Interceptor) ConfigOption {
	return func(cfg *configInstance) error {
		cfg.interceptors = ics
		return nil
	}
}

const defaultBufferSize = 4096

func BufferSize(n int) ConfigOption {
	if n < 1 {
		n = defaultBufferSize
	}
	return func(cfg *configInstance) error {
		cfg.bufferSize = n
		return nil
	}
}

func New(opts ...ConfigOption) Config {
	cfg := &configInstance{
		host:          "",
		port:          defaultPort,
		documentRoot:  ".",
		mimeDetectors: detectors.BuiltInMimeDetectors,
		interceptors:  []interceptors.Interceptor{},
		bufferSize:    defaultBufferSize,
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}
