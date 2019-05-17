package main

import (
	"github.com/islishude/bip39"
)

func main() {

	m, err := bip39.NewMnemonic(12, bip39.English)
	if err != nil {
		panic(err)
	}

	if err := bip39.CheckMnemonic(m, bip39.English); err != nil {
		panic(err)
	}
}
