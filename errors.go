package hlserv

import "errors"

var (
	ErrUnknownFormat      = errors.New("unknown format")
	ErrInvalidURL         = errors.New("invalid url")
	ErrStreamIsNotDefined = errors.New("stream is not defined")
	ErrNotReady           = errors.New("not ready")
)
