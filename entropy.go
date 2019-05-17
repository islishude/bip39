package bip39

import (
	"crypto/sha256"
	"math/big"
	"strings"
)

// 0b11111111111
var last11BitsMask = big.NewInt(2047)
var first11BitsMask = big.NewInt(2048)

// fromEntropy creates mnemonic from an entropy
func fromEntropy(entropy []byte, wordLen int, lg Language) string {
	// entropy hash
	hash := sha256.New()
	hash.Write(entropy)
	checksum := hash.Sum(nil)[0:1]

	csInt := new(big.Int).SetBytes(checksum)
	csBitLen := uint(len(entropy) / 4)
	csInt.Div(csInt, big.NewInt(1<<(8-csBitLen)))

	// entropy big int
	entInt := new(big.Int).SetBytes(entropy)
	entInt.Lsh(entInt, uint(csBitLen))
	entInt.Add(entInt, csInt)

	wordList := make([]string, wordLen)
	lgList := lg.List()
	wordIdx := new(big.Int)
	for i := wordLen - 1; i >= 0; i-- {
		wordIdx.And(entInt, last11BitsMask)
		entInt.Div(entInt, first11BitsMask)
		wordList[i] = lgList[wordIdx.Int64()]
	}

	if lg == Japanese {
		return strings.Join(wordList, "\u3000")
	}
	return strings.Join(wordList, "\x20")
}
