package main

import "log"

func add32(a, b int) int {
	A := uint32(a)
	B := uint32(b)
	return int(A + B)
}

func rightRotate32(x, n int) int {
	if x >= 4294967296 {
		log.Fatal("x is too large. Did you use + instead of add32 somewhere?", x)
	}
	right := int(uint32(x) >> n)
	left := int(uint32(x) << (32 - n))
	return add32(left, right)
}

func littleSigma0(x int) int {
	return rightRotate32(x, 7) ^ rightRotate32(x, 18) ^ (x >> 3)
}

func littleSigma1(x int) int {
	return rightRotate32(x, 17) ^ rightRotate32(x, 19) ^ (x >> 10)
}
