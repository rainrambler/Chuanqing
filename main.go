package main

import (
	"fmt"
)

const KeyDemo = "aabbccdd"

func gcm_demo() {
	s := EncryptGcm(`今天中午去哪里吃？`, KeyDemo)
	DecryptGcm(s, KeyDemo)
}

func ecb_demo() {
	s := EncryptMsg(`今天中午去哪里吃？`, KeyDemo)
	fmt.Printf("Enc: %s\n", s)
	res := DecryptMsg(s, KeyDemo)
	fmt.Printf("Plain: %s\n", res)
}

func ecb_char_demo() {
	s := EncryptToSongci(`今天中午去哪里吃？`, KeyDemo)
	fmt.Printf("Enc: %s\n", s)
	res := DecryptFromSongci(s, KeyDemo)
	fmt.Printf("Plain: %s\n", res)
}

func main() {
	ecb_char_demo()
}
