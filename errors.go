package bip39

import "errors"

// Error list
var (
	ErrWordLen           = errors.New("Invalid mnemonic list length")
	ErrEntropyLen        = errors.New("Invalid entropy length")
	ErrChecksumIncorrect = errors.New("checksum incorrect")
)
