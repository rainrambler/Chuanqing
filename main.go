package main

import (
	"fmt"
)

const KeyDemo = "aabbccdd"

var convInst CharConv

func ecb_char_demo() {
	s := EncryptToSongci(`今天中午去哪里吃？`, KeyDemo)
	fmt.Printf("Enc: %s\n", s)
	res := DecryptFromSongci(s, KeyDemo)
	fmt.Printf("Plain: %s\n", res)
}

func main() {
	convInst.Load(`SongciRank.txt`)
	ecb_char_demo()
}
