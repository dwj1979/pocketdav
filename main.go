/*
 * pocketdav is a tiny 'webdav' server written in Go.

 * It only supports GET, HEAD, PUT, and DELETE.
 * It supports in-place restarts. You can simply
    'sudo mv pocketdav /usr/bin/pocketdav' while it's running
    and run kill -SIGUSR2 `/sbin/pidof pocketdav` to restart it.

    There are some potential CAVEATS about the current in-place restart approach.

 * It has extensive logging capabilities due to rbastic/webdav using 'glog'
 *  but it does them in a friendly way by default:
 * - directories are automatically created as necessary with a PUT operation
 * - no caching happens between GET/HEAD so a DELETE takes effect immediately
*/
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
