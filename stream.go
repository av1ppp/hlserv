package hlserv

import (
	"github.com/av1ppp/hlserv/ffmpeg"
	"github.com/thanhpk/randstr"
)

var streams = []*Stream{}

type Stream struct {
	ID     string
	Worker *ffmpeg.FFmpegWorker
	Config StreamConfig

	ready bool
	dir   string
}

func (stream *Stream) buildWorker() error {
	var input, output ffmpeg.OptionIO
	var err error

	switch stream.Config.Format {
	case "rtsp":
		input, output = rtspToHLS(stream)
	case "mp4":
		input, output = mp4ToHLS(stream)
	default:
		return ErrUnknownFormat
	}

	if stream.Worker, err = ffmpeg.NewWorker(stream.ID, input, output); err != nil {
		return err
	}

	return nil
}

type StreamConfig struct {
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
func CreateStream(conf StreamConfig) (*Stream, error) {
	var err error

	streamID := randstr.String(12)
	stream := Stream{
		Config: conf,
		ID:     streamID,
		dir:    streamID,
	}

	if err = stream.buildWorker(); err != nil {
		return nil, err
	}

	streams = append(streams, &stream)

	if err = stream.Worker.Start(); err != nil {
		return nil, err
	}

	// wait
	// TODO: add timeout
	for {
		if stream.ready {
			break
		}
	}

	return &stream, nil
}

func GetStream(id string) (*Stream, error) {
	for _, stream := range streams {
		if stream.ID == id {
			return stream, nil
		}
	}
	return nil, ErrStreamIsNotDefined
}

func RemoveStream(id string) error {
	for i, stream := range streams {
		if stream.ID == id {
			if err := stream.Worker.Stop(); err != nil {
				return err
			}
			ffmpeg.RmWorker(id)
			streams = append(streams[:i], streams[i+1:]...)

			return nil
		}
	}

	return ErrStreamIsNotDefined
}

func Streams() []*Stream {
	return streams
}
