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
			name: "Deutsch",
			lan:  Deutsch,
			want: wordlist.Deutsch,
		},
		{
			name: "Esperanto",
			lan:  Esperanto,
			want: wordlist.Esperanto,
		},
		{
			name: "Greek",
			lan:  Greek,
			want: wordlist.Greek,
		},
		{
			name: "Hindi",
			lan:  Hindi,
			want: wordlist.Hindi,
		},
		{
			name: "Latin",
			lan:  Latin,
			want: wordlist.Latin,
		},
		{
			name: "Russian",
			lan:  Russian,
			want: wordlist.Russian,
		},
		{
			name: "Unsupports",
			lan:  math.MaxInt,
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.lan.list(false)
			if len(tt.want) == 0 && len(got) != 0 {
				t.Errorf("Language.List() = nil, want %v", tt.want)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
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
		{"Deutsch", Deutsch, &deutschMapping},
		{"Esperanto", Esperanto, &esperantoMapping},
		{"Greek", Greek, &greekMapping},
		{"Hindi", Hindi, &hindiMapping},
		{"Latin", Latin, &latinMapping},
		{"Russian", Russian, &russianMapping},
		{"Unknown", math.MaxInt, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.lan.mapping()
			if tt.name != "Unknown" && !reflect.DeepEqual(got, *tt.want) {
				t.Errorf("Language.mapping() = %v, want %v", got, tt.want)
			}
			if tt.name != "Unknown" && len(got) != mnemonicWordCount {
				t.Errorf("Language.mapping() wants %d elements but got %d", mnemonicWordCount, len(got))
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
	for _, lan := range []Language{
		ChineseSimplified,
		ChineseTraditional,
		English,
		French,
		Italian,
		Japanese,
		Korean,
		Spanish,
		Czech,
		Portuguese,
		Deutsch,
		Esperanto,
		Greek,
		Hindi,
		Latin,
		Russian,
	} {
		words := lan.Words()
		if len(words) != mnemonicWordCount || cap(words) != mnemonicWordCount {
			t.Errorf("Language.Words() wants %d elements but got %d", mnemonicWordCount, len(words))
		}
		if lan == English {
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
