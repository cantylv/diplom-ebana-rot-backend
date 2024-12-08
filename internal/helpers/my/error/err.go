package myerrors

import "errors"

// outer errors
var (
	ErrInternal               = errors.New("internal server error, please try again later")
	ErrBadPayload             = errors.New("server received bad data")
	ErrBadUUID                = errors.New("id of song must be in uuid format")
	ErrNotExist               = errors.New("target resource does not exist")
	ErrNoRequestIdInContext   = errors.New("no request_id in context")
	ErrInvalidSongLimit       = errors.New("parameter limit must be non negative number")
	ErrInvalidSongOffset      = errors.New("parameter offset must be non negative number")
	ErrInvalidFromReleaseDate = errors.New("parameter from_release_date must be in format DD-MM-YYY")
	ErrInvalidSongID          = errors.New("parameter ids must be in uuid format")
	ErrInvalidToReleaseDate   = errors.New("parameter to_release_date must be in format DD-MM-YYY")
)
