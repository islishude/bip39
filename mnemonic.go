package bip39

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"strings"

	"golang.org/x/text/unicode/norm"
)

// IsMnemonicValid validate menemonic
func IsMnemonicValid(m string, lg Language) bool {
	return CheckMnemonic(m, lg) == nil
}

// CheckMnemonic creates entropy from mnemonic
func CheckMnemonic(mnemonic string, lg Language) error {
	mnemonic = norm.NFKD.String(mnemonic)
	wordList := strings.Split(mnemonic, "\x20")

	wordCount := len(wordList)
	// invalid word list length
	if wordCount%3 != 0 || wordCount < 12 || wordCount > 24 {
		return ErrWordLen
	}

	entBig := new(big.Int)
	mapping := lg.mapping()
	for wordIdx, word := range wordList {
		idx, ok := mapping[word]
		// not includes the word
		if !ok {
			return fmt.Errorf("word `%s` at `%d` not found in mnemonic mapping", word, wordIdx)
		}

		partBig := big.NewInt(int64(idx))
		partBig.Lsh(partBig, uint(wordCount-wordIdx-1)*11)
		entBig.Add(entBig, partBig)
	}

	var shift int64 = 1 << uint(wordCount/3)
	// get checksum
	csBig := new(big.Int).And(entBig, big.NewInt(shift-1))

	// get real entropy
	entBytes := entBig.Quo(entBig, big.NewInt(shift)).Bytes()
	// get checksum from real entropy
	hash := sha256.New()
	_, _ = hash.Write(entBytes)
	sum := new(big.Int).SetBytes(hash.Sum(nil)[0:1])
	sum.Quo(sum, big.NewInt(1<<(8-uint(wordCount/3))))

	// compare checksum
	if sum.Cmp(csBig) != 0 {
		return ErrChecksumIncorrect
	}
	return nil
}
