package bip39

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"testing"
)

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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsMnemonicValid(tt.args.mnemonic, tt.args.lang); got != tt.want {
				t.Errorf("IsMnemonicValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleIsMnemonicValid() {
	var mnemonic = "check fiscal fit sword unlock" +
		" rough lottery tool sting pluck bulb random"
	fmt.Println(IsMnemonicValid(mnemonic, English))

	// Output:
	// true
}

func TestNewMnemonic(t *testing.T) {
	type args struct {
		wordsLen int
		lang     Language
		skip     bool
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "words length less than 12",
			args:    args{wordsLen: 1},
			want:    "",
			wantErr: true,
		},
		{
			name:    "words length greater than 24",
			args:    args{wordsLen: 25},
			want:    "",
			wantErr: true,
		},
		{
			name:    "words length isn't multiple of 3",
			args:    args{wordsLen: 13},
			want:    "",
			wantErr: true,
		},
		{
			name:    "words length is ok",
			args:    args{wordsLen: 12, skip: true},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMnemonic(tt.args.wordsLen, tt.args.lang)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMnemonic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want && !tt.args.skip {
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
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Mnemonic is empty",
			args: args{
				mnemonic: "",
				passwd:   "",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Without password",
			args: args{
				mnemonic: "moment butter trigger coffee divert choose slim tiger ice series cup enough",
				passwd:   "",
			},
			want:    "4b8c14466dbad77f6ff3adf016d372fbccfb0308ea5a36c9ab0c6f6eb1162ca461c02a1df2a1b854291785e59f0d98eb39af4d02a0ca8ffae5f66ff2dd0e2a48",
			wantErr: false,
		},
		{
			name: "With password",
			args: args{
				mnemonic: "coffee purity language speed anger whisper ramp burden response brief coast trigger",
				passwd:   "bip39",
			},
			want:    "ddfb143f00d7c135a59f1a05d00d2477a3eaa8ebfc1d4a4ddf2875d03cb74635458161a40faab128b4b8e1aeed75a919508a2816e7ef0a282105ad8ae48c91eb",
			wantErr: false,
		},
		{
			name: "Janpanese",
			args: args{
				mnemonic: "こころ　いどう　きあつ　そうがんきょう　へいあん　せつりつ　ごうせい　はいち　いびき　きこく　あんい　おちつく　きこえる　けんとう　たいこ　すすめる　はっけん　ていど　はんおん　いんさつ　うなぎ　しねま　れいぼう　みつかる",
				passwd:   "㍍ガバヴァぱばぐゞちぢ十人十色",
			},
			want:    "43de99b502e152d4c198542624511db3007c8f8f126a30818e856b2d8a20400d29e7a7e3fdd21f909e23be5e3c8d9aee3a739b0b65041ff0b8637276703f65c2",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MnemonicToSeed(tt.args.mnemonic, tt.args.passwd)
			if (err != nil) != tt.wantErr {
				t.Errorf("MnemonicToSeed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(hex.EncodeToString(got), tt.want) {
				t.Errorf("MnemonicToSeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_entropyToMnemonic(t *testing.T) {
	type args struct {
		hexdata  string
		wordsLen int
		language Language
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "0",
			args: args{
				hexdata:  "00000000000000000000000000000000",
				language: English,
			},
			want:    "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about",
			wantErr: false,
		},
		{
			name: "1",
			args: args{
				hexdata:  "7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f",
				language: English,
			},
			want:    "legal winner thank year wave sausage worth useful legal winner thank yellow",
			wantErr: false,
		},
		{
			name: "2",
			args: args{
				hexdata:  "000000000000000000000000000000000000000000000000",
				language: English,
			},
			want:    "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon agent",
			wantErr: false,
		},
		{
			name: "3",
			args: args{
				hexdata:  "8080808080808080808080808080808080808080808080808080808080808080",
				language: Japanese,
			},
			want:    "そとづら　あまど　おおう　あこがれる　いくぶん　けいけん　あたえる　いよく　そとづら　あまど　おおう　あこがれる　いくぶん　けいけん　あたえる　いよく　そとづら　あまど　おおう　あこがれる　いくぶん　けいけん　あたえる　うめる",
			wantErr: false,
		},
		{
			name: "4",
			args: args{
				hexdata:  "3e141609b97933b66a060dcddc71fad1d91677db872031e85f4c015c5e7e8982",
				language: Japanese,
			},
			want:    "くのう　てぬぐい　そんかい　すろっと　ちきゅう　ほあん　とさか　はくしゅ　ひびく　みえる　そざい　てんすう　たんぴん　くしょう　すいようび　みけん　きさらぎ　げざん　ふくざつ　あつかう　はやい　くろう　おやゆび　こすう",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entropy, err := hex.DecodeString(tt.args.hexdata)
			if err != nil {
				t.Errorf("entropyToMnemonic() decode hex string error %v", err)
				return
			}
			got, err := entropyToMnemonic(entropy, tt.args.wordsLen, tt.args.language)
			if (err != nil) != tt.wantErr {
				t.Errorf("entropyToMnemonic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("entropyToMnemonic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleNewMnemonic() {
	// Words length can be 12 | 15 | 18 | 21 | 24
	NewMnemonic(12, ChineseSimplified)
	NewMnemonic(24, ChineseTraditional)
	NewMnemonic(12, English)
	NewMnemonic(15, French)
	NewMnemonic(18, Italian)
	NewMnemonic(21, Japanese)
	NewMnemonic(24, French)
	NewMnemonic(15, Korean)
	NewMnemonic(15, Spanish)
}

func ExampleMnemonicToSeed() {
	mnemonic := "jungle devote wisdom slim" +
		" census orbit merge order flip sketch add mass"

	fmt.Println(IsMnemonicValid(mnemonic, English))

	// Output:
	// true
}

func ExampleNewMnemonicByEntropy() {
	entropy, _ := hex.DecodeString("79079bf165e25537e2dce15919440cc4")
	mnemonic, _ := NewMnemonicByEntropy(entropy, English)
	fmt.Println(mnemonic)

	// Output:
	// jungle devote wisdom slim census orbit merge order flip sketch add mass
}
