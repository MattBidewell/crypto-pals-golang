package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
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

func readFile(filename string) []string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(content), "\n")
}

type testResult struct {
	xorValue string
	score int
	xorRune rune
}


/**
*	Iterates through each file
*	calculate every possible decoded value using all possible bites.
* only save the highest scored decoded string.
* take all the highest strings and get the hightest string from that.
**/
func main() {

	file := readFile("data.txt")

	// save highest iterations of each mutation
	highestScoredSet := make([]testResult, len(file))

	// iterate through all lines in file
	for _, str := range file {
		hex := decodeHexString(string(str))
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
		highestScoredSet = append(highestScoredSet, highest)
		// fmt.Printf("String: %s \nRune: %s\nScore: %d", highest.xorValue, string(highest.xorRune), highest.score)
	}

	highestScoredResult := testResult{}
	for _, result := range highestScoredSet {
		if highestScoredResult.score < result.score {
			highestScoredResult = result
		}
	}

	fmt.Printf("String: %sRune: %s\nScore: %d", highestScoredResult.xorValue, string(highestScoredResult.xorRune), highestScoredResult.score)

}
