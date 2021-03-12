package errutil

import (
	"fmt"
)

type ErrReason int

const (
	ErrReadClientSecret = iota
	ErrReadAuthCode
	ErrRetrieveJwtToken
	ErrSaveJwtToken
	ErrParseClientSecret
	ErrRetrieveSheetsClient
	ErrParsingRangeResponse
)

// Error contains error reason and the actual error if any
// satisfies golang's error interface
type Error struct {
	reason ErrReason
	err    error
}

func New(reason ErrReason, err error) *Error {
	return &Error{reason: reason, err: err}
}

func (err Error) Error() string {
	if err.err != nil {
		return fmt.Sprintf("%s (%v)", err.desc(), err.err)
	}
	return err.desc()
}

func (err Error) desc() string {
	switch err.reason {
	case ErrReadClientSecret:
		return "unable to read client secret"
	case ErrReadAuthCode:
		return "unable to read auth code from command line"
	case ErrRetrieveJwtToken:
		return "unable to retrieve jwt token for sheets from the web"
	case ErrParseClientSecret:
		return "unable to parse client secret"
	case ErrRetrieveSheetsClient:
		return "unable to retrieve google sheets client"
	case ErrSaveJwtToken:
		return "unable to save jwt token"
	case ErrParsingRangeResponse:
		return "error parsing range response"
	default:
		return "unknown error"
	}
}
