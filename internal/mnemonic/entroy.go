package mnemonic

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/islishude/bip39/internal/lang"
	"golang.org/x/text/unicode/norm"
)

// 0b11111111111
var last11BitsMask = big.NewInt(2047)
var shift11BitsMask = big.NewInt(2048)

// FromEntropy creates mnemonic from an entropy
func FromEntropy(entropy []byte, wordLen int, lg lang.Language) string {
	// entropy hash
	hash := sha256.New()
	hash.Write(entropy)
	checksum := hash.Sum(nil)[0:1]

	csInt := new(big.Int).SetBytes(checksum)
	csBitLen := uint(len(entropy) / 4)

	csInt.Div(csInt, big.NewInt(1<<(8-csBitLen)))

	// entropy big int
	entInt := big.NewInt(0).SetBytes(entropy)
	entInt.Lsh(entInt, uint(csBitLen))
	entInt.Add(entInt, csInt)

	wordList := make([]string, wordLen)

	lgList := lg.List()

	for i := wordLen - 1; i >= 0; i-- {
		word := new(big.Int)
		word.And(entInt, last11BitsMask)
		entInt.Div(entInt, shift11BitsMask)
		wordList[i] = lgList[word.Int64()]
	}

	if lg == lang.Japanese {
		return strings.Join(wordList, "\u3000")
	}
	return strings.Join(wordList, "\x20")
}

// ValidMnemonic creates entropy from mnemonic
func ValidMnemonic(mnemonic string, lg lang.Language) error {
	mnemonic = norm.NFKD.String(mnemonic)
	wordList := strings.Split(mnemonic, "\x20")

	wordCount := len(wordList)
	if wordCount%3 != 0 || wordCount < 12 || wordCount > 24 {
		// invalid word list length
		return errors.New("Invalid word list length")
	}

	mapping := lg.Mapping()
	entBig := new(big.Int)
	for _, word := range wordList {
		idx, ok := mapping[word]
		if !ok {
			// not includes the word
			return fmt.Errorf("word `%s` not found in mnemonic map", word)
		}
		// tmp := big.NewInt(idx)
		// tmp.Lsh(tmp, uint(wordCount-index))
		// entBig.Add(entBig, tmp)
		entBig.Mul(entBig, shift11BitsMask)
		entBig.Add(entBig, big.NewInt(int64(idx)))
	}

	entroppyBytes := entBig.Bytes()
	checksum := entroppyBytes[len(entroppyBytes)-1:]
	csInt := new(big.Int).SetBytes(checksum)

	entBytes := entroppyBytes[:len(entroppyBytes)-1]
	hash := sha256.New()
	hash.Write(entBytes)
	sum := new(big.Int).SetBytes(hash.Sum(nil)[0:1])

	if sum.Cmp(csInt) != 0 {
		// checksum faild
		return errors.New("checksum incorrect")
	}
	return nil
}
