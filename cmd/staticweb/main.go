package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/tachesimazzoca/go-httppack/interceptors"

	"github.com/tachesimazzoca/go-httppack/config"
	"github.com/tachesimazzoca/go-httppack/servers"
)

func main() {
	// flags
	flgs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	docRoot := flgs.String("d", ".", "document root")
	port := flgs.Int("p", 4000, "port")
	if err := flgs.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}

	// interceptors
	logging := &interceptors.LoggingInterceptor{}
	interceptors := []interceptors.Interceptor{logging}
	cfg := config.New(
		config.DocumentRoot(*docRoot),
		config.Port(*port),
		config.Interceptors(interceptors))

	// serve
	err := servers.NewStaticWebServer(cfg).ListenAndServe()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
