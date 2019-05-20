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
		want [2048]string
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
			name: "Unsupports",
			lan:  100,
			want: wordlist.English,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.lan.list(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Language.List() = %v, want %v", got, tt.want)
			}
		})
	}
}
