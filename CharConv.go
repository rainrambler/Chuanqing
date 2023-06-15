package main

import (
	"fmt"
)

const (
	MultiFactor = 1024
	DeltaPos    = 10
	ByteLen     = 8
)

type CharConv struct {
	chars []rune
	c2pos map[rune]int // Key: rune, val: position in chars
}

func (p *CharConv) Load(filename string) {
	s, err := ReadTextFile(filename)
	if err != nil {
		fmt.Printf("Cannot read file: %s: %v!\n", filename, err)
		return
	}

	rs := []rune(s)
	p.chars = append(p.chars, rs...)

	p.c2pos = make(map[rune]int)
	for i, c := range p.chars {
		p.c2pos[c] = i
	}

	fmt.Printf("Loaded. Total %d chars.\n", len(p.chars))
}

func (p *CharConv) Encrypt(bs []byte) string {
	fmt.Printf("CharConv input: [%X]:%d\n", bs, len(bs))
	s := ""
	num := len(bs) / 5
	for i := 0; i < num; i++ {
		part := bs[i*5 : (i+1)*5]
		v := Bytes2UintManual(part)
		s += p.ValtoStr(v)
	}

	rbs := bs[num*5:]
	s += p.ValtoStr(Bytes2UintManual(rbs))
	return s
}

func (p *CharConv) Decrypt(s string) []byte {
	fmt.Printf("Decrypt In: [%s]\n", s)
	b := []byte{}
	rs := []rune(s)
	num := len(rs) / 4
	for i := 0; i < num; i++ {
		part := rs[i*4 : (i+1)*4]
		v := p.StrtoVal(part)
		bs := Uint2BytesManual(v)
		//fmt.Printf("DEC: [%d]: %v (%X) to [%X]\n", i, v, v, bs)
		b = append(b, bs...)
	}

	part := rs[num*4:]
	v := p.StrtoVal(part)
	bs := Uint2BytesManual(v)
	b = append(b, bs...)

	//fmt.Printf("[%X]:%d\n", b, len(b))
	return b
}

func Uint2BytesManual(v uint64) []byte {
	arr := []byte{}
	u := v
	for u > 0xFF {
		arr = append(arr, byte(u))
		u = u >> 8
	}
	arr = append(arr, byte(u))
	return Reverse(arr)
}

func Reverse(input []byte) []byte {
	inputLen := len(input)
	output := make([]byte, inputLen)

	for i, n := range input {
		j := inputLen - i - 1

		output[j] = n
	}

	return output
}

// Go\src\encoding\binary\binary.go littleEndian Uint64
// First byte SHOULD NOT be zero!!
func Bytes2UintManual(b []byte) uint64 {
	var v uint64
	for i := 0; i < len(b); i++ {
		v = v<<8 | uint64(b[i])
	}
	return v
}

func val2intarr(v uint64) []int {
	arr := []int{}

	// Go\src\strconv\itoa.go:formatBits(...)
	u := v
	b := uint64(MultiFactor)
	for u >= b {
		q := u / b
		remain := int(u - q*b)
		arr = append(arr, remain)
		u = q
	}

	arr = append(arr, int(u))
	return arr
}

func (p *CharConv) ValtoStr(v uint64) string {
	rs := []rune{}
	arr := val2intarr(v)
	for _, item := range arr {
		c := p.chars[item]
		rs = append(rs, c)
	}

	res := string(rs)
	return res
}

func (p *CharConv) StrtoVal(rs []rune) uint64 {
	v := uint64(0)
	for i := len(rs) - 1; i >= 0; i-- {
		c := rs[i]
		pos, exists := p.c2pos[c]
		if exists {
			v = v<<DeltaPos | uint64(pos)
		} else {
			fmt.Printf("Cannot find char in array: %s!\n", string(c))
			v = v << DeltaPos // ???
		}
	}
	return v
}
