package ascendex

import "errors"

// Common errors in project.
var (
	ErrConnectionClosed     = errors.New("connection closed")
	ErrParsingOrderBookData = errors.New("failed parse order book data")
	ErrInvalidSymbolValue   = errors.New("invalid symbol value")
)
