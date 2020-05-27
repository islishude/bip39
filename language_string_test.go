package bip39

import "testing"

func TestLanguage_String(t *testing.T) {
	tests := []struct {
		name string
		i    Language
		want string
	}{
		{"ChineseSimplified", ChineseSimplified, "ChineseSimplified"},
		{"ChineseTraditional", ChineseTraditional, "ChineseTraditional"},
		{"English", English, "English"},
		{"French", French, "French"},
		{"Italian", Italian, "Italian"},
		{"Japanese", Japanese, "Japanese"},
		{"Korean", Korean, "Korean"},
		{"Spanish", Spanish, "Spanish"},
		{"Czech", Czech, "Czech"},
		{"Unknown", 10000, "Language(10000)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.String(); got != tt.want {
				t.Errorf("Language.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
