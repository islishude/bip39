package bip39

import (
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
			if got := tt.lan.list(); !reflect.DeepEqual(got, tt.want) || len(got) != 2048 {
				t.Errorf("Language.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLanguage_mapping(t *testing.T) {
	tests := []struct {
		name string
		lan  Language
		want map[string]int64
	}{
		{"ChineseSimplified", ChineseSimplified, chineseSimplifiedMapping},
		{"ChineseTraditional", ChineseTraditional, chineseTraditionalMapping},
		{"English", English, englishMapping},
		{"Italian", Italian, italianMapping},
		{"Japanese", Japanese, japaneseMapping},
		{"Spanish", Spanish, spanishMapping},
		{"Korean", Korean, koreanMapping},
		{"Czech", Czech, czechMapping},
		{"Portuguese", Portuguese, portugueseMapping},
		{"Unknown", 100, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.lan.mapping()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Language.mapping() = %v, want %v", got, tt.want)
			}
			if got != nil && len(got) != 2048 {
				t.Errorf("Language.mapping() wants 2048 elements but got %d", len(got))
			}
		})
	}
}
