package bip39

import (
	"iter"
	"sync"

	"github.com/islishude/bip39/internal/wordlist"
)

const mnemonicWordCount = 2048

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
func (lan Language) list(copy bool) (res []string) {
	if copy {
		res = make([]string, 0, mnemonicWordCount)
	}
	switch lan {
	case English:
		if copy {
			return append(res, wordlist.English...)
		}
		return wordlist.English
	case ChineseSimplified:
		if copy {
			return append(res, wordlist.ChineseSimplified...)
		}
		return wordlist.ChineseSimplified
	case ChineseTraditional:
		if copy {
			return append(res, wordlist.ChineseTraditional...)
		}
		return wordlist.ChineseTraditional
	case French:
		if copy {
			return append(res, wordlist.French...)
		}
		return wordlist.French
	case Italian:
		if copy {
			return append(res, wordlist.Italian...)
		}
		return wordlist.Italian
	case Japanese:
		if copy {
			return append(res, wordlist.Japanese...)
		}
		return wordlist.Japanese
	case Spanish:
		if copy {
			return append(res, wordlist.Spanish...)
		}
		return wordlist.Spanish
	case Korean:
		if copy {
			return append(res, wordlist.Korean...)
		}
		return wordlist.Korean
	case Czech:
		if copy {
			return append(res, wordlist.Czech...)
		}
		return wordlist.Czech
	case Portuguese:
		if copy {
			return append(res, wordlist.Portuguese...)
		}
		return wordlist.Portuguese
	default:
		if copy {
			return append(res, wordlist.English...)
		}
		return wordlist.English
	}
}

// Words gets word list
func (lan Language) Words() []string {
	return lan.list(true)
}

func (lan Language) Iter() iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		for idx, word := range lan.list(false) {
			if !yield(idx, word) {
				break
			}
		}
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
	chineseSimplifiedMapping  map[string]int64
	chineseTraditionalMapping map[string]int64
	englishMapping            map[string]int64
	frenchMapping             map[string]int64
	italianMapping            map[string]int64
	japaneseMapping           map[string]int64
	koreanMapping             map[string]int64
	spanishMapping            map[string]int64
	czechMapping              map[string]int64
	portugueseMapping         map[string]int64
)

// mapping returns word index mapping
func (lan Language) mapping() map[string]int64 {
	switch lan {
	case ChineseSimplified:
		chineseSimplifiedOnce.Do(func() {
			chineseSimplifiedMapping = make(map[string]int64, mnemonicWordCount)
			for idx, word := range wordlist.ChineseSimplified {
				chineseSimplifiedMapping[word] = int64(idx)
			}
		})
		return chineseSimplifiedMapping
	case ChineseTraditional:
		chineseTraditionalOnce.Do(func() {
			chineseTraditionalMapping = make(map[string]int64, mnemonicWordCount)
			for idx, word := range wordlist.ChineseTraditional {
				chineseTraditionalMapping[word] = int64(idx)
			}
		})
		return chineseTraditionalMapping
	case English:
		englishOnce.Do(func() {
			englishMapping = make(map[string]int64, mnemonicWordCount)
			for idx, word := range wordlist.English {
				englishMapping[word] = int64(idx)
			}
		})
		return englishMapping
	case French:
		frenchOnce.Do(func() {
			frenchMapping = make(map[string]int64, mnemonicWordCount)
			for idx, word := range wordlist.French {
				frenchMapping[word] = int64(idx)
			}
		})
		return frenchMapping
	case Italian:
		italianOnce.Do(func() {
			italianMapping = make(map[string]int64, mnemonicWordCount)
			for idx, word := range wordlist.Italian {
				italianMapping[word] = int64(idx)
			}
		})
		return italianMapping
	case Japanese:
		japaneseOnce.Do(func() {
			japaneseMapping = make(map[string]int64, mnemonicWordCount)
			for idx, word := range wordlist.Japanese {
				japaneseMapping[word] = int64(idx)
			}
		})
		return japaneseMapping
	case Spanish:
		spanishOnce.Do(func() {
			spanishMapping = make(map[string]int64, mnemonicWordCount)
			for idx, word := range wordlist.Spanish {
				spanishMapping[word] = int64(idx)
			}
		})
		return spanishMapping
	case Korean:
		koreanOnce.Do(func() {
			koreanMapping = make(map[string]int64, mnemonicWordCount)
			for idx, word := range wordlist.Korean {
				koreanMapping[word] = int64(idx)
			}
		})
		return koreanMapping
	case Czech:
		czechOnce.Do(func() {
			czechMapping = make(map[string]int64, mnemonicWordCount)
			for idx, word := range wordlist.Czech {
				czechMapping[word] = int64(idx)
			}
		})
		return czechMapping
	case Portuguese:
		portugueseOnce.Do(func() {
			portugueseMapping = make(map[string]int64, mnemonicWordCount)
			for idx, word := range wordlist.Portuguese {
				portugueseMapping[word] = int64(idx)
			}
		})
		return portugueseMapping
	}
	return nil
}
