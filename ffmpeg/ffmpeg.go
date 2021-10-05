// Этот пакет - надстройка для использования FFmpeg в Go.
// Можно использовать более адекватные варианты, такие
// как https://github.com/giorgisio/goav, но пока нет времени
// разбираться с этим.

package ffmpeg

import (
	"os"
	"strconv"
	"sync"
)

type OptionIO interface {
	String() string
}

type report struct {
	File     string
	LogLevel int
}

type FFmpeg struct {
	BinPath string
	Report  report
	workers []*FFmpegWorker

	sync.RWMutex
}

func New() *FFmpeg {
	var (
		bin, reportFile string
		reportLogLevel  int
	)

	if _bin := os.Getenv("FFMPEG_BIN"); _bin != "" {
		bin = _bin
	} else {
		bin = "ffmpeg"
	}

	// if _reportFile := os.Getenv("FFMPEG_REPORT_FILE"); _reportFile != "" {
	// 	reportFile = _reportFile
	// } else {
	// 	reportFile = "logs/ffreport-%t.log"
	// }
	reportFile = "/dev/null"

	if _reportLogLevel, err := strconv.Atoi(os.Getenv("FFMPEG_REPORT_LOG_LEVEL")); err != nil {
		reportLogLevel = _reportLogLevel
	} else {
		reportLogLevel = 16 // error
	}

	return &FFmpeg{
		BinPath: bin,
		Report: report{
			File:     reportFile,
			LogLevel: reportLogLevel,
		},
		workers: []*FFmpegWorker{},
	}
}
