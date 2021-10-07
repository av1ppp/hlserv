package hlserv_test

import (
	"bytes"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/av1ppp/hlserv"
)

func TestHLSServer(t *testing.T) {
	go http.ListenAndServe(hlserv.EndPoint(), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hlserv.Handler.ServeHTTP(w, r)
	}))

	stream, err := hlserv.CreateStream(hlserv.StreamConfig{
		Format: "rtsp",
		Source: "rtsp://admin:12345678@192.168.1.20:554/ch01/0",
	})
	if err != nil {
		t.Fatalf("error adding stream: %s", err)
	}
	t.Log("Stream ID", stream.ID)

	time.Sleep(time.Second * 5)

	resp, err := http.Get("http://" + hlserv.EndPoint() + "/stream/" + stream.ID + "/stream.m3u8")
	if err != nil {
		t.Fatalf("error request: %s", err)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("error read body: %s", err)
	}

	time.Sleep(time.Second * 5)

	if err := hlserv.RemoveStream(stream.ID); err != nil {
		t.Fatalf("error removeing stream: %s", err)
	}

	if !bytes.Contains(data, []byte("#EXTM3U")) {
		t.Fatalf("m3u8 file is invalid: %s", string(data))
	}
}
