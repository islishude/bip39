package bip39

import (
	"math"
	"reflect"
	"testing"

	"github.com/islishude/bip39/internal/wordlist"
)

func TestLanguage_List(t *testing.T) {
	tests := []struct {
		name string
		lan  Language
		want []string
	}{
		{
			name: "ChineseSimplified",
			lan:  ChineseSimplified,
			want: wordlist.ChineseSimplified,
		},
		{
			name: "ChineseTraditional",
			lan:  ChineseTraditional,
			want: wordlist.ChineseTraditional,
		},
		{
			name: "English",
			lan:  English,
			want: wordlist.English,
		},
		{
			name: "French",
			lan:  French,
			want: wordlist.French,
		},
		{
			name: "Italian",
			lan:  Italian,
			want: wordlist.Italian,
		},
		{
			name: "Japanese",
			lan:  Japanese,
			want: wordlist.Japanese,
		},
		{
			name: "Korean",
			lan:  Korean,
			want: wordlist.Korean,
		},
		{
			name: "Spanish",
			lan:  Spanish,
			want: wordlist.Spanish,
		},
		{
			name: "Czech",
			lan:  Czech,
			want: wordlist.Czech,
		},
		{
			name: "Portuguese",
			lan:  Portuguese,
			want: wordlist.Portuguese,
		},
		{
			name: "Unsupports",
			lan:  100,
			want: wordlist.English,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.lan.list(false); !reflect.DeepEqual(got, tt.want) || len(got) != 2048 {
				t.Errorf("Language.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLanguage_mapping(t *testing.T) {
	tests := []struct {
		name string
		lan  Language
		want *map[string]int64
	}{
		{"ChineseSimplified", ChineseSimplified, &chineseSimplifiedMapping},
		{"ChineseTraditional", ChineseTraditional, &chineseTraditionalMapping},
		{"English", English, &englishMapping},
		{"Italian", Italian, &italianMapping},
		{"Japanese", Japanese, &japaneseMapping},
		{"Spanish", Spanish, &spanishMapping},
		{"Korean", Korean, &koreanMapping},
		{"Czech", Czech, &czechMapping},
		{"Portuguese", Portuguese, &portugueseMapping},
		{"Unknown", 100, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.lan.mapping()
			if tt.name != "Unknown" && !reflect.DeepEqual(got, *tt.want) {
				t.Errorf("Language.mapping() = %v, want %v", got, tt.want)
			}
			if tt.name != "Unknown" && len(got) != 2048 {
				t.Errorf("Language.mapping() wants 2048 elements but got %d", len(got))
			}
		})
	}
}

func TestLanguage_Iter(t *testing.T) {
	const testMnemonic = "test test test test test test test test test test test "
	for _, word := range wordlist.English {
		tmp := testMnemonic + word
		if IsMnemonicValid(tmp, English) {
			break
		}
	}
}

func TestLanguage_Words(t *testing.T) {
	for _, lan := range []Language{ChineseSimplified, ChineseTraditional, English, French, Italian, Japanese, Korean, Spanish, Czech, Portuguese, math.MaxInt} {
		words := lan.Words()
		if len(words) != 2048 || cap(words) != 2048 {
			t.Errorf("Language.Words() wants 2048 elements but got %d", len(words))
		}
		if lan == English || lan == math.MaxInt {
			if !reflect.DeepEqual(words, wordlist.English) {
				t.Errorf("Language.Words() = %v, want %v", words, wordlist.English)
			}
			words[0] = "nonexistentword"
			if wordlist.English[0] == words[0] {
				t.Errorf("Language.Words() should return a copy of the word list, but it seems to be a reference to the original list")
			}
		}
	}
}
