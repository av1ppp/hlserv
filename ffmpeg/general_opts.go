package ffmpeg

import (
	"strconv"
	"strings"
)

type (
	// OptsInGeneral - general input 'file' options.
	OptsInGeneral struct {
		// See: https://www.ffmpeg.org/ffmpeg.html#Generic-options
		LogLevel []string // option -loglevel

		// See: https://ffmpeg.org/ffmpeg-formats.html#toc-Format-Options
		Analyzeduration int      // option -analyzeduration
		Probesize       int      // option -probesize
		MaxProbePackets int      // option -max_probe_packets
		Packetsize      int      // option -probesize
		FFlags          []string // option -fflags

		// See: https://ffmpeg.org/ffmpeg.html#toc-Audio-Options
		Audio bool // option -an

		Framerate int    // option -framerate
		VideoSize string // option -video_size
		Start     string // option -ss
	}

	// OptionsOGeneral - general output 'file' options.
	//  See: https://www.ffmpeg.org/ffmpeg-scaler.html#toc-Scaler-Options
	OptsOutGeneral struct {
		VCodec         string   // option -c:v
		VBitrate       string   // option -b:v
		VFilter        string   // option -vf
		Bufsize        string   // option -bufisze
		Audio          bool     // option -an
		Scaling        string   // option -s
		SwsFlags       []string // option -sws_flags
		Flags          []string // option -flags
		Preset         string   // option -preset
		Tune           string   // option -tune
		FPS            int      // option -r
		Gop            int      // option -g
		Vsync          string   // option -vsync
		ForceKeyFrames string   // option -force_key_frames
		KeyintMin      int      // option -keyint_min
		VProfile       string   // option -profile
		Level          string   // option -level
		MovFlags       []string // option -movflags
		Start          string   // option -ss
		CRF            int      // option -crf
		Time           string   // option -t
	}
)

var DefaultOptionsIGeneral = OptsInGeneral{
	LogLevel:        []string{},
	Analyzeduration: 10000000,
	Probesize:       8000000,
}

var DefaultOptionsOGeneral = OptsOutGeneral{
	VCodec: "copy",
}

func (o OptsInGeneral) String() string {
	var str string

	if len(o.FFlags) > 0 {
		str += "-fflags " + strings.Join(o.FFlags, "+") + " "
	}

	if len(o.LogLevel) > 0 {
		str += "-loglevel " + strings.Join(o.LogLevel, "+") + " "
	}

	if o.Start != "" {
		str += "-ss " + o.Start + " "
	}

	if o.Analyzeduration > 0 {
		str += "-analyzeduration " + strconv.Itoa(o.Analyzeduration) + " "
	}

	if o.Probesize > 0 {
		str += "-probesize " + strconv.Itoa(o.Probesize) + " "
	}

	if o.MaxProbePackets > 0 {
		str += "-max_probe_packets " + strconv.Itoa(o.MaxProbePackets) + " "
	}

	if o.Packetsize > 0 {
		str += "-packetsize " + strconv.Itoa(o.Packetsize) + " "
	}

	if !o.Audio {
		str += "-an "
	}

	if o.Framerate > 0 {
		str += "-framerate " + strconv.Itoa(o.Framerate) + " "
	}

	if o.VideoSize != "" {
		str += "-video_size " + o.VideoSize + " "
	}

	return str
}

func (o OptsOutGeneral) String() string {
	var str string

	str += "-y "

	if o.Start != "" {
		str += "-ss " + o.Start + " "
	}

	if !o.Audio {
		str += "-an "
	}

	if o.VCodec != "" {
		str += "-c:v " + o.VCodec + " "
	}

	if o.VBitrate != "" {
		str += "-b:v " + o.VBitrate + " -maxrate " + o.VBitrate + " "
	}

	if o.Bufsize != "" {
		str += "-bufsize " + o.Bufsize + " "
	}

	if o.Scaling != "" {
		str += "-s " + o.Scaling + " "
	}

	if o.VFilter != "" {
		str += "-vf " + o.VFilter + " "
	}

	if len(o.SwsFlags) > 0 {
		str += "-sws_flags " + strings.Join(o.SwsFlags, "+") + " "
	}

	if len(o.Flags) > 0 {
		str += "-flags " + strings.Join(o.Flags, "+") + " "
	}

	if o.Preset != "" {
		str += "-preset " + o.Preset + " "
	}

	if o.Tune != "" {
		str += "-tune " + o.Tune + " "
	}

	if o.FPS > 0 {
		str += "-r " + strconv.Itoa(o.FPS) + " "
	}

	if o.Gop > 0 {
		str += "-g " + strconv.Itoa(o.Gop) + " "
	}

	if o.KeyintMin > 0 {
		str += "-keyint_min " + strconv.Itoa(o.KeyintMin) + " "
	}

	if o.CRF > 0 {
		str += "-crf " + strconv.Itoa(o.CRF) + " "
	}

	if o.Time != "" {
		str += "-t " + o.Time + " "
	}

	if o.Vsync != "" {
		str += "-vsync " + o.Vsync + " "
	}

	if o.ForceKeyFrames != "" {
		str += "-force_key_frames " + o.ForceKeyFrames + " "
	}

	if o.VProfile != "" {
		str += "-profile:v " + o.VProfile + " "
	}

	if o.Level != "" {
		str += "-level " + o.Level + " "
	}

	if len(o.MovFlags) > 0 {
		str += "-movflags " + strings.Join(o.MovFlags, "+") + " "
	}

	return str
}
