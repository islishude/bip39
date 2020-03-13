package bip39

import "testing"

func TestIsMnemonicValid(t *testing.T) {
	type args struct {
		mnemonic string
		lang     Language
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "English",
			args: args{
				mnemonic: "check fiscal fit sword unlock rough lottery tool sting pluck bulb random",
				lang:     English,
			},
			want: true,
		},
		{
			name: "Englishx2",
			args: args{
				mnemonic: "rich soon pool legal busy add couch tower goose security raven anger",
				lang:     English,
			},
			want: true,
		},
		{
			name: "EnglishValidLength",
			args: args{
				mnemonic: "rich soon pool legal busy add couch tower goose security raven",
				lang:     English,
			},
			want: false,
		},
		{
			name: "EnglishNoWord",
			args: args{
				mnemonic: "rich soon pool legal busy add couch tower goose security women",
				lang:     English,
			},
			want: false,
		},
		{
			name: "English with invalid word",
			args: args{
				mnemonic: "bip39 english mnemonic test case bip39 english mnemonic test case bip39 english mnemonic test case",
				lang:     English,
			},
			want: false,
		},
		{
			name: "EnglishChecksumError",
			args: args{
				mnemonic: "rich soon pool legal busy add couch tower goose security base",
				lang:     English,
			},
			want: false,
		},
		{
			name: "ChineseSimplified",
			args: args{
				mnemonic: "氮 冠 锋 枪 做 到 容 枯 获 槽 弧 部",
				lang:     ChineseSimplified,
			},
			want: true,
		},
		{
			name: "ChineseTraditional",
			args: args{
				mnemonic: "氮 冠 鋒 槍 做 到 容 枯 獲 槽 弧 部",
				lang:     ChineseTraditional,
			},
			want: true,
		},
		{
			name: "Japanese",
			args: args{
				mnemonic: "ねほりはほり　ひらがな　とさか　そつう　おうじ　あてな　きくらげ　みもと　してつ　ぱそこん　にってい　いこつ",
				lang:     Japanese,
			},
			want: true,
		},
		{
			name: "Spanish",
			args: args{
				mnemonic: "posible ruptura ozono ligero bobina acto chuleta tetera gol realidad pez alerta",
				lang:     Spanish,
			},
			want: true,
		},
		{
			name: "French",
			args: args{
				mnemonic: "pieuvre revivre nuptial implorer blinder accroche chute syntaxe félin promener parcelle aimable",
				lang:     French,
			},
			want: true,
		},
		{
			name: "Italian",
			args: args{
				mnemonic: "risultato siccome prenotare mimosa bosco adottare continuo tifare ignaro sbloccato residente alticcio",
				lang:     Italian,
			},
			want: true,
		},
		{
			name: "Korean",
			args: args{
				mnemonic: "전망 차선 이전 실장 기간 간판 대접 판단 생명 존재 잠깐 건축",
				lang:     Korean,
			},
			want: true,
		},
		{
			name: "invalid checksum",
			args: args{
				mnemonic: "ivory disorder hawk slot oil promote north fat zebra useless device cargo",
				lang:     English,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsMnemonicValid(tt.args.mnemonic, tt.args.lang); got != tt.want {
				t.Errorf("IsMnemonicValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
