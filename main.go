/**
 * Created on 23/2/13 by cos
 */
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

var root string
var port int

func main() {
	flag.StringVar(&root, "h", "", "home directory")
	flag.IntVar(&port, "p", 8080, "port")
	flag.Parse()

	if len(root) <= 0 {
		log.Fatal("usage: server -h <Home directory> -p <PORT>")
	}

	http.HandleFunc("/", router)
	log.Print("start listen at ", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func router(w http.ResponseWriter, r *http.Request) {
	log.Printf("serving %s", r.RequestURI)

	path := r.URL.Path
	log.Print("path = ", path)

	if "/" == path {
		// log.Print("list dir")
		http.ServeFile(w, r, root)
		return
	}

	if path[0] == '/' {
		local, err := url.QueryUnescape(root + "/" + path[1:])
		if err != nil {
			http.Error(w, path[1:], http.StatusNotFound)
			return
		}
		log.Printf("serve file: %s", local)
		http.ServeFile(w, r, local)
		return
	}

	http.Error(w, fmt.Sprintf("Page not found: %s", r.RequestURI), http.StatusNotFound)
}
