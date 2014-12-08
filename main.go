package main

import (
	"flag"
	"net/http"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/rbastic/webdav"
)

var (
	port = flag.String("httpport", "8080", "port to listen on")
	dir  = flag.String("docroot", "www", "document root folder to serve from")
)

func main() {
	// TODO: Make sure we can read, write, etc., to *dir

	flag.Parse()

	// http.StripPrefix is not working, webdav.Server has no knowledge
	// of stripped component, but needs for COPY/MOVE methods.
	// Destination path is supplied as header and needs to be stripped.

	mux := http.NewServeMux()
	mux.Handle("/", &webdav.Server{
		Fs:         webdav.Dir(*dir),
		TrimPrefix: "/",
		Listings:   true,
	})

	gracehttp.Serve(
		&http.Server{Addr: ":" + *port, Handler: mux},
	)

	webdav.FlushGlog()
}
