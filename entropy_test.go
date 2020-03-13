package bip39

import (
	"encoding/hex"
	"testing"
)

func Test_fromEntropy(t *testing.T) {
	type args struct {
		entropy string
		wordLen int
		lg      Language
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "15",
			args: args{
				entropy: "8e8bf76c330d126d3ba872a96f70af1b96e4b549",
				wordLen: 15,
				lg:      English,
			},
			want: "model garden swallow gravity spell custom upgrade atom practice knee cloth damp hour follow category",
		},
		{
			name: "12",
			args: args{
				entropy: "126f3c8b10757e43bbfd48d79e861d03",
				wordLen: 12,
				lg:      English,
			},
			want: "bar ketchup carpet can fitness canyon useful poverty stuff vintage mansion all",
		},
		{
			name: "japanese",
			args: args{
				entropy: "823be3d84e6ce7494001d42949f9ce391fe45616",
				wordLen: 15,
				lg:      Japanese,
			},
			want: "そらまめ　ほとんど　らくがき　ていか　ひみつ　てんらんかい　あいこくしん　くうふく　かいほう　こさめ　せいかつ　すめし　ろれつ　かたい　さつえい",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entropy, err := hex.DecodeString(tt.args.entropy)
			if err != nil {
				t.Error(err)
			}
			if got := fromEntropy(entropy, tt.args.wordLen, tt.args.lg); got != tt.want {
				t.Errorf("fromEntropy() = %v, want %v", got, tt.want)
			}
		})
	}
}
