package hlserv

import "errors"

var (
	ErrUnknownFormat = errors.New("unknown format")
	ErrInvalidURL    = errors.New("invalid url")
)
