package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/av1ppp/hlserv"
	"github.com/go-chi/chi/v5"
)

const (
	addr = "localhost:5555"
)

// GET /
var homeHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("examples/mp4-stream/index.html")
	if err != nil {
		panic(err)
	}
	io.Copy(w, file)
})

// GET /streams
var getStreamsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var resp []struct {
		ID string `json:"id"`
	}

	for _, stream := range hlserv.Streams() {
		resp = append(resp, struct {
			ID string `json:"id"`
		}{
			ID: stream.ID,
		})
	}

	data, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	w.Write(data)
})

// POST /streams
var startStreamHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Source string `json:"source"`
		CRF    int    `json:"crf"` // quality
		Scale  string `json:"scale"`
	}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(data, &req); err != nil {
		panic(err)
	}

	var resp struct {
		StreamID string `json:"stream_id"`
	}

	// если уже есть такой стрим - возвращаем его
	for _, stream := range hlserv.Streams() {
		if stream.Config.Source == req.Source &&
			stream.Config.CRF == req.CRF &&
			stream.Config.Scale == req.Scale {
			resp.StreamID = stream.ID
		}
	}

	// если нет - создаем новый
	if resp.StreamID == "" {
		stream, err := hlserv.CreateStream(hlserv.StreamConfig{
			Format: "mp4",
			Source: req.Source,
			FPS:    10,
			CRF:    req.CRF,
			Scale:  req.Scale,
		})
		if err != nil {
			panic(err)
		}
		resp.StreamID = stream.ID
	}

	if data, err = json.Marshal(resp); err != nil {
		panic(err)
	}

	w.Header().Add("Conten-Type", "applicatoin/json")
	w.Write(data)
})

// DELETE /streams
var stopStreamHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	streamID := chi.URLParam(r, "streamID")

	if err := hlserv.RemoveStream(streamID); err != nil {
		panic(err)
	}
})

// GET /streams/<streamId>/<filename>
var getStreamSegment = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	hlserv.GetHandler.ServeHTTP(w, r)
})

// PUT /streams/<streamId>/<filename>
var updateStreamSegment = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	hlserv.PutHandler.ServeHTTP(w, r)
})

// DELETE /streams/<streamId>/<filename>
var deleteStreamSegment = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	hlserv.DeleteHandler.ServeHTTP(w, r)
})

func main() {
	r := chi.NewRouter()

	r.Get("/", homeHandler)
	r.Get("/streams", getStreamsHandler)
	r.Post("/streams", startStreamHandler)
	r.Delete("/streams/{streamID}", stopStreamHandler)
	r.Get("/streams/{streamID}/{filename}", getStreamSegment)
	r.Put("/streams/{streamID}/{filename}", updateStreamSegment)
	r.Delete("/streams/{streamID}/{filename}", deleteStreamSegment)

	fmt.Printf("Open http://%s/ in your browser\n", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		panic(err)
	}
}
