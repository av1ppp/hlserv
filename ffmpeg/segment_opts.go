package ffmpeg

import (
	"strconv"
	"strings"
)

// OptsOutSegment - options for segments format output file.
// See: https://ffmpeg.org/ffmpeg-formats.html#toc-Options-12
type OptsOutSegment struct {
	IncrementTc                  bool              // option -increment_tc
	ReferenceStream              string            // option -reference_stream
	SegmentFormat                string            // option -segment_format
	SegmentFormatOptions         map[string]string // option -segment_format_options
	SegmentList                  string            // option -segment_list
	SegmentListFlags             []string          // option -segment_list_flags
	SegmentListSize              int               // option -segment_list_size
	SegmentListEntryPrefix       string            // option -segment_list_entry_prefix
	SegmentListType              string            // option -segment_list_type
	SegmentTime                  int               // option -segment_tiem
	SegmentAtClockTime           bool              // option -segment_atclocktime
	SegmentClockTimeOffset       int               // option -segment_clocktime_offset
	SegmentClockTimeWrapDuration int               // option -segment_clocktime_wrap_duration
	SegmentTimeDelta             int               // option -segment_time_delta
	SegmentTimes                 []int             // option -segment_times
	SegmentFrames                []int             // option -segment_frames
	SegmentWrap                  int               // option -segment_wrap
	SegmentStartNumber           int               // option -segment_start_number
	Strftime                     bool              // option -strftime
	BreakNonKeyframes            bool              // option -break_non_keyframes
	ResetTimestamps              bool              // option -reset_timestamps
	InitialOffset                int               // option -initial_offset
	WriteEmptySegments           bool              // option -write_empty_segments

	File string // output file

	General *OptsOutGeneral
}

func (o OptsOutSegment) String() string {
	var str string

	if o.General == nil {
		str += DefaultOptionsOGeneral.String()
	} else {
		str += o.General.String()
	}

	str += "-f segment "

	if o.IncrementTc {
		str += "-increment_tc "
	}

	if o.ReferenceStream != "" {
		str += "-reference_stream " + o.ReferenceStream + " "
	}

	if o.SegmentFormat != "" {
		str += "-segment_format " + o.SegmentFormat + " "
	}

	if o.SegmentFormatOptions != nil {
		str += "-segment_format_options "
		for key, value := range o.SegmentFormatOptions {
			str += key + "=" + value + ":"
		}
		str = strings.TrimSuffix(str, ":") + " "
	}

	if o.SegmentList != "" {
		str += "-segment_list " + o.SegmentList + " "
	}

	if len(o.SegmentListFlags) > 0 {
		str += "-segment_list_flags " + strings.Join(o.SegmentListFlags, "+") + " "
	}

	if o.SegmentListSize > 0 {
		str += "-segment_list_size " + strconv.Itoa(o.SegmentListSize) + " "
	}

	if o.SegmentListEntryPrefix != "" {
		str += "-segment_list_entry_prefix " + o.SegmentListEntryPrefix + " "
	}

	if o.SegmentListType != "" {
		str += "-segment_list_type " + o.SegmentListType + " "
	}

	if o.SegmentTime > 0 {
		str += "-segment_time " + strconv.Itoa(o.SegmentTime) + " "
	}

	if o.SegmentAtClockTime {
		str += "-segment_atclocktime 1 "
	}

	if o.SegmentClockTimeOffset > 0 {
		str += "-segment_clocktime_offset " + strconv.Itoa(o.SegmentClockTimeOffset) + " "
	}

	if o.SegmentClockTimeWrapDuration > 0 {
		str += "-segment_clocktime_wrap_duration " + strconv.Itoa(o.SegmentClockTimeWrapDuration) + " "

	}

	if o.SegmentTimeDelta > 0 {
		str += "-segment_time_delta " + strconv.Itoa(o.SegmentTimeDelta) + " "
	}

	if len(o.SegmentTimes) > 0 {
		str += "-segment_times "
		for _, t := range o.SegmentTimes {
			str += strconv.Itoa(t) + ","
		}
		str = strings.TrimSuffix(str, ",") + " "
	}

	if len(o.SegmentFrames) > 0 {
		str += "-segment_frames "
		for _, t := range o.SegmentFrames {
			str += strconv.Itoa(t) + ","
		}
		str = strings.TrimSuffix(str, ",") + " "
	}

	if o.SegmentWrap > 0 {
		str += "-segment_wrap " + strconv.Itoa(o.SegmentWrap) + " "
	}

	if o.SegmentStartNumber > 0 {
		str += "-segment_start_number " + strconv.Itoa(o.SegmentStartNumber) + " "
	}

	if o.Strftime {
		str += "-strftime 1 "
	}

	if o.BreakNonKeyframes {
		str += "-break_non_keyframes 1 "
	}

	if o.ResetTimestamps {
		str += "-reset_timestamps 1 "
	}

	if o.InitialOffset > 0 {
		str += "-initial_offset " + strconv.Itoa(o.InitialOffset) + " "
	}

	if o.WriteEmptySegments {
		str += "-write_empty_segments 1 "
	}

	str += o.File

	return str
}
