package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/av1ppp/hlserv"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var addr = "localhost:5555"

func main() {
	var streamID string
	var err error

	hlserv.EndPoint = "http://" + addr + "/streams"

	// start streaming
	conf := hlserv.StreamConfig{
		Format: "rtsp",
		Input:  "rtsp://admin:12345678@192.168.1.20:554/ch01/0",
	}
	if streamID, err = hlserv.CreateStream(&conf); err != nil {
		panic(err)
	}

	// start http server
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	// index.html
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open("examples/chi/index.html")
		if err != nil {
			panic(err)
		}
		io.Copy(w, file)
	})

	r.Route("/streams", func(r chi.Router) {
		// send all streams
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			if streamID == "" {
				w.Write([]byte("[]"))
			} else {
				w.Write([]byte(fmt.Sprintf(`[ { "id": "%s" } ]`, streamID)))
			}
		})

		r.Route("/{streamID}/{filename}", func(r chi.Router) {
			r.Handle("/", hlserv.Handler)
		})
	})

	log.Printf("Open http://%s/ in your browser", addr)
	if err = http.ListenAndServe(addr, r); err != nil {
		panic(err)
	}
}
