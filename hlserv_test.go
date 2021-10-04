package hlserv_test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/av1ppp/hlserv"
)

var addr = "127.0.0.1:5555"

func init() {
	hlserv.Addr = addr
}

func TestHLSServer(t *testing.T) {
	go http.ListenAndServe(addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			hlserv.O.ServeHTTP(w, r)
		} else {
			hlserv.I.ServeHTTP(w, r)
		}
	}))

	streamID, err := hlserv.CreateStream(&hlserv.StreamConfig{
		Format: "rtsp",
		Input:  "rtsp://admin:12345678@192.168.1.20:554/ch01/0",
	})
	if err != nil {
		t.Fatalf("error adding stream: %s", err)
	}
	fmt.Println("STREAM ID", streamID)

	time.Sleep(time.Second * 5)

	resp, err := http.Get("http://" + addr + "/stream/" + streamID + "/stream.m3u8")
	if err != nil {
		t.Fatalf("error request: %s", err)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("error read body: %s", err)
	}

	time.Sleep(time.Second * 5)

	if err := hlserv.RemoveStream(streamID); err != nil {
		t.Fatalf("error removeing stream: %s", err)
	}

	if !bytes.Contains(data, []byte("#EXTM3U")) {
		t.Fatalf("m3u8 file is invalid: %s", string(data))
	}
}
