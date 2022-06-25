package main

import (
	"encoding/hex"
	"fmt"
)

func xorTwoHex(hex1 []byte, hex2 []byte)  ([]byte) {
	byteAr3 := make([]byte, len(hex1))

	for i := range hex1 {
		byteAr3[i] = hex1[i] ^ hex2[i]
	}

	return byteAr3
}

func decodeHexString(str string) ([]byte) {
	hex, err := hex.DecodeString(str)
	if err != nil {
		panic(err)
	}

	return hex
}

func encodeStringToHex(hexArr []byte) (string) {
	return hex.EncodeToString(hexArr)
}

func main() {

	hex1 := decodeHexString("1c0111001f010100061a024b53535009181c")
	hex2 := decodeHexString("686974207468652062756c6c277320657965")

	byteAr3 := xorTwoHex(hex1, hex2)

	strOutput := encodeStringToHex(byteAr3)
	fmt.Println(strOutput)
}