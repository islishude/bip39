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

// Words Mapping
var (
	chineseSimplifiedMapping  = make(map[string]int64, 2048)
	chineseTraditionalMapping = make(map[string]int64, 2048)
	englishMapping            = make(map[string]int64, 2048)
	frenchMapping             = make(map[string]int64, 2048)
	italianMapping            = make(map[string]int64, 2048)
	japaneseMapping           = make(map[string]int64, 2048)
	koreanMapping             = make(map[string]int64, 2048)
	spanishMapping            = make(map[string]int64, 2048)
)

func init() {
	for idx, word := range wordlist.ChineseSimplified {
		chineseSimplifiedMapping[word] = int64(idx)
	}
	for idx, word := range wordlist.ChineseTraditional {
		chineseTraditionalMapping[word] = int64(idx)
	}
	for idx, word := range wordlist.English {
		englishMapping[word] = int64(idx)
	}
	for idx, word := range wordlist.French {
		frenchMapping[word] = int64(idx)
	}
	for idx, word := range wordlist.Italian {
		italianMapping[word] = int64(idx)
	}
	for idx, word := range wordlist.Japanese {
		japaneseMapping[word] = int64(idx)
	}
	for idx, word := range wordlist.Spanish {
		spanishMapping[word] = int64(idx)
	}
	for idx, word := range wordlist.Korean {
		koreanMapping[word] = int64(idx)
	}
}

// list gets word list
func (lan Language) list() []string {
	switch lan {
	case English:
		return wordlist.English
	case ChineseSimplified:
		return wordlist.ChineseSimplified
	case ChineseTraditional:
		return wordlist.ChineseTraditional
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
	default:
		return wordlist.English
	}
}

// mapping returns word index mapping
func (lan Language) mapping() map[string]int64 {
	switch lan {
	case ChineseSimplified:
		return chineseSimplifiedMapping
	case ChineseTraditional:
		return chineseTraditionalMapping
	case English:
		return englishMapping
	case French:
		return frenchMapping
	case Italian:
		return italianMapping
	case Japanese:
		return japaneseMapping
	case Spanish:
		return spanishMapping
	case Korean:
		return koreanMapping
	}
	return nil
}
