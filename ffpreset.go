package hlserv

import (
	"fmt"

	"github.com/av1ppp/hlserv/ffmpeg"
)

func presetMP4(stream *Stream) (ffmpeg.OptsInMp4, ffmpeg.OptsOutHLS) {
	var (
		input = ffmpeg.OptsInMp4{
			File: stream.Config.Source,
			General: &ffmpeg.OptsInGeneral{
				Speed: stream.Config.Speed,
				Start: fmt.Sprintf("%f", stream.Config.OffsetSec),
			},
		}
		output = ffmpeg.OptsOutHLS{
			HLSTime:        4,
			HLSListSize:    0, // unlimit
			HLSSegmentType: "mpegts",
			HLSFlags:       []string{"program_date_time"},
			File:           fmt.Sprintf("%s/%s/stream.m3u8", EndPoint, stream.ID),
			General: &ffmpeg.OptsOutGeneral{
				Audio:          stream.Config.AudioEnabled,
				VCodec:         "libx264",
				FPS:            stream.Config.FPS,
				Gop:            stream.Config.FPS,
				KeyintMin:      stream.Config.FPS,
				ForceKeyFrames: "expr:gte(t,n_forced*1)",
				Tune:           "zerolatency",
				Preset:         stream.Config.Preset,

				VFilter:  formatVFilter(stream.Config.Scale),
				Level:    stream.Config.Level,
				VProfile: stream.Config.Profile,
				VBitrate: stream.Config.VideoBitrate,
				Bufsize:  stream.Config.Bufsize,
			},
		}
	)

	return input, output
}

func formatVFilter(scale string) string {
	if scale == "" {
		return ""
	}

	return "scale=" + scale
}

func presetRTSP(stream *Stream) (ffmpeg.OptsInRTSP, ffmpeg.OptsOutHLS) {
	var (
		input = ffmpeg.OptsInRTSP{
			RTSPTransport:     "udp",
			AllowedMediaTypes: []string{"video", "data"},
			Stimeout:          20_000_000, // 20 sec.
			File:              stream.Config.Source,
		}
		output = ffmpeg.OptsOutHLS{
			HLSTime:        1,
			HLSListSize:    32,
			HLSSegmentType: "mpegts",
			HLSFlags:       []string{"delete_segments", "omit_endlist"},
			File:           fmt.Sprintf("%s/%s/stream.m3u8", EndPoint, stream.ID),
			General: &ffmpeg.OptsOutGeneral{
				Audio:          stream.Config.AudioEnabled,
				VCodec:         "libx264",
				FPS:            stream.Config.FPS,
				Gop:            stream.Config.FPS,
				KeyintMin:      stream.Config.FPS,
				ForceKeyFrames: "expr:gte(t,n_forced*1)",
				Tune:           "zerolatency",
				Preset:         stream.Config.Preset,

				VFilter:  formatVFilter(stream.Config.Scale),
				Level:    stream.Config.Level,
				VProfile: stream.Config.Profile,
				VBitrate: stream.Config.VideoBitrate,
				Bufsize:  stream.Config.Bufsize,
			},
		}
	)

	return input, output
}
