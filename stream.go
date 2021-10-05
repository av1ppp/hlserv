package hlserv

import (
	"github.com/av1ppp/hlserv/ffmpeg"
	"github.com/thanhpk/randstr"
)

type StreamConfig struct {
	id string

	Source string

	Format       string
	AudioEnabled bool
	FPS          int
	Preset       ffmpeg.Preset
	Level        string
	Profile      string
	VideoBitrate string
	Bufsize      string
	Scale        string
	CRF          int
}

// Добавляет и запускает стрим с заданными настройками.
// Возвращает ID стрима.
func CreateStream(conf *StreamConfig) (string, error) {
	conf.id = randstr.String(12)

	var input, output ffmpeg.OptionIO

	switch conf.Format {
	case "rtsp":
		input, output = rtspToHLS(conf)
	case "mp4":
		input, output = mp4ToHLS(conf)
	default:
		return "", ErrUnknownFormat
	}

	w, err := ffmpeg.NewWorker(conf.id, input, output)

	if err != nil {
		return "", err
	}

	if err := w.Start(); err != nil {
		return "", err
	}

	return conf.id, nil
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

func Streams() []string {
	ids := []string{}
	for _, w := range ffmpeg.Workers() {
		ids = append(ids, w.Name)
	}
	return ids
}
