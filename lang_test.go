package bip39_test

import (
	"reflect"
	"testing"

	"github.com/islishude/bip39"
	"github.com/islishude/bip39/internal/wordlist"
)

func TestLanguage_List(t *testing.T) {
	tests := []struct {
		name string
		lan  bip39.Language
		want []string
	}{
		{
			name: "ChineseSimplified",
			lan:  bip39.ChineseSimplified,
			want: wordlist.ChineseSimplified,
		},
		{
			name: "ChineseTraditional",
			lan:  bip39.ChineseTraditional,
			want: wordlist.ChineseTraditional,
		},
		{
			name: "English",
			lan:  bip39.English,
			want: wordlist.English,
		},
		{
			name: "French",
			lan:  bip39.French,
			want: wordlist.French,
		},
		{
			name: "Italian",
			lan:  bip39.Italian,
			want: wordlist.Italian,
		},
		{
			name: "Japanese",
			lan:  bip39.Japanese,
			want: wordlist.Japanese,
		},
		{
			name: "Korean",
			lan:  bip39.Korean,
			want: wordlist.Korean,
		},
		{
			name: "Spanish",
			lan:  bip39.Spanish,
			want: wordlist.Spanish,
		},
		{
			name: "Unsupports",
			lan:  100,
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.lan.List(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Language.List() = %v, want %v", got, tt.want)
			}
		})
	}
}
