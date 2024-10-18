package sevenzip

import "errors"

type ErrorCode int

const (
	NoError ErrorCode = iota
	ProcessNotFound
	CouldNotExtract
	CouldNotCompress
)

var ErrSevenZipNotFound = errors.New("7zip was not found")
