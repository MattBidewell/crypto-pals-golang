package main

import (
	"encoding/hex"
	"fmt"
	"strings"
)

// decode a hexadecimal string to a byte array
func decodeHexString(str string) ([]byte) {
	hex, err := hex.DecodeString(str)
	if err != nil {
		panic(err)
	}

	return hex
}

// xored against a single character (rune)
func xorHex(hex1 []byte, test rune)  ([]byte) {
	byteAr3 := make([]byte, len(hex1))

	// xor against test byte
	// xor example
	// 110101
	// 101010
	// ------ xored
	// 011111
	for i := range hex1 {
		byteAr3[i] = hex1[i] ^byte(test)
	}

	return byteAr3
}

func calculateScore(str string) (int) {

	testString := "etaoin shrdlu"

	score := 0
	for _, char := range str {
		if strings.ContainsRune(testString, char) {
			score++
		}
	}
	return score
}

type testResult struct {
	xorValue string
	score int
	xorRune rune
}

func main() {

	hex := decodeHexString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")

	results := make([]testResult, 256)

	for i := 1; i < 256; i++ {
		xorHex := string(xorHex(hex, rune(i)))

		score := calculateScore(xorHex)

		result := testResult{}
		result.xorValue = xorHex
		result.score = score
		result.xorRune = rune(i)

		results = append(results, result)
	}

	highest := testResult{}
	for _, result := range results {
		if highest.score < result.score {
			highest = result
		}
	}
	fmt.Printf("String: %s \nRune: %s\nScore: %d", highest.xorValue, string(highest.xorRune), highest.score)
}