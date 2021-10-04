package ffmpeg

// OptsInMp4 - options for mp4 input file.
type OptsInMp4 struct {
	File string // option -i

	// Set general input "file" options.
	General *OptsInGeneral
}

func (o OptsInMp4) String() string {
	var str string

	if o.General == nil {
		str += DefaultOptionsIGeneral.String()
	} else {
		str += o.General.String()
	}

	str += "-i " + o.File

	return str
}

type OptsOutMp4 struct {
	File string

	// Set general input "file" options.
	General *OptsOutGeneral
}

func (o OptsOutMp4) String() string {
	var str string

	if o.General == nil {
		str += DefaultOptionsIGeneral.String()
	} else {
		str += o.General.String()
	}

	str += o.File

	return str
}
