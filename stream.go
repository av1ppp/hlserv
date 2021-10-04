package hlserv

import (
	"github.com/av1ppp/hlserv/ffmpeg"
	"github.com/thanhpk/randstr"
)

type StreamConfig struct {
	Format string
	Input  string

	// Level   string
	// Profile string

	// VideoBitrate string
	// VideoProfile string
	// Bufsize      string
	// Scale        string
}

// Добавляет и запускает стрим с заданными настройками.
// Возвращает ID стрима.
func CreateStream(conf *StreamConfig) (string, error) {
	id := randstr.String(12)

	var input, output ffmpeg.OptionIO

	if conf.Format == "rtsp" {
		input, output = rtspToHLS(id, conf.Input)
	} else {
		// TODO: Return error
	}

	w, err := ffmpeg.NewWorker(id, input, output)

	if err != nil {
		return "", err
	}

	if err := w.Start(); err != nil {
		return "", err
	}

	return id, nil
}

func RemoveStream(id string) error {
	w, find := ffmpeg.Worker(id)
	if !find {
		return nil
	}
	if err := w.Stop(); err != nil {
		return err
	}

	ffmpeg.RmWorker(id)
	return nil
}
