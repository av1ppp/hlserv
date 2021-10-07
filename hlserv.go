package hlserv

import "sync"

var endPoint = "http://localhost:5555/streams"
var mu sync.Mutex

func SetEndPoint(p string) {
	mu.Lock()
	defer mu.Unlock()
	endPoint = p
}

func EndPoint() string {
	mu.Lock()
	defer mu.Unlock()
	return endPoint
}
