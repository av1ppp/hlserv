package hlserv

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
)

type urlParams struct {
	filename string
	streamID string
}

func (params *urlParams) name() string {
	return params.streamID + "/" + params.filename
}

func parseUrlParams(u *url.URL) (*urlParams, error) {
	splitPath := strings.Split(u.Path, "/")

	if len(splitPath) < 3 {
		return nil, ErrInvalidURL
	}

	return &urlParams{
		filename: splitPath[len(splitPath)-1],
		streamID: splitPath[len(splitPath)-2],
	}, nil
}

// Обработчик запросов на загрузку сегментов
var PutHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params, err := parseUrlParams(r.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := store.Write(params.name(), data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
})

// Обработчик запросов на удаление сегментов
var DeleteHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params, err := parseUrlParams(r.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := store.Remove(params.name()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
})

// Обработчик запросов на получение сегментов
var GetHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params, err := parseUrlParams(r.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file, err := store.File(params.name())
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	defer file.Reset()

	if filepath.Ext(params.filename) == ".ts" {
		w.Header().Add("Content-Type", TS_MIMETYPE)
	} else {
		w.Header().Add("Content-Type", M3U8_MIMETYPE)
	}

	io.Copy(w, file)
})

var Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetHandler.ServeHTTP(w, r)
	case http.MethodPut:
		PutHandler.ServeHTTP(w, r)
	case http.MethodDelete:
		DeleteHandler.ServeHTTP(w, r)
	}
})
