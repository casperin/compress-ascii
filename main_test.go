package main

import (
	"bytes"
	"testing"
)

func TestCompress(t *testing.T) {
	bts, err := Compress("hello world")
	if err != nil {
		t.Fail()
	}

	//0b11010001
	//0b10010111
	//0b01100110
	//0b11001101
	//0b11101000
	//0b00111011
	//0b11101111
	//0b11100101
	//0b10110011
	//0b00100000
	expectedBytes := []byte{
		0xd1,
		0x97,
		0x66,
		0xcd,
		0xe8,
		0x3b,
		0xef,
		0xe5,
		0xb3,
		0x20,
	}

	if !bytes.Equal(bts, expectedBytes) {
		t.Errorf("%v != %v", bts, expectedBytes)
	}

}

func TestCompressError(t *testing.T) {
	_, err := Compress("ñato")
	if err == nil {
		t.Fail()
	}
}

func TestDecompress(t *testing.T) {
	str, err := Decompress([]byte{
		0xd1,
		0x97,
		0x66,
		0xcd,
		0xe8,
		0x3b,
		0xef,
		0xe5,
		0xb3,
		0x20,
	})

	if err != nil {
		t.Fail()
	}

	if str != "hello world" {
		t.Errorf("%s != %s, %v != %v", str, "hello world", []byte(str), []byte("hello world"))
	}

}

func TestCompressAndDecompress(t *testing.T) {
	loremIpsum := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum"
	bts, _ := Compress(loremIpsum)
	decompressed, _ := Decompress(bts)

	if decompressed != loremIpsum {
		t.Fail()
	}
}

func TestCompressAndDecompress8Chars(t *testing.T) {
	txt := "12345678"
	bts, _ := Compress(txt)
	decompressed, _ := Decompress(bts)
	if decompressed != txt {
		t.Fail()
	}
}

func TestCompressAndDecompress16Chars(t *testing.T) {
	txt := "1234567812345678"
	bts, _ := Compress(txt)
	decompressed, _ := Decompress(bts)
	if decompressed != txt {
		t.Fail()
	}
}
