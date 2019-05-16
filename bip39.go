// Package bip39 is the Golang implementation of the BIP39 spec.
package bip39

import (
	"crypto/rand"
	"crypto/sha512"
	"errors"

	"github.com/islishude/bip39/internal/lang"
	"github.com/islishude/bip39/internal/mnemonic"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/text/unicode/norm"
)

// Error list
var (
	ErrWordLen         = errors.New("Invalid word list length")
	ErrEntropyLen      = errors.New("Invalid entropy length")
	ErrInvalidMnemonic = errors.New("Invalid mnemonic")
)

// NewMnemonicByEntropy generates new mnemonic by entropy provided
func NewMnemonicByEntropy(entropy []byte, lang lang.Language) (string, error) {
	entLen := len(entropy)
	// 128 <= ENT <= 256
	if entLen < 16 || entLen > 32 || entLen%4 != 0 {
		return "", ErrEntropyLen
	}
	return mnemonic.FromEntropy(entropy, entLen/4*3, lang), nil
}

// NewMnemonic generates new mnemonic by words length
func NewMnemonic(wordsLen int, lang lang.Language) (string, error) {
	// word length should be 12 | 15 | 18 | 21 | 24
	if wordsLen < 12 || wordsLen > 24 || wordsLen%3 != 0 {
		return "", ErrWordLen
	}

	/*
		CS = ENT / 32
		MS = (ENT + CS) / 11

		|  ENT  | CS | ENT+CS |  MS  |
		+-------+----+--------+------+
		|  128  |  4 |   132  |  12  |
		|  160  |  5 |   165  |  15  |
		|  192  |  6 |   198  |  18  |
		|  224  |  7 |   231  |  21  |
		|  256  |  8 |   264  |  24  |
	*/
	entropy := make([]byte, wordsLen+wordsLen/3)
	if _, err := rand.Read(entropy); err != nil {
		return "", err
	}

	return mnemonic.FromEntropy(entropy, wordsLen, lang), nil
}

// MnemonicToSeed creates seed by mnemonic.
// param passwd can be empty string
func MnemonicToSeed(mnemonic string, passwd string) ([]byte, error) {
	if mnemonic == "" {
		return nil, ErrInvalidMnemonic
	}
	password := []byte(norm.NFKD.String(mnemonic))
	salt := []byte(norm.NFKD.String("mnemonic" + passwd))
	return pbkdf2.Key(password, salt, 2048, 64, sha512.New), nil
}

// IsMnemonicValid validate menemonic
func IsMnemonicValid(m string, lg lang.Language) bool {
	if err := mnemonic.ValidMnemonic(m, lg); err != nil {
		return false
	}
	return true
}
