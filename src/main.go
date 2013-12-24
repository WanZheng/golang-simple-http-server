/**
 * Created on 23/2/13 by cos
 */
package main

import (
	"net/http"
	"log"
	"fmt"
	"net/url"
	// "mime"
)

func main() {
	// mime.AddExtensionType(".rmvb", "application/vnd.rn-realmedia-vbr")

	http.HandleFunc("/", router)
	log.Print("start listen at " + PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}

func router(w http.ResponseWriter, r *http.Request) {
	log.Printf("uri = %s", r.RequestURI)

	uri := r.RequestURI
	if "/" == uri {
		log.Print("list dir")
		/*
		if err := listDir(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		*/
		http.ServeFile(w, r, ROOT_DIR)
		return
	}

	if uri[0] == '/' {
		path, err := url.QueryUnescape(ROOT_DIR + "/" + uri[1:])
		if err != nil {
			http.Error(w, uri[1:], http.StatusNotFound)
			return
		}
		log.Printf("serve file: %s", path)
		http.ServeFile(w, r, path)
		return
	}

	http.Error(w, fmt.Sprintf("Page not found: %s", r.RequestURI), http.StatusNotFound)
}

/*
func listDir(w http.ResponseWriter, r *http.Request) error{
	dir, err := os.Open(ROOT_DIR)
	if err != nil {
		return err
	}
	defer dir.Close()

	list, err := dir.Readdir(-1)
	if err != nil {
		return err
	}

	files := make([]os.FileInfo, 0, len(list))
	for _, file := range list {
		if file.IsDir() {
			continue
		}
		files = append(files, file)
	}

	w.Header().Set("Content-Type", "text/html")

	fmt.Fprint(w, "<html><header><meta http-equiv='Content-Type' content='text/html; charset=utf-8'></header>")

	fmt.Fprint(w, "<body><ul>\n")
	for _, file := range files {
		fmt.Fprintf(w, "  <li><a href='/%s'>%s</a></li>\n", file.Name(), file.Name())
	}
	fmt.Fprint(w, "</ul></body></html>\n")

	return nil
}
*/
