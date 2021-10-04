package ffmpeg

import (
	"strconv"
	"strings"
)

// OptsOutHLS - options for HLS output file.
// See: https://ffmpeg.org/ffmpeg-formats.html#toc-Options-6
type OptsOutHLS struct {
	HLSInitTime          int      // option -hls_init_time
	HLSTime              int      // option -hls_time
	HLSListSize          int      // option -hls_list_size
	HLSDeleteThreshold   int      // option -hls_delete_threshold
	HLSTsOptions         []string // option -hls_ts_options
	HLSWrap              int      // option -hls_wrap
	HLSStartNumberSource string   // option -hls_start_number_source
	StartNumber          int      // option -start_number
	HLSAllowCache        bool     // option -hls_allow_cache
	HLSBaseURL           string   // option -hls_base_url
	HLSSegmentFilename   string   // option -hls_segment_filename
	UseLocaltime         bool     // option -use_localtime
	Strftime             bool     // option -strftime
	UseLocaltimeMkdir    bool     // option -use_localtime_mkdir
	StrftimeMkdir        bool     // option -strftime_mkdir
	HLSKeyInfoFile       string   // option -hls_key_info_file
	HLSEnc               string   // option -hls_enc
	HLSEncKey            string   // option -hls_enc_key
	HLSEncKeyURL         string   // option -hls_enc_key_url
	HLSEncIv             string   // option -hls_enc_iv
	HLSSegmentType       string   // option -hls_segment_type
	HLSFMP4InitFilename  string   // option -hls_fmp4_init_filename
	HLSFMP4InitResend    string   // option -hls_fmp4_init_resend
	HLSFlags             []string // option -hls_flags
	HLSPlaylistType      string   // option -hls_segment_type
	Method               string   // option -method
	HTTPUserAgent        string   // option -http_user_agent
	VarStreamMap         string   // option -var_stream_map
	CCStreamMap          string   // option -cc_stream_map
	MasterPlName         string   // option -master_pl_name
	MasterPlPublishRate  int      // option -master_pl_publish_rate
	HTTPPersistent       string   // option -http_persistent
	Timeout              int      // option -timeout
	IgnoreIoErrors       bool     // option -ignore_io_errors
	Headers              string   // option -header

	File string // output file

	General *OptsOutGeneral // general output options
}

func (o OptsOutHLS) String() string {
	var str string

	if o.General == nil {
		str += DefaultOptionsOGeneral.String()
	} else {
		str += o.General.String()
	}

	str += "-f hls "

	if o.HLSInitTime != 0 {
		str += "-hls_init_time " + strconv.Itoa(o.HLSInitTime) + " "
	}

	if o.HLSTime != 0 {
		str += "-hls_time " + strconv.Itoa(o.HLSTime) + " "
	}

	if o.HLSListSize != -1 {
		str += "-hls_list_size " + strconv.Itoa(o.HLSListSize) + " "
	}

	if o.HLSDeleteThreshold != 0 {
		str += "-hls_delete_threshold " + strconv.Itoa(o.HLSDeleteThreshold) + " "
	}

	if len(o.HLSTsOptions) > 0 {
		str += "-hls_ts_options " + strings.Join(o.HLSTsOptions, "+") + " "
	}

	if o.HLSWrap != 0 {
		str += "-hls_wrap " + strconv.Itoa(o.HLSWrap) + " "
	}

	if o.HLSStartNumberSource != "" {
		str += "-hls_start_number_source " + o.HLSStartNumberSource + " "
	}

	if o.StartNumber != 0 {
		str += "-start_number " + strconv.Itoa(o.StartNumber) + " "
	}

	if o.HLSAllowCache {
		str += "-hls_allow_cache 1 "
	}

	if o.HLSBaseURL != "" {
		str += "-hls_base_url " + o.HLSBaseURL + " "
	}

	if o.HLSSegmentFilename != "" {
		str += "-hls_segment_filename " + o.HLSSegmentFilename + " "
	}

	if o.UseLocaltime {
		str += "-use_localtime "
	}

	if o.Strftime {
		str += "-strftime "
	}

	if o.UseLocaltimeMkdir {
		str += "-use_localtime_mkdir "
	}

	if o.StrftimeMkdir {
		str += "-strftime_mkdir "
	}

	if o.HLSKeyInfoFile != "" {
		str += "-hls_key_info_file " + o.HLSKeyInfoFile + " "
	}

	if o.HLSEnc != "" {
		str += "-hls_enc " + o.HLSEnc + " "
	}

	if o.HLSEncKey != "" {
		str += "-hls_enc_key " + o.HLSEncKey + " "
	}

	if o.HLSEncKeyURL != "" {
		str += "-hls_enc_key_url " + o.HLSEncKeyURL + " "
	}

	if o.HLSEncIv != "" {
		str += "-hls_enc_iv " + o.HLSEncIv + " "
	}

	if o.HLSSegmentType != "" {
		str += "-hls_segment_type " + o.HLSSegmentType + " "
	}

	if o.HLSFMP4InitFilename != "" {
		str += "-hls_fmp4_init_filename " + o.HLSFMP4InitFilename + " "
	}

	if o.HLSFMP4InitResend != "" {
		str += "-hls_fmp4_init_resend " + o.HLSFMP4InitResend + " "
	}

	if len(o.HLSFlags) > 0 {
		str += "-hls_flags " + strings.Join(o.HLSFlags, "+") + " "
	}

	if o.HLSPlaylistType != "" {
		str += "-hls_playlist_type " + o.HLSPlaylistType + " "
	}

	if o.Method != "" {
		str += "-method " + o.Method + " "
	}

	if o.HTTPUserAgent != "" {
		str += "-http_user_agent " + o.HTTPUserAgent + " "
	}

	if o.VarStreamMap != "" {
		str += "-var_stream_map " + o.VarStreamMap + " "
	}

	if o.CCStreamMap != "" {
		str += "cc_stream_map " + o.CCStreamMap + " "
	}

	if o.MasterPlName != "" {
		str += "-master_pl_name " + o.MasterPlName + " "
	}

	if o.MasterPlPublishRate > 0 {
		str += "-master_pl_publish_rate " + strconv.Itoa(o.MasterPlPublishRate) + " "
	}

	if o.HTTPPersistent != "" {
		str += "-http_persistent " + o.HTTPPersistent + " "
	}

	if o.Timeout != 0 {
		str += "-timeout " + strconv.Itoa(o.Timeout) + " "
	}

	if o.IgnoreIoErrors {
		str += "-ignore_io_errors "
	}

	if o.Headers != "" {
		str += "-headers " + o.Headers + " "
	}

	str += o.File

	return str
}
