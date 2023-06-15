package main

import (
	"fmt"
	"testing"
)

func TestVal2intarr1(t *testing.T) {
	var v uint64
	v = 3*1024 + 2
	arr := val2intarr(v)
	res := arr[1] // reverse order!!!
	expected := 3

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}

func TestVal2intarr2(t *testing.T) {
	var v uint64
	v = 0xAB<<20 + 3<<10 + 2
	arr := val2intarr(v)
	res := arr[2] // reverse order!!!
	expected := 0xAB

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}

func TestBytes2UintManual1(t *testing.T) {
	bs := []byte{0xA, 0x1}
	res := Bytes2UintManual(bs)
	expected := uint64(0xA*0x100 + 1)

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}

func TestBytes2UintManual2(t *testing.T) {
	bs := []byte{0xAB, 0xCD}
	res := Bytes2UintManual(bs)
	expected := uint64(0xAB*0x100 + 0xCD)

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}

func TestUint2Bytes2(t *testing.T) {
	bs := []byte{0xAB, 0xCD}
	v := Bytes2UintManual(bs)

	//v := uint64(0xAB*0x100 + 0xCD)
	bs2 := Uint2BytesManual(v)
	res := bs2[0]
	expected := byte(0xAB)

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}

func TestUint2Bytes3(t *testing.T) {
	bs := []byte{0xFF, 0xFE}
	v := Bytes2UintManual(bs)

	//v := uint64(0xAB*0x100 + 0xCD)
	bs2 := Uint2BytesManual(v)
	res := bs2[0]
	expected := byte(0xFF)

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}

	res = bs2[1]
	expected = byte(0xFE)

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}

func TestUint2Bytes4(t *testing.T) {
	bs := []byte{0xFF, 0xFE, 0xAB, 0x02}
	v := Bytes2UintManual(bs)

	bs2 := Uint2BytesManual(v)
	for i := 0; i < len(bs); i++ {
		res := bs[i]
		expected := bs2[i]

		if res != expected {
			t.Errorf("Result: %v, want: %v", res, expected)
		}
	}
}

func TestUint2Bytes5(t *testing.T) {
	var bs []byte
	for i := 0; i < 7; i++ {
		bs = append(bs, byte(i+0xAA))
	}

	fmt.Printf("Source: %X\n", bs)

	//bs := []byte{0xFF, 0xFE, 0xAB, 0x02}
	v := Bytes2UintManual(bs)
	fmt.Printf("UInt64: %v\n", v)

	bs2 := Uint2BytesManual(v)
	fmt.Printf("Result: %X\n", bs2)

	for i := 0; i < len(bs); i++ {
		res := bs[i]
		expected := bs2[i]

		if res != expected {
			t.Errorf("Result: %v, want: %v", res, expected)
		}
	}
}

func TestUint2BytesBug1(t *testing.T) {
	var bs []byte
	for i := 0; i < 6; i++ {
		bs = append(bs, byte(i+1))
	}

	fmt.Printf("Source: %X\n", bs)

	//bs := []byte{0xFF, 0xFE, 0xAB, 0x02}
	v := Bytes2UintManual(bs)
	fmt.Printf("UInt64: %v\n", v)

	bs2 := Uint2BytesManual(v)
	fmt.Printf("Result: %X\n", bs2)

	for i := 0; i < len(bs); i++ {
		res := bs2[i]
		expected := bs[i]

		if res != expected {
			t.Errorf("Result: %v, want: %v", res, expected)
		}
	}
}
