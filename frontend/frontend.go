package frontend

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gobuffalo/packr/v2"
)

var appBox = packr.New("app", "./app/dist")

func IndexRoute(w http.ResponseWriter, r *http.Request) {
	indexHTML, err := appBox.Find("index.html")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("content-type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(indexHTML))
}

func FaviconRoute(w http.ResponseWriter, r *http.Request) {
	favicon, err := appBox.Find("favicon.ico")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(favicon))
}

func StaticRoute(w http.ResponseWriter, r *http.Request) {
	path := r.URL.RequestURI()
	if strings.HasSuffix(path, ".css") {
		w.Header().Set("content-type", "text/css")
	}
	if strings.HasSuffix(path, ".js") {
		w.Header().Set("content-type", "text/javascript")
	}
	file, err := appBox.Find(path)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(file))
}
