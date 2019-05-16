package bip39

import "github.com/islishude/bip39/internal/wordlist"

// Language is bip39 word lang type
type Language uint8

// Language list
const (
	ChineseSimplified Language = iota
	ChineseTraditional
	English
	French
	Italian
	Japanese
	Korean
	Spanish
)

// List gets word list
func (lan Language) List() []string {
	switch lan {
	case ChineseSimplified:
		return wordlist.ChineseSimplified
	case ChineseTraditional:
		return wordlist.ChineseTraditional
	case English:
		return wordlist.English
	case French:
		return wordlist.French
	case Italian:
		return wordlist.Italian
	case Japanese:
		return wordlist.Japanese
	case Spanish:
		return wordlist.Spanish
	case Korean:
		return wordlist.Korean
	}
	return nil
}
