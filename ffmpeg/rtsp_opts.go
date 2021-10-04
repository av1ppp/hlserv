package ffmpeg

import (
	"strconv"
	"strings"
)

// OptsInRTSP - options for RTSP input file.
type OptsInRTSP struct {
	// See: https://www.ffmpeg.org/ffmpeg-protocols.html#toc-rtsp
	InitialPause      bool     // option -initial_pause
	RTSPTransport     string   // option -rtsp_transport
	RTSPFlags         []string // option -rtsp_flags
	AllowedMediaTypes []string // option -allowed_media_types
	MinPort           int      // option -min_port
	MaxPort           int      // option -max_port
	Timeout           int      // option -timeout
	RecorderQueueSize int      // option -recorder_queue_size
	Stimeout          int      // option -stimeout
	UserAgent         string   // option -user-agent
	File              string   // option -i

	// Set general input "file" options.
	General *OptsInGeneral
}

func (o OptsInRTSP) String() string {
	var str string

	if o.General == nil {
		str += DefaultOptionsIGeneral.String()
	} else {
		str += o.General.String()
	}

	str += "-f rtsp "

	if o.InitialPause {
		str += "-initial_pause 1 "
	}

	if o.RTSPTransport != "" {
		str += "-rtsp_transport " + o.RTSPTransport + " "
	}

	if len(o.RTSPFlags) > 0 {
		str += "-rtsp_flags " + strings.Join(o.RTSPFlags, "+") + " "
	}

	if len(o.AllowedMediaTypes) > 0 {
		str += "-allowed_media_types " + strings.Join(o.AllowedMediaTypes, "+") + " "
	}

	if o.MinPort != 0 {
		str += "-min_port " + strconv.Itoa(o.MinPort) + " "
	}

	if o.MaxPort != 0 {
		str += "-max_port " + strconv.Itoa(o.MaxPort) + " "
	}

	if o.Timeout != 0 {
		str += "-timeout " + strconv.Itoa(o.Timeout) + " "
	}

	if o.RecorderQueueSize != 0 {
		str += "-reorder_queue_size " + strconv.Itoa(o.RecorderQueueSize) + " "
	}

	if o.Stimeout != 0 {
		str += "-stimeout " + strconv.Itoa(o.Stimeout) + " "
	}

	if o.UserAgent != "" {
		str += "-user-agent " + o.UserAgent + " "
	}

	str += "-i " + o.File

	return str
}
