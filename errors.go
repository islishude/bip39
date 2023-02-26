package bip39

import "errors"

// Error list
var (
	ErrWordLen           = errors.New("invalid mnemonic list length")
	ErrEntropyLen        = errors.New("invalid entropy length")
	ErrChecksumIncorrect = errors.New("checksum incorrect")
)
