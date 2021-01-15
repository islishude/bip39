package bip39

import (
	"sync"

	"github.com/islishude/bip39/internal/wordlist"
)

//go:generate stringer -type=Language

// Language is bip39 word lang type
type Language int

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
	Czech
	Portuguese
)

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
	case Czech:
		return wordlist.Czech
	case Portuguese:
		return wordlist.Portuguese
	default:
		return wordlist.English
	}
}

var (
	chineseSimplifiedOnce  sync.Once
	chineseTraditionalOnce sync.Once
	englishOnce            sync.Once
	frenchOnce             sync.Once
	italianOnce            sync.Once
	japaneseOnce           sync.Once
	koreanOnce             sync.Once
	spanishOnce            sync.Once
	czechOnce              sync.Once
	portugueseOnce         sync.Once
)

// Words Mapping
var (
	chineseSimplifiedMapping  = make(map[string]int64)
	chineseTraditionalMapping = make(map[string]int64)
	englishMapping            = make(map[string]int64)
	frenchMapping             = make(map[string]int64)
	italianMapping            = make(map[string]int64)
	japaneseMapping           = make(map[string]int64)
	koreanMapping             = make(map[string]int64)
	spanishMapping            = make(map[string]int64)
	czechMapping              = make(map[string]int64)
	portugueseMapping         = make(map[string]int64)
)

// mapping returns word index mapping
func (lan Language) mapping() map[string]int64 {
	switch lan {
	case ChineseSimplified:
		chineseSimplifiedOnce.Do(func() {
			for idx, word := range wordlist.ChineseSimplified {
				chineseSimplifiedMapping[word] = int64(idx)
			}
		})
		return chineseSimplifiedMapping
	case ChineseTraditional:
		chineseTraditionalOnce.Do(func() {
			for idx, word := range wordlist.ChineseTraditional {
				chineseTraditionalMapping[word] = int64(idx)
			}
		})
		return chineseTraditionalMapping
	case English:
		englishOnce.Do(func() {
			for idx, word := range wordlist.English {
				englishMapping[word] = int64(idx)
			}
		})
		return englishMapping
	case French:
		frenchOnce.Do(func() {
			for idx, word := range wordlist.French {
				frenchMapping[word] = int64(idx)
			}
		})
		return frenchMapping
	case Italian:
		italianOnce.Do(func() {
			for idx, word := range wordlist.Italian {
				italianMapping[word] = int64(idx)
			}
		})
		return italianMapping
	case Japanese:
		japaneseOnce.Do(func() {
			for idx, word := range wordlist.Japanese {
				japaneseMapping[word] = int64(idx)
			}
		})
		return japaneseMapping
	case Spanish:
		spanishOnce.Do(func() {
			for idx, word := range wordlist.Spanish {
				spanishMapping[word] = int64(idx)
			}
		})
		return spanishMapping
	case Korean:
		koreanOnce.Do(func() {
			for idx, word := range wordlist.Korean {
				koreanMapping[word] = int64(idx)
			}
		})
		return koreanMapping
	case Czech:
		czechOnce.Do(func() {
			for idx, word := range wordlist.Czech {
				czechMapping[word] = int64(idx)
			}
		})
		return czechMapping
	case Portuguese:
		portugueseOnce.Do(func() {
			for idx, word := range wordlist.Portuguese {
				portugueseMapping[word] = int64(idx)
			}
		})
		return portugueseMapping
	}
	return nil
}
