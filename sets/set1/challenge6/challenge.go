// cracking the file
package main

import (
	"fmt"
	"strconv"
)

// 1. base64 decode file.
// & write function to calculate the distance between two strings.

func xorTwoHex(hex1 []byte, hex2 []byte)  ([]byte) {
	byteAr3 := make([]byte, len(hex1))

	for i := range hex1 {
		byteAr3[i] = hex1[i] ^ hex2[i]
	}

	return byteAr3
}

func calculateHammingWeight(str1 string, str2 string) (int) {
	xoredVal := xorTwoHex([]byte(str1), []byte(str2))
	totalOneBits := 0
	for _, val := range xoredVal {
		byteStr := strconv.FormatInt(int64(val), 2)
		for _, str := range byteStr {
			if byte(str) == "1"[0] {
				totalOneBits++
			}
		}
	}
	return totalOneBits
}

func main() {
	str1 := "this is a test"
	str2 := "wokka wokka!!!"
	tmp := calculateHammingWeight(str1, str2)
	fmt.Println(tmp)
}