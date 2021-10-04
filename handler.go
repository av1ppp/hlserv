package hlserv

import (
	"io"
	"net/http"
	"path/filepath"
	"strings"
)

var O = http.HandlerFunc(outHandler)
var I = http.HandlerFunc(inHandler)

func outHandler(w http.ResponseWriter, r *http.Request) {
	splitPath := strings.Split(r.URL.Path, "/")
	filename := splitPath[len(splitPath)-1]

	file, err := store.File(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	defer file.Reset()

	ext := filepath.Ext(filename)

	if ext == ".ts" {
		w.Header().Add("Content-Type", TS_MIMETYPE)
	} else {
		w.Header().Add("Content-Type", M3U8_MIMETYPE)
	}

	// w.WriteHeader(http.StatusOK)
	io.Copy(w, file)
}

func inHandler(w http.ResponseWriter, r *http.Request) {
	splitPath := strings.Split(r.URL.Path, "/")
	filename := splitPath[len(splitPath)-1]

	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := store.Create(filename, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
