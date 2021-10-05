package hlserv

import (
	"fmt"
	"time"

	"github.com/av1ppp/hlserv/ffmpeg"
)

func mp4ToHLS(streamID string, src string, offset time.Duration) (ffmpeg.OptsInMp4, ffmpeg.OptsOutHLS) {
	var (
		input = ffmpeg.OptsInMp4{
			File: src,
			General: &ffmpeg.OptsInGeneral{
				LogLevel: []string{},
				Start:    fmt.Sprintf("%f", offset.Seconds()),
			},
		}
		output = ffmpeg.OptsOutHLS{
			HLSTime:        4,
			HLSListSize:    0, // unlimit
			HLSSegmentType: "mpegts",
			HLSFlags:       []string{"program_date_time"},
			File:           fmt.Sprintf("%s/%s/stream.m3u8", EndPoint, streamID),
			General: &ffmpeg.OptsOutGeneral{
				Audio:          false,
				VCodec:         "libx264",
				FPS:            15,
				Gop:            15,
				KeyintMin:      15,
				ForceKeyFrames: "expr:gte(t,n_forced*1)",
				Tune:           "zerolatency",
				Preset:         "ultrafast",
				CRF:            25,

				Level:    "5.0",
				VProfile: "main",
				VBitrate: "5000k",
				Bufsize:  "2500k",
			},
		}
	)

	// switch quality {
	// case QualityHigh:
	// 	output.General.Level = "5.0"
	// 	output.General.VProfile = "main"
	// 	output.General.VBitrate = "5000k"
	// 	output.General.Bufsize = "2500k"

	// case QualityMedium:
	// 	output.General.Level = "3.0"
	// 	output.General.VProfile = "main"
	// 	output.General.VBitrate = "1600k"
	// 	output.General.Bufsize = "800k"
	// 	output.General.VFilter = "scale=-2:480"

	// case QualityLow:
	// 	output.General.Level = "1.1"
	// 	output.General.VProfile = "baseline"
	// 	output.General.VBitrate = "1000k"
	// 	output.General.Bufsize = "500k"
	// 	output.General.VFilter = "scale=-2:240"
	// }

	return input, output
}

func formatVFilter(scale string) string {
	if scale == "" {
		return ""
	}

	return "scale=" + scale
}

func rtspToHLS(conf *StreamConfig) (ffmpeg.OptsInRTSP, ffmpeg.OptsOutHLS) {
	// Creating ffmpeg worker
	var (
		input = ffmpeg.OptsInRTSP{
			RTSPTransport:     "udp",
			AllowedMediaTypes: []string{"video", "data"},
			Stimeout:          20_000_000, // 20 sec.
			File:              conf.Source,
		}
		output = ffmpeg.OptsOutHLS{
			HLSTime:        1,
			HLSListSize:    32,
			HLSSegmentType: "mpegts",
			HLSFlags:       []string{"delete_segments", "omit_endlist"},
			File:           fmt.Sprintf("%s/%s/stream.m3u8", EndPoint, conf.id),
			General: &ffmpeg.OptsOutGeneral{
				Audio:          conf.AudioEnabled,
				VCodec:         "libx264",
				FPS:            conf.FPS,
				Gop:            conf.FPS,
				KeyintMin:      conf.FPS,
				ForceKeyFrames: "expr:gte(t,n_forced*1)",
				Tune:           "zerolatency",
				Preset:         ffmpeg.FormatPreset(conf.Preset),
				CRF:            conf.CRF,

				VFilter:  formatVFilter(conf.Scale),
				Level:    conf.Level,
				VProfile: conf.Profile,
				VBitrate: conf.VideoBitrate,
				Bufsize:  conf.Bufsize,
			},
		}
	)

	return input, output
}
