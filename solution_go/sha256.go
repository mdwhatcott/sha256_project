package main

import "encoding/binary"

func rightRotate32(x uint32, n int) uint32 {
	right := x >> n
	left := x << (32 - n)
	return left + right
}

func littleSigma0(x uint32) uint32 {
	return rightRotate32(x, 7) ^ rightRotate32(x, 18) ^ (x >> 3)
}

func littleSigma1(x uint32) uint32 {
	return rightRotate32(x, 17) ^ rightRotate32(x, 19) ^ (x >> 10)
}

func messageSchedule(input string) (result []uint32) {
	result = make([]uint32, 64)
	for i := 0; i < 16; i++ {
		n := i * 4
		result[i] = binary.BigEndian.Uint32([]byte(input[n : n+4]))
	}
	for i := 16; i < 64; i++ {
		result[i] = result[i-16] + littleSigma0(result[i-15]) + result[i-7] + littleSigma1(result[i-2])
	}
	return result
}
