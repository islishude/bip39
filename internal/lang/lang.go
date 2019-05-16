package lang

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
	ChineseSimplifiedMapping  = make(map[string]int64, 2048)
	ChineseTraditionalMapping = make(map[string]int64, 2048)
	EnglishMapping            = make(map[string]int64, 2048)
	FrenchMapping             = make(map[string]int64, 2048)
	ItalianMapping            = make(map[string]int64, 2048)
	JapaneseMapping           = make(map[string]int64, 2048)
	KoreanMapping             = make(map[string]int64, 2048)
	SpanishMapping            = make(map[string]int64, 2048)
)

func init() {
	for idx, word := range wordlist.ChineseSimplified {
		ChineseSimplifiedMapping[word] = int64(idx)
	}
	for idx, word := range wordlist.ChineseTraditional {
		ChineseTraditionalMapping[word] = int64(idx)
	}
	for idx, word := range wordlist.English {
		EnglishMapping[word] = int64(idx)
	}
	for idx, word := range wordlist.French {
		FrenchMapping[word] = int64(idx)
	}
	for idx, word := range wordlist.Italian {
		ItalianMapping[word] = int64(idx)
	}
	for idx, word := range wordlist.Japanese {
		JapaneseMapping[word] = int64(idx)
	}
	for idx, word := range wordlist.Spanish {
		SpanishMapping[word] = int64(idx)
	}
	for idx, word := range wordlist.Korean {
		KoreanMapping[word] = int64(idx)
	}
}

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

// Mapping returns word index mapping
func (lan Language) Mapping() map[string]int64 {
	switch lan {
	case ChineseSimplified:
		return ChineseSimplifiedMapping
	case ChineseTraditional:
		return ChineseTraditionalMapping
	case English:
		return EnglishMapping
	case French:
		return FrenchMapping
	case Italian:
		return ItalianMapping
	case Japanese:
		return JapaneseMapping
	case Spanish:
		return SpanishMapping
	case Korean:
		return KoreanMapping
	}
	return nil
}
