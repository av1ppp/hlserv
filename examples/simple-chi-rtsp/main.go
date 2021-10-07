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
	var stream *hlserv.Stream
	var err error

	hlserv.SetEndPoint("http://" + addr + "/streams")

	// start http server
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	// index.html
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open("examples/simple-chi-rtsp/index.html")
		if err != nil {
			panic(err)
		}
		io.Copy(w, file)
	})

	r.Route("/streams", func(r chi.Router) {
		// send all streams
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			w.Write([]byte(fmt.Sprintf(`[ { "id": "%s" } ]`, stream.ID)))
		})

		r.Route("/{streamID}/{filename}", func(r chi.Router) {
			r.Handle("/", hlserv.Handler)
		})
	})

	ch := make(chan error)
	go func() {
		if err = http.ListenAndServe(addr, r); err != nil {
			ch <- err
		}
	}()

	// start streaming
	stream, err = hlserv.CreateStream(hlserv.StreamConfig{
		Format: "rtsp",
		Source: "rtsp://admin:12345678@192.168.1.20:554/ch01/0",
		FPS:    10,
	})
	if err != nil {
		panic(err)
	}

	log.Printf("Open http://%s/ in your browser", addr)

	if err := <-ch; err != nil {
		panic(err)
	}
}
