package bip39

import (
	"encoding/hex"
	"fmt"
)

func ExampleMnemonicToSeed() {
	mnemonic := "jungle devote wisdom slim census orbit merge order flip sketch add mass"
	res := MnemonicToSeed(mnemonic, "")
	fmt.Println(hex.EncodeToString(res))

	// Output:
	// f38fb5a8ca30abe8cd8a35428911d7277d9125ebc2f019a5d400cd765d1f6f56a552f0b92f8149b0d834f7e5b64a71548840ab7d4c826a58783429e366b21caf
}

func ExampleNewMnemonic() {
	// Words length can be 12 | 15 | 18 | 21 | 24
	NewMnemonic(12, English)
	NewMnemonic(24, ChineseTraditional)
	NewMnemonic(12, English)
	NewMnemonic(15, French)
	NewMnemonic(18, Italian)
	NewMnemonic(21, Japanese)
	NewMnemonic(24, French)
}

func ExampleIsMnemonicValid() {
	var mnemonic = "check fiscal fit sword unlock rough lottery tool sting pluck bulb random"
	res := IsMnemonicValid(mnemonic, English)
	fmt.Println(res)

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
