package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
    "flag"

	"github.com/rbastic/webdav"
	"github.com/facebookgo/grace/gracehttp"
)

var (
	path = "./tmp"
    port = flag.String("httpport", "8080", "port to listen on")
)

func newHandler(name string) http.Handler {

	mux := http.NewServeMux()
	mux.Handle("/", &webdav.Server{
		Fs:         webdav.Dir(path),
		Listings:   true,
	})

    return mux;
}

func main() {
	os.Mkdir(path, 0755)

    flag.Parse()

	// http.StripPrefix is not working, webdav.Server has no knowledge
	// of stripped component, but needs for COPY/MOVE methods.
	// Destination path is supplied as header and needs to be stripped.

	gracehttp.Serve(
		&http.Server{Addr: ":"+*port, Handler: newHandler("Zero  ")},
    )

	log.Println("Listening on http://127.0.0.1:"+*port)
}

// TODO: do something else with this
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q\n", r.URL.Path)
}
