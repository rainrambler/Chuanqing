package main

import (
	"fmt"

	"golang.org/x/text/encoding/simplifiedchinese"
)

// Convert UTF-8 to GBK
func GbkEncode(s string) []byte {
	ret, err := simplifiedchinese.GBK.NewEncoder().Bytes([]byte(s))
	if err != nil {
		fmt.Printf("Cannot convert to GBK: %s\n", s)
		return []byte{}
	}
	return ret
}

// Convert GBK to UTF-8
func GbkDecode(bs []byte) string {
	ret, err := simplifiedchinese.GBK.NewDecoder().Bytes(bs)
	if err != nil {
		fmt.Printf("Cannot convert to UTF-8: %s\n", string(bs))
		return ""
	}
	return string(ret)
}
