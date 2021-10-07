package hlserv

import (
	"testing"

	"github.com/av1ppp/hlserv/ffmpeg"
)

func TestPresetMP4(t *testing.T) {
	var (
		input1 = ffmpeg.OptsInMp4{
			File: "/mnt/4tb1/.archive/1/1/8d8f-6153d9b8-6153dae5.mp4",
		}
		input2 = ffmpeg.OptsInMp4{
			File: "/mnt/4tb1/.archive/1/3/8c8c-61538e12-61538f3e.mp4",
		}
		output = ffmpeg.OptsOutHLS{
			HLSTime:        20,
			HLSSegmentType: "mpegts",
			HLSFlags:       []string{"program_date_time"},
			File:           "/tmp/reccc/stream.m3u8",
			General: &ffmpeg.OptsOutGeneral{
				Audio:          false,
				VCodec:         "libx264",
				FPS:            15,
				Gop:            15,
				KeyintMin:      15,
				ForceKeyFrames: "expr:gte(t,n_forced*1)",
				Tune:           "zerolatency",
				Preset:         ffmpeg.PresetUltraFast,
			},
		}
	)

	w, err := ffmpeg.NewWorker("test", input1, input2, output)
	if err != nil {
		t.Fatal(err)
	}

	if err := w.Run(); err != nil {
		t.Fatal(err)
	}
}
