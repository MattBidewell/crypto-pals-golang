package main

import (
	"encoding/hex"
	"fmt"
)

func encodeHex(str []byte) (string) {
return hex.EncodeToString(str)
}

func xorHex(hex []byte, key string) ([]byte) {
	encryptedHex := make([]byte, len(hex))

	for i, char := range hex {
		// fmt.Printf("encoded %s with key %s\n", string(char), string(key[i % len(key)]))
		encryptedHex[i] = char ^ byte(rune(key[i % len(key)]))
	}

	return encryptedHex
}

/**
* Takes a string and a key and XORs all bytes with the
* equivilent byte in the key. Once at the end of the key,use
* the remanider function to pick off the first byte again.
*/
func main() {
	key := "ICE"
	hexString := "Burning 'em, if you ain't quick and nimble I go crazy when I hear a cymbal"

	hex := []byte(hexString)
	encryptedString := xorHex(hex, key)
	fmt.Println(string(encodeHex(encryptedString)))
}


