package bip39

import (
	"iter"
	"strings"
	"sync"

	"github.com/islishude/bip39/internal/wordlist"
)

const mnemonicWordCount = 2048

//go:generate stringer -type=Language

// Language is bip39 word lang type
type Language int

// Language list
const (
	English Language = iota
	ChineseSimplified
	ChineseTraditional
	French
	Italian
	Japanese
	Korean
	Spanish
	Czech
	Portuguese
	Deutsch
	Esperanto
	Greek
	Hindi
	Latin
	Russian
)

func LanguageByName(lang string) (Language, bool) {
	switch strings.ToLower(lang) {
	case "english":
		return English, true
	case "chinese_simplified":
		return ChineseSimplified, true
	case "chinese_traditional":
		return ChineseTraditional, true
	case "french":
		return French, true
	case "italian":
		return Italian, true
	case "japanese":
		return Japanese, true
	case "korean":
		return Korean, true
	case "spanish":
		return Spanish, true
	case "czech":
		return Czech, true
	case "portuguese":
		return Portuguese, true
	case "deutsch":
		return Deutsch, true
	case "esperanto":
		return Esperanto, true
	case "greek":
		return Greek, true
	case "hindi":
		return Hindi, true
	case "latin":
		return Latin, true
	case "russian":
		return Russian, true
	default:
		return 0, false
	}
}

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
	case Deutsch:
		if copy {
			return append(res, wordlist.Deutsch...)
		}
		return wordlist.Deutsch
	case Esperanto:
		if copy {
			return append(res, wordlist.Esperanto...)
		}
		return wordlist.Esperanto
	case Greek:
		if copy {
			return append(res, wordlist.Greek...)
		}
		return wordlist.Greek
	case Hindi:
		if copy {
			return append(res, wordlist.Hindi...)
		}
		return wordlist.Hindi
	case Latin:
		if copy {
			return append(res, wordlist.Latin...)
		}
		return wordlist.Latin
	case Russian:
		if copy {
			return append(res, wordlist.Russian...)
		}
		return wordlist.Russian
	}
	return nil
}

// Words gets word list
func (lan Language) Words() []string {
	return lan.list(true)
}

func (lan Language) Valid() bool {
	return lan >= English && lan <= Russian
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
	frenchOnce             sync.Once
	italianOnce            sync.Once
	japaneseOnce           sync.Once
	koreanOnce             sync.Once
	spanishOnce            sync.Once
	czechOnce              sync.Once
	portugueseOnce         sync.Once
	deutschOnce            sync.Once
	esperantoOnce          sync.Once
	greekOnce              sync.Once
	hindiOnce              sync.Once
	latinOnce              sync.Once
	russianOnce            sync.Once
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
	deutschMapping            map[string]int64
	esperantoMapping          map[string]int64
	greekMapping              map[string]int64
	hindiMapping              map[string]int64
	latinMapping              map[string]int64
	russianMapping            map[string]int64
)

func init() {
	englishMapping = make(map[string]int64, mnemonicWordCount)
	for idx, word := range wordlist.English {
		englishMapping[word] = int64(idx)
	}
}

// mapping returns word index mapping
func (lan Language) mapping() map[string]int64 {
	switch lan {
	case English:
		return englishMapping
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
	case Deutsch:
		deutschOnce.Do(func() {
			deutschMapping = make(map[string]int64, mnemonicWordCount)
			for idx, word := range wordlist.Deutsch {
				deutschMapping[word] = int64(idx)
			}
		})
		return deutschMapping
	case Esperanto:
		esperantoOnce.Do(func() {
			esperantoMapping = make(map[string]int64, mnemonicWordCount)
			for idx, word := range wordlist.Esperanto {
				esperantoMapping[word] = int64(idx)
			}
		})
		return esperantoMapping
	case Greek:
		greekOnce.Do(func() {
			greekMapping = make(map[string]int64, mnemonicWordCount)
			for idx, word := range wordlist.Greek {
				greekMapping[word] = int64(idx)
			}
		})
		return greekMapping
	case Hindi:
		hindiOnce.Do(func() {
			hindiMapping = make(map[string]int64, mnemonicWordCount)
			for idx, word := range wordlist.Hindi {
				hindiMapping[word] = int64(idx)
			}
		})
		return hindiMapping
	case Latin:
		latinOnce.Do(func() {
			latinMapping = make(map[string]int64, mnemonicWordCount)
			for idx, word := range wordlist.Latin {
				latinMapping[word] = int64(idx)
			}
		})
		return latinMapping
	case Russian:
		russianOnce.Do(func() {
			russianMapping = make(map[string]int64, mnemonicWordCount)
			for idx, word := range wordlist.Russian {
				russianMapping[word] = int64(idx)
			}
		})
		return russianMapping
	}
	return nil
}
