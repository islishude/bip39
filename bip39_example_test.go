package bip39

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func ExampleMnemonicToSeed() {
	mnemonic := "jungle devote wisdom slim census orbit merge order flip sketch add mass"
	seed := MnemonicToSeed(mnemonic, "")
	fmt.Println(base64.StdEncoding.EncodeToString(seed))

	// Output:
	// 84+1qMowq+jNijVCiRHXJ32RJevC8Bml1ADNdl0fb1alUvC5L4FJsNg09+W2SnFUiECrfUyCalh4NCnjZrIcrw==
}

func ExampleIsMnemonicValid() {
	const mnemonic = "check fiscal fit sword unlock rough lottery tool sting pluck bulb random"
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
