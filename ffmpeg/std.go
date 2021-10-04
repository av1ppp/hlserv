package ffmpeg

import (
	"sync"
)

type std struct {
	ffmpeg *FFmpeg
	// mu     sync.Mutex
	once sync.Once
}

var stdImpl = std{}

func stdFFmpeg() *FFmpeg {
	stdImpl.once.Do(func() {
		stdImpl.ffmpeg = New()
	})
	return stdImpl.ffmpeg
}

func RunOnceWorker(files ...OptionIO) error {
	return stdFFmpeg().RunOnceWorker(files...)
}

func NewWorker(key string, files ...OptionIO) (*FFmpegWorker, error) {
	return stdFFmpeg().NewWorker(key, files...)
}

func Worker(key string) (*FFmpegWorker, bool) {
	return stdFFmpeg().Worker(key)
}

func RmWorker(key string) {
	stdFFmpeg().RmWorker(key)
}
