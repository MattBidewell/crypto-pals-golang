package main

import (
	"encoding/hex"
	"fmt"
)

func decodeHexString(str string) ([]byte) {
	hex, err := hex.DecodeString(str)
	if err != nil {
		panic(err)
	}

	return hex
}

func xorHex(hex1 []byte, test byte)  ([]byte) {
	byteAr3 := make([]byte, len(hex1))

	// xor against test byte
	for i := range hex1 {
		byteAr3[i] = hex1[i] ^ test
	}

	return byteAr3
}


// func testString()

func main() {
	// testString := "etaoin shrdlu"

	hex := decodeHexString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")

	for i := 1; i < 256; i++ {
		xorHex := xorHex(hex, byte(i))
		fmt.Println(string(xorHex))
	}

}