package bip39

import (
	"strings"
	"testing"
	"unicode"
)

func ToSnakeCaseManual(s string) string {
	var res strings.Builder
	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) {
			res.WriteRune('_')
		}
		res.WriteRune(unicode.ToLower(r))
	}
	return res.String()
}

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
		{"Portuguese", Portuguese, "Portuguese"},
		{"Deutsch", Deutsch, "Deutsch"},
		{"Esperanto", Esperanto, "Esperanto"},
		{"Greek", Greek, "Greek"},
		{"Hinidi", Hindi, "Hindi"},
		{"Latin", Latin, "Latin"},
		{"Russian", Russian, "Russian"},
		{"Unknown", 10000, "Language(10000)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.String(); got != tt.want {
				t.Errorf("Language.String() = %v, want %v", got, tt.want)
			}
			langName := ToSnakeCaseManual(tt.i.String())
			if !tt.i.Valid() {
				got, ok := LanguageByName(langName)
				if ok {
					t.Errorf("LanguageByName(%v) = %v, want invalid", langName, got)
				}
				return
			}
			if got, ok := LanguageByName(langName); !ok || got != tt.i {
				t.Errorf("LanguageByName(%v) = %v, want %v", langName, got, tt.i)
			}
		})
	}
}
