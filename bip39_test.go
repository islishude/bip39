package bip39

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"io"
	"testing"
)

func TestNewMnemonic(t *testing.T) {
	type args struct {
		wordsLen int
		lang     Language
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
		rander  io.Reader
	}{
		{
			name:    "invalid words length",
			args:    args{wordsLen: 1},
			want:    "",
			wantErr: true,
			rander:  rand.Reader,
		},
		{
			name:    "words length greater than 24",
			args:    args{wordsLen: 25},
			want:    "",
			wantErr: true,
			rander:  rand.Reader,
		},
		{
			name:    "words length isn't multiple of 3",
			args:    args{wordsLen: 13},
			want:    "",
			wantErr: true,
			rander:  rand.Reader,
		},
		{
			name:    "invalid random reader",
			args:    args{wordsLen: 12},
			want:    "",
			wantErr: true,
			rander:  bytes.NewReader(nil),
		},
		{
			name: "words length is ok",
			args: args{wordsLen: 12, lang: English},
			want: "betray shoe olive vivid nurse concert wonder early image castle route avocado",
			rander: bytes.NewBuffer([]byte{
				21, 120, 206, 104, 250,
				153, 120, 93, 127, 66,
				41, 113, 68, 114, 242,
				7,
			}),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cryptoRander = tt.rander
			got, err := NewMnemonic(tt.args.wordsLen, tt.args.lang)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMnemonic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewMnemonic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMnemonicByEntropy(t *testing.T) {
	type args struct {
		entropy []byte
		lang    Language
		skip    bool
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "entropy length is less than 16",
			args:    args{entropy: make([]byte, 1)},
			want:    "",
			wantErr: true,
		},
		{
			name:    "entropy length is greater than 32",
			args:    args{entropy: make([]byte, 33)},
			want:    "",
			wantErr: true,
		},
		{
			name:    "entropy length is not multiple of 4",
			args:    args{entropy: make([]byte, 17)},
			want:    "",
			wantErr: true,
		},
		{
			name:    "entropy length is ok",
			args:    args{entropy: make([]byte, 16), skip: true},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMnemonicByEntropy(tt.args.entropy, tt.args.lang)
			if (err != nil) != tt.wantErr {
				t.Errorf("MnemonicByEntropy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want && !tt.args.skip {
				t.Errorf("MnemonicByEntropy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMnemonicToSeed(t *testing.T) {
	type args struct {
		mnemonic string
		passwd   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Without password",
			args: args{
				mnemonic: "moment butter trigger coffee divert choose slim tiger ice series cup enough",
				passwd:   "",
			},
			want: "4b8c14466dbad77f6ff3adf016d372fbccfb0308ea5a36c9ab0c6f6eb1162ca461c02a1df2a1b854291785e59f0d98eb39af4d02a0ca8ffae5f66ff2dd0e2a48",
		},
		{
			name: "With password",
			args: args{
				mnemonic: "coffee purity language speed anger whisper ramp burden response brief coast trigger",
				passwd:   "bip39",
			},
			want: "ddfb143f00d7c135a59f1a05d00d2477a3eaa8ebfc1d4a4ddf2875d03cb74635458161a40faab128b4b8e1aeed75a919508a2816e7ef0a282105ad8ae48c91eb",
		},
		{
			name: "Janpanese",
			args: args{
				mnemonic: "こころ　いどう　きあつ　そうがんきょう　へいあん　せつりつ　ごうせい　はいち　いびき　きこく　あんい　おちつく　きこえる　けんとう　たいこ　すすめる　はっけん　ていど　はんおん　いんさつ　うなぎ　しねま　れいぼう　みつかる",
				passwd:   "㍍ガバヴァぱばぐゞちぢ十人十色",
			},
			want: "43de99b502e152d4c198542624511db3007c8f8f126a30818e856b2d8a20400d29e7a7e3fdd21f909e23be5e3c8d9aee3a739b0b65041ff0b8637276703f65c2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MnemonicToSeed(tt.args.mnemonic, tt.args.passwd)
			if hex.EncodeToString(got) != tt.want {
				t.Errorf("MnemonicToSeed() = %v, want %v", got, tt.want)
			}
		})
	}
}
