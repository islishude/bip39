// Package bip39 is the Golang implementation of the BIP39 spec.
package bip39

import (
	"crypto/rand"
	"crypto/sha512"
	"io"

	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/text/unicode/norm"
)

// cryptoRander is a test stub for NewMnemonic func
var cryptoRander = rand.Reader

// NewMnemonicByEntropy create new mnemonic by entropy provided
func NewMnemonicByEntropy(entropy []byte, lang Language) (string, error) {
	entLen := len(entropy)
	// 128 <= ENT <= 256
	if entLen < 16 || entLen > 32 || entLen%4 != 0 {
		return "", ErrEntropyLen
	}
	return fromEntropy(entropy, entLen/4*3, lang), nil
}

// NewMnemonic creates new mnemonic by words length for language provided
func NewMnemonic(length int, lang Language) (string, error) {
	// word length should be 12 | 15 | 18 | 21 | 24
	if length < 12 || length > 24 || length%3 != 0 {
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
	entropy := make([]byte, length+length/3)
	if _, err := io.ReadFull(cryptoRander, entropy); err != nil {
		return "", err
	}

	return fromEntropy(entropy, length, lang), nil
}

// MnemonicToSeed creates 64 bytes seed by pbkdf
// passphrace is optional,it can be empty string
func MnemonicToSeed(mnemonic, passphrase string) []byte {
	password := []byte(norm.NFKD.String(mnemonic))
	salt := []byte(norm.NFKD.String("mnemonic" + passphrase))
	return pbkdf2.Key(password, salt, 2048, 64, sha512.New)
}
