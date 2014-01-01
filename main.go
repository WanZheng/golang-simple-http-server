/**
 * Created on 23/2/13 by cos
 */
package main

import (
	"net/http"
	"log"
	"fmt"
	"net/url"
	"flag"
)

var root = flag.String("h", "", "home directory")
var port = flag.Int("p", 8080, "port")

func main() {
	flag.Parse()

	if len(*root) <= 0 {
		log.Fatal("usage: server -h <Home directory> -p <PORT>")
	}

	http.HandleFunc("/", router)
	log.Print("start listen at ", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

func router(w http.ResponseWriter, r *http.Request) {
	log.Printf("serving %s", r.RequestURI)

	uri := r.RequestURI
	if "/" == uri {
		// log.Print("list dir")
		http.ServeFile(w, r, *root)
		return
	}

	if uri[0] == '/' {
		path, err := url.QueryUnescape(*root + "/" + uri[1:])
		if err != nil {
			http.Error(w, uri[1:], http.StatusNotFound)
			return
		}
		// log.Printf("serve file: %s", path)
		http.ServeFile(w, r, path)
		return
	}

	http.Error(w, fmt.Sprintf("Page not found: %s", r.RequestURI), http.StatusNotFound)
}
