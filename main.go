package main

import (
	"errors"
)

func Compress(str string) ([]byte, error) {
	inBytes := []byte(str)
	howManyBits := uint(0)
	result := []byte{}

	for i, b := range inBytes {
		if (b & 0x80) == 0x80 {
			return nil, errors.New("Can't compress non-ascii character")
		}

		howManyBits += 1
		if howManyBits == 8 {
			howManyBits = 0
			continue
		}

		if len(inBytes)-1 == i {
			result = append(result, (b << howManyBits))
			return result, nil
		}

		nextByte := (b << howManyBits) |
			(inBytes[i+1] >> (7 - howManyBits))
		result = append(result, nextByte)
	}
	return nil, nil
}

func Decompress(bts []byte) (string, error) {
	resultBytes := []byte{}
	bytesToTakeFromPrevious := uint(0)

	for i, b := range bts {
		nextByte := (b >> bytesToTakeFromPrevious)
		if bytesToTakeFromPrevious > 0 {
			nextByte = nextByte | (bts[i-1] << (8 - bytesToTakeFromPrevious))
		}
		nextByte = nextByte >> 1

		resultBytes = append(resultBytes, nextByte)

		if bytesToTakeFromPrevious == 7 {
			resultBytes = append(resultBytes, (b>>1)&0x7f)
			bytesToTakeFromPrevious = 0
		}

		bytesToTakeFromPrevious += 1
	}

	return string(resultBytes), nil
}
