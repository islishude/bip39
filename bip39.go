// Package bip39 is the Golang implementation of the BIP39 spec.
package bip39

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"fmt"
	"strconv"
	"strings"

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
func NewMnemonicByEntropy(entropy []byte, lang Language) (string, error) {
	entLen := len(entropy)
	// 128 <= ENT <= 256
	if entLen < 16 || entLen > 32 || entLen%4 != 0 {
		return "", ErrEntropyLen
	}
	return entropyToMnemonic(entropy, entLen/32*3, lang)
}

// NewMnemonic generates new mnemonic by words length
func NewMnemonic(wordsLen int, lang Language) (string, error) {
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

	return entropyToMnemonic(entropy, wordsLen, lang)
}

func entropyToMnemonic(entropy []byte, wordsLen int, lang Language) (string, error) {
	var binEntBuf strings.Builder
	{
		for _, v := range entropy {
			tmp := fmt.Sprintf("%08b", v)
			binEntBuf.WriteString(tmp)
		}
		hash := sha256.New()
		hash.Write(entropy)
		cs1stByte := hash.Sum(nil)[0]
		csBinStr := fmt.Sprintf("%08b", cs1stByte)
		// len(entropy)/4 = len(entroy) * 8 / 32
		binEntBuf.WriteString(csBinStr[:len(entropy)/4])
	}

	binEnt := binEntBuf.String()

	words := make([]string, 0, wordsLen)
	wordList := lang.List()
	for i := 0; i < len(binEnt); i += 11 {
		idx, err := strconv.ParseInt(binEnt[i:i+11], 2, 32)
		if err != nil {
			return "", err
		}
		words = append(words, wordList[idx])
	}

	if lang == Japanese {
		return strings.Join(words, "\u3000"), nil
	}
	return strings.Join(words, "\x20"), nil
}

// MnemonicToSeed creates seed by mnemonic.
// param passwd can be empty string
func MnemonicToSeed(mnemonic string, passwd string) ([]byte, error) {
	if mnemonic == "" {
		return nil, ErrInvalidMnemonic
	}
	reschan := make(chan []byte)
	go func() {
		defer close(reschan)
		password := []byte(norm.NFKD.String(mnemonic))
		salt := []byte(norm.NFKD.String("mnemonic" + passwd))
		reschan <- pbkdf2.Key(password, salt, 2048, 64, sha512.New)
	}()
	return <-reschan, nil
}

// IsMnemonicValid validate menemonic
func IsMnemonicValid(mnemonic string, lang Language) bool {
	mnemonic = norm.NFKD.String(mnemonic)
	wordList := strings.Split(mnemonic, "\x20")

	wordCount := len(wordList)
	if wordCount%3 != 0 || wordCount < 12 || wordCount > 24 {
		return false
	}

	// record index of word
	wordMapping := make(map[string]int)
	for idx, v := range lang.List() {
		wordMapping[v] = idx
	}

	binEnt := mnemonicToEntropy(wordList, wordMapping)
	if binEnt == "" {
		return false
	}
	binEntLen := len(binEnt)

	entropy := make([]byte, 0, wordCount+wordCount/3)
	for i := 0; i < binEntLen-4; i += 8 {
		b, err := strconv.ParseInt(binEnt[i:i+8], 2, 32)
		if err != nil {
			return false
		}
		entropy = append(entropy, byte(b))
	}

	hash := sha256.New()
	hash.Write(entropy)
	csBitsLen := wordCount / 3
	return binEnt[binEntLen-4:binEntLen] == fmt.Sprintf("%08b", hash.Sum(nil)[0])[:csBitsLen]
}

func mnemonicToEntropy(wordList []string, wordMapping map[string]int) string {
	var entBuf strings.Builder
	for _, v := range wordList {
		idx, has := wordMapping[v]
		if !has {
			return ""
		}
		x := fmt.Sprintf("%08b", idx)
		// padding to length 11 with 0
		if rpt := 11 - len(x); rpt != 0 {
			x = strings.Repeat("0", rpt) + x
		}
		entBuf.WriteString(x)
	}
	return entBuf.String()
}
